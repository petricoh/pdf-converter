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
	outputPath string
}

func NewRunner(pwd string, outputPath string) *Runner {
	fm := file.NewFileManager(pwd)
	return &Runner{fm, outputPath}
}

func (r *Runner) Run(inputPath string) {
	dirs, _ := r.fm.GetDirPaths(inputPath)
	result := utils.AwaitAll(dirs, r.genPdf)

	failures := utils.FilterFailures(result)
	for _, fail := range failures {
		fmt.Println(fail.Err)
	}
}

func (r *Runner) genPdf(dirPath string) (*exec.Cmd, error) {
	infos, err := r.fm.GetFileInfos(dirPath)
	if err != nil {
		return nil, err
	}

	abs, err := utils.GetProjectAbsolutePath()
	if err != nil {
		return nil, err
	}

	cmdPath := filepath.Join(abs, "bin", "pdfcpu.exe")
	outPath := filepath.Join(abs, r.outputPath, utils.GetFilename(dirPath)+".pdf")

	args := utils.Concat([]string{"import", outPath}, r.fm.GetPathsByInfos(infos))
	c := exec.Command(cmdPath, args...)
	c.Stderr = os.Stderr
	err = c.Run()

	filename := utils.GetFilename(dirPath)
	fmt.Printf("complete: %s\n", filename)

	return c, err
}
