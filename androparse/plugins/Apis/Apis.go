package main

import (
	"bytes"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
	"unicode"
)

func NeedLock() bool { return true }

func GetKey() string { return "Apis" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	prog := "java"
	args := []string{"-Dfile.encoding=UTF-8",
		"-Xmx512m",
		"-cp",
		config.CodeDir + "/plugins/Apis/:" + config.CodeDir + "/plugins/Apis/Rapid.jar",
		"APIParser",
		path}
	cmd := exec.Command(prog, args...)
	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		return []string{"ERROR: no DEX file is found in the APK file."}, err
	}
	if errout.String() != "" {
		log.Printf(errout.String())
	}
	data := strings.Split(out.String(), "\n")
	data = data[5:]
	for i := 0; i < len(data)-1; i++ {
		previous := false
		tmp := strings.Replace(data[i], "\t", "", -1)
		for _, c := range tmp {
			if !unicode.IsDigit(c) && previous {
				break
			}
			if unicode.IsDigit(c) {
				tmp = tmp[1:]
				previous = true
			}
		}
		data[i] = tmp
	}
	return data, nil
}
