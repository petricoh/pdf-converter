package executor

import (
	"errors"
	"fmt"
	"os"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/rwiv/pdfconv/pkg/utils/apath"
	"github.com/rwiv/pdfconv/pkg/utils/await"
	"github.com/rwiv/pdfconv/pkg/utils/fileutil"
	"github.com/rwiv/pdfconv/pkg/utils/list"
)

type PdfCpuExecutor struct {
	pwd       string
	inDirAbs  string
	outDirAbs string
}

func NewPdfCpuExecutor(pwd string, inputPath string, outputPath string) *PdfCpuExecutor {
	inDirAbs := apath.ToAbsPath(pwd, inputPath)
	outDirAbs := apath.ToAbsPath(pwd, outputPath)
	return &PdfCpuExecutor{pwd, inDirAbs, outDirAbs}
}

func (e *PdfCpuExecutor) ExecSync() error {
	// check out dir
	if _, err := os.Stat(e.outDirAbs); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(e.outDirAbs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// generate pdfs sync
	dirs, err := fileutil.ReadDir(e.inDirAbs)
	if err != nil {
		return err
	}
	for _, dirInfo := range dirs {
		err := e.genPdf(dirInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *PdfCpuExecutor) ExecParallel() error {
	// check out dir
	if _, err := os.Stat(e.outDirAbs); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(e.outDirAbs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// generate pdfs parallel
	dirs, err := fileutil.ReadDir(e.inDirAbs)
	if err != nil {
		return err
	}
	result := await.AwaitAll(dirs, func(info *fileutil.FileInfo) (*string, error) {
		err := e.genPdf(info)
		result := ""
		return &result, err
	})
	failures := await.FilterFailures(result)
	for _, fail := range failures {
		fmt.Println(fail.Err)
	}
	return nil
}

func (e *PdfCpuExecutor) genPdf(dirInfo *fileutil.FileInfo) error {
	files, err := fileutil.ReadDir(dirInfo.AbsPath)
	if err != nil {
		return err
	}
	imgFiles := list.Map(files, func(file *fileutil.FileInfo) string {
		return file.AbsPath
	})
	outFile := apath.ToAbsPath(e.outDirAbs, dirInfo.Name+".pdf")
	err = pdfcpu.ImportImagesFile(imgFiles, outFile, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
