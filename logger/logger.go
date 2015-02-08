package logger

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Log(msg string) {
	fmt.Println("[Log] " + msg)
}

func Warn(msg string) {
	fmt.Println("[Warning] " + msg)
}

func Error(msg string) {
	fmt.Println("[Error] " + msg)
	debug.PrintStack()
	os.Exit(1)
}
