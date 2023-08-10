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

func (e *PdfCpuExecutor) Exec() error {
	if _, err := os.Stat(e.outputPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(e.outputPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dirRelPaths, _ := e.fm.GetDirRelPaths(file.RelPath(e.inputPath))
	result := utils.AwaitAll(dirRelPaths, e.genPdf)

	failures := utils.FilterFailures(result)
	for _, fail := range failures {
		fmt.Println(fail.Err)
	}

	return nil
}

func (e *PdfCpuExecutor) genPdf(dirRelPath file.RelPath) (*exec.Cmd, error) {
	infos, err := e.fm.GetFileInfos(dirRelPath)
	if err != nil {
		return nil, err
	}

	cmdPath, err := exec.LookPath("pdfcpu")
	if err != nil {
		return nil, err
	}

	outPath := filepath.Join(e.pwd, e.outputPath, utils.GetFilename(string(dirRelPath))+".pdf")
	absPaths := e.fm.ToStringsAbsPaths(e.fm.GetAbsPathsByInfos(infos))

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
