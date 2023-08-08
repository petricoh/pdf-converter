package utils

import (
	"path/filepath"
	"strings"
)

func GetFilename(path string) string {
	s := string(filepath.Separator)
	chunks := strings.Split(path, s)
	return chunks[len(chunks)-1]
}
