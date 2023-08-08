package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	projectName = "pdf-converter"
)

func GetProjectAbsolutePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return GetAbsolutePathBy(wd, string(filepath.Separator))
}

func GetAbsolutePathBy(path string, separator string) (string, error) {
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
