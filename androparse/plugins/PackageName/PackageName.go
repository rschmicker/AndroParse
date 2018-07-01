package main

import (
	"bytes"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
)

func NeedLock() bool { return false }

func GetKey() string { return "PackageName" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	prog := "aapt"
	args := []string{
		"dump",
		"permissions",
		path}
	cmd := exec.Command(prog, args...)
	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	if errout.String() != "" {
		log.Printf(errout.String())
	}
	data := strings.Split(out.String(), "\n")
	return strings.Split(data[0], "package: ")[1], nil
}
