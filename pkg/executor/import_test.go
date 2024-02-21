package executor

import (
	"fmt"
	"testing"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/rwiv/pdfconv/pkg/utils/apath"
)

func TestName(t *testing.T) {
	ap, err := apath.GetProjectAbsPath()
	if err != nil {
		t.Fatal(err)
	}
	base := ap + "\\test\\images"
	img1 := base + "\\test1.png"
	img2 := base + "\\test2.png"
	fmt.Println(base)
	imgFiles := []string{img1, img2}
	outFile := ap + "\\test" + "\\out.pdf"
	err = pdfcpu.ImportImagesFile(imgFiles, outFile, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	//fm := file.NewFileManager(ap)
	//dirRelPaths, _ := fm.GetFileInfos(file.RelPath(ap + "\\test\\images"))
	//for _, elem := range dirRelPaths {
	//	fmt.Println(elem)
	//}
}
