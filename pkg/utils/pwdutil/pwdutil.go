package pwdutil

import (
	"os"
)

func GetPwd() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}
