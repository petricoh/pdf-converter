package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"pdf-converter/pkg/file"
	"pdf-converter/pkg/utils"
)

type FileInfo struct {
	Name        string
	Path        string
	IsDirectory bool
}

type Runner struct {
	fm         *file.Manager
	inputPath  string
	outputPath string
	pwd        string
}

func NewRunner(pwd string, inputPath string, outputPath string) *Runner {
	fm := file.NewFileManager(pwd)
	return &Runner{fm, inputPath, outputPath, pwd}
}

func (r *Runner) Run() {
	dirRelPaths, _ := r.fm.GetDirRelPaths(file.RelPath(r.inputPath))
	result := utils.AwaitAll(dirRelPaths, r.genPdf)

	failures := utils.FilterFailures(result)
	for _, fail := range failures {
		fmt.Println(fail.Err)
	}
}

func (r *Runner) genPdf(dirRelPath file.RelPath) (*exec.Cmd, error) {
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

	c := exec.Command(cmdPath, args...)
	c.Stderr = os.Stderr
	err = c.Run()

	filename := utils.GetFilename(string(dirRelPath))
	fmt.Printf("complete: %s\n", filename)

	return c, err
}
