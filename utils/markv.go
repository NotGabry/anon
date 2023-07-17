package AnonUtils

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Error(str string, exit bool) {

	red := color.New(color.FgRed).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()

	ErrorLog := fmt.Sprintf("%s%s%s %s\n", gray("["), red("x"), gray("]"), str)
	fmt.Print(ErrorLog)

	if exit {
		os.Exit(1)
	}
}

func SpinError(str string, exit bool, s *spinner.Spinner) {

	red := color.New(color.FgRed).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()

	s.FinalMSG = fmt.Sprintf("%s%s%s %s\n", gray("["), red("x"), gray("]"), str)
	s.Stop()

	if exit {
		os.Exit(1)
	}
}

func SpinCheck(str string, exit bool, s *spinner.Spinner) {

	green := color.New(color.FgGreen).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()

	s.FinalMSG = fmt.Sprintf("%s%s%s %s\n", gray("["), green("✔"), gray("]"), str)
	s.Stop()

	if exit {
		os.Exit(1)
	}
}
