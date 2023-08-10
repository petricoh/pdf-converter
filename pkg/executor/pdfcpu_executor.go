package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"pdf-converter/pkg/file"
	"pdf-converter/pkg/utils"
)

type PdfCpuExecutor struct {
	fm         *file.Manager
	inputPath  string
	outputPath string
	pwd        string
}

func NewPdfCpuExecutor(pwd string, inputPath string, outputPath string) *PdfCpuExecutor {
	fm := file.NewFileManager(pwd)
	return &PdfCpuExecutor{fm, inputPath, outputPath, pwd}
}

func (r *PdfCpuExecutor) Exec() error {
	if _, err := os.Stat(r.outputPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(r.outputPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dirRelPaths, _ := r.fm.GetDirRelPaths(file.RelPath(r.inputPath))
	result := utils.AwaitAll(dirRelPaths, r.genPdf)

	failures := utils.FilterFailures(result)
	for _, fail := range failures {
		fmt.Println(fail.Err)
	}

	return nil
}

func (r *PdfCpuExecutor) genPdf(dirRelPath file.RelPath) (*exec.Cmd, error) {
	infos, err := r.fm.GetFileInfos(dirRelPath)
	if err != nil {
		return nil, err
	}

	cmdPath, err := exec.LookPath("pdfcpu")
	if err != nil {
		return nil, err
	}

	outPath := filepath.Join(r.pwd, r.outputPath, utils.GetFilename(string(dirRelPath))+".pdf")
	absPaths := r.fm.ToStringsAbsPaths(r.fm.GetAbsPathsByInfos(infos))

	args := utils.Concat([]string{"import", outPath}, absPaths)

	cmd := exec.Command(cmdPath, args...)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	filename := utils.GetFilename(string(dirRelPath))
	fmt.Printf("complete: %s\n", filename)

	return cmd, nil
}
