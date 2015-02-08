package utils

import (
	"fmt"
	"github.com/kumquatexpress/Vroom/logger"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func FindOrCreateDir(dir string) bool {
	if !Exists(dir) {
		logger.Log(fmt.Sprintf("No directory found, creating %s", dir))
		os.MkdirAll(dir, os.ModePerm)
		return false
	}
	return true
}

func MergeMap(oldM map[string]interface{}, newM map[string]interface{}) map[string]interface{} {
	for inner, val := range newM {
		oldM[inner] = val
	}
	return oldM
}
