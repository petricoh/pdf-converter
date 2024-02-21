package executor

import (
	"testing"

	"github.com/rwiv/pdfconv/pkg/utils/apath"
)

func TestPdfCpuExecutor(t *testing.T) {
	ap, err := apath.GetProjectAbsPath()
	if err != nil {
		t.Fatal(err)
	}
	//inDirAbs := apath.ToAbsPath(ap, "\\test\\input")
	//outDirAbs := apath.ToAbsPath(ap, "\\test\\out")
	inDirAbs := "\\test\\input"
	outDirAbs := "\\test\\out"

	e := NewPdfCpuExecutor(ap, inDirAbs, outDirAbs)
	//err = e.ExecSync()
	err = e.ExecParallel()
	if err != nil {
		t.Fatal(err)
	}
}
