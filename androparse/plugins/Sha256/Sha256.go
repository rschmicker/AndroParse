package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/AndroParse/androparse/utils"
	"io"
	"os"
)

func NeedLock() bool { return false }

func GetKey() string { return "Sha256" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
