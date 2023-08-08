package main

import (
	"fmt"
	"os"
	"pdf-converter/pkg/runner"
	"pdf-converter/pkg/utils"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("not found args")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	pwd, err := utils.GetPwd()
	if err != nil {
		fmt.Println("not found pwd")
		return
	}
	r := runner.NewRunner(pwd, inputPath, outputPath)
	r.Run()
	time.Sleep(time.Second * 2)
}
