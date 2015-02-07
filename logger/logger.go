package logger

import (
	"fmt"
	_ "io/ioutil"
)

func Log(msg string) {
	fmt.Println("[Log] " + msg)
}

func Error(msg string) {
	fmt.Println("[Error] " + msg)
}
