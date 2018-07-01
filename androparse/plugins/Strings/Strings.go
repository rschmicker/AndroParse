package main

import (
	"bytes"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
)

func NeedLock() bool { return true }

func GetKey() string { return "Strings" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	prog := "java"
	args := []string{"-Dfile.encoding=UTF-8",
		"-Xmx512m",
		"-cp",
		config.CodeDir + "/plugins/Strings/:" + config.CodeDir + "/plugins/Apis/Rapid.jar",
		"StringParser",
		path}
	cmd := exec.Command(prog, args...)
	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		return []string{"ERROR:	no DEX file is found in the APK file."}, err
	}
	if errout.String() != "" {
		log.Printf(errout.String())
	}
	data := strings.Split(out.String(), "\n")
	data = data[5:]
	return data, nil
}
