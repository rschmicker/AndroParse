package main

import (
	"github.com/AndroParse/androparse/utils"
	"os"
)

func NeedLock() bool { return false }

func GetKey() string { return "FileSize" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return fi.Size(), nil
}
