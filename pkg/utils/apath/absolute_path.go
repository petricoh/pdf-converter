package apath

import (
	"fmt"
	"path/filepath"
	"regexp"
)

const (
	projectName = "pdfconv"
)

func GetProjectAbsPath() (string, error) {
	ex, err := filepath.Abs("")
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return GetProjectAbsPathBy(exPath, string(filepath.Separator))
}

func GetProjectAbsPathBy(path string, separator string) (string, error) {
	if separator == `\` {
		separator = `\\`
	}
	expr := fmt.Sprintf(".*?%s%s", separator, projectName)
	reg, err := regexp.Compile(expr)
	if err != nil {
		return "", err
	}
	m := reg.FindString(path)

	return m, nil
}
