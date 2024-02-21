package apath

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
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

func IsAbsPath(path string) bool {
	if path[0] == '/' {
		return true
	} else if strings.Contains(path, ":") {
		return true
	} else {
		return false
	}
}

func ToAbsPath(base string, path string) string {
	if IsAbsPath(path) {
		return path
	}
	if strings.HasPrefix(path, string(filepath.Separator)) {
		return base + path
	} else {
		return base + string(filepath.Separator) + path
	}
}
