package main

import (
	"crypto/md5"
	"fmt"
	"github.com/AndroParse/androparse/utils"
	"io"
	"os"
)

func NeedLock() bool { return false }

func GetKey() string { return "Md5" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
