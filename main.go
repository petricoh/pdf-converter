package main

import (
	"pdf-converter/pkg/runner"
	"pdf-converter/pkg/utils"
	"time"
)

const (
	inputPath  = "inputs"
	outputPath = "outs"
)

func main() {
	r := runner.NewRunner(utils.GetPwd(), outputPath)
	r.Run(inputPath)
	time.Sleep(time.Second * 2)
}
