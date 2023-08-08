package utils

import (
	"os"
	"path/filepath"
)

func GetPwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
