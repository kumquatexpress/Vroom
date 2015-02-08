package logger

import (
	"fmt"
	"os"
)

func Log(msg string) {
	fmt.Println("[Log] " + msg)
}

func Warn(msg string) {
	fmt.Println("[Warning] " + msg)
}

func Error(msg string) {
	fmt.Println("[Error] " + msg)
	os.Exit(1)
}
