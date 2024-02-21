package apath

import (
	"fmt"
	"testing"
)

func TestToAbsPath(t *testing.T) {
	p1 := "haha.pdf"
	p2 := "C:\\a\\haha.pdf"
	p3 := "/root/haha.pdf"
	fmt.Println(ToAbsPath("C:\\hello", p1))
	fmt.Println(ToAbsPath("C:\\hello", p2))
	fmt.Println(ToAbsPath("C:\\hello", p3))
}
