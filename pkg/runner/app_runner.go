package runner

import (
	"fmt"
	"os"
	"time"

	"github.com/rwiv/pdfconv/pkg/executor"
	"github.com/rwiv/pdfconv/pkg/utils/pwdutil"
)

type AppRunner struct {
}

func NewAppRunner() AppRunner {
	return AppRunner{}
}

func (r *AppRunner) Run() {
	if len(os.Args) < 3 {
		fmt.Println("not found args")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	pwd, err := pwdutil.GetPwd()
	if err != nil {
		fmt.Println("not found pwd")
		return
	}

	e := executor.NewPdfCpuExecutor(pwd, inputPath, outputPath)
	err = e.ExecParallel()
	if err != nil {
		fmt.Println("not found pwd")
		return
	}

	time.Sleep(time.Second * 2)
}
