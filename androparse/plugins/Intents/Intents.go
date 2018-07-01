package main

import (
	"bytes"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os/exec"
	"strings"
)

func NeedLock() bool { return false }

func GetKey() string { return "Intents" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	prog := "aapt"
	args := []string{
		"dump",
		"xmltree",
		path,
		"AndroidManifest.xml"}
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
	data := strings.Split(out.String(), "\n")
	intents := []string{}
	for _, line := range data {
		if strings.Contains(line, "android.intent.") {
			line = strings.Split(line, "(Raw: \"")[1]
			line = strings.Split(line, "\")")[0]
			intents = append(intents, line)
		}
	}
	return intents, nil
}
