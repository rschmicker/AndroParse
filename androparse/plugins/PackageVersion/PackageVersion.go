package main

import (
	"bytes"
	"errors"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
)

func NeedLock() bool { return false }

func GetKey() string { return "Version" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	prog := "aapt"
	args := []string{
		"dump",
		"badging",
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
	if strings.Contains(out.String(), "ERROR") {
		return "", errors.New("Not an APK")
	}
	data := strings.Split(out.String(), "\n")
	version := data[0]
	version = strings.Split(version, "versionName='")[1]
	version = strings.Split(version, "'")[0]
	return version, nil
}
