package main

import (
	"bytes"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
)

func NeedLock() bool { return false }

func GetKey() string { return "Permissions" }

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
		return []string{}, err
	}
	if errout.String() != "" {
		log.Printf(errout.String())
	}
	tmp := strings.Split(out.String(), "\n")
	tmp = tmp[1:]
	data := []string{}
	for _, line := range tmp {
		if !strings.Contains(line, "permission: ") {
			continue
		}
		line = strings.Trim(line, " ")
		line = strings.Split(line, "permission: ")[1]
		if strings.Contains(line, "name='") {
			line = strings.Split(line, "name='")[1]
		}
		line = strings.Trim(line, "'")
		data = append(data, line)
	}
	return data, nil
}
