package cli

import (
	"log"
	"github.com/fatih/color"
)

func Log(arg1 string, arg2 string) {
	log.Println(color.BlueString(arg1), color.RedString(arg2))
}

func ErrLog(reason string) {
	log.Println(color.YellowString(reason))
}