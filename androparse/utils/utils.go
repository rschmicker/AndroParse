package utils

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ConfigData struct {
	ApkDir     string `yaml:"apkDir"`
	OutputDir  string `yaml:"outputDir"`
	CodeDir    string `yaml:"codeDir"`
	CacheDir   string `yaml:"cacheDir"`
	Clean      bool
	VtApiCheck bool
	Append     bool
	VtApiKey   string `yaml:"vtapikey"`
	Parsers    []string
}

func ReadConfig(configPath string) ConfigData {
	data, err := ioutil.ReadFile(configPath)
	Check(err)
	config := ConfigData{}
	err = yaml.Unmarshal(data, &config)
	Check(err)

	config.ApkDir, err = filepath.Abs(config.ApkDir)
	Check(err)
	config.OutputDir, err = filepath.Abs(config.OutputDir)
	Check(err)
	config.CodeDir, err = filepath.Abs(config.CodeDir)
	Check(err)
	config.CacheDir, err = filepath.Abs(config.CacheDir)
	Check(err)

	return config
}

func GetPaths(dir string, Containing string) []string {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileInfo, errS := os.Stat(path)
		Check(errS)
		if strings.Contains(path, Containing) && fileInfo.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		return err
	})
	Check(err)
	return fileList
}

func CrossCompare(todoFiles []string, doneFiles []string) []string {
	ret := []string{}
	found := false
	for _, todo := range todoFiles {
		_, name := filepath.Split(todo)
		name = name[:len(name)-4]
		for _, done := range doneFiles {
			if strings.Contains(done, name) {
				found = true
				log.Printf("Skipping: %v already completed...", todo)
				break
			}
		}
		if !found {
			ret = append(ret, todo)
		}
		found = false
	}
	return ret
}

func Check(err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStr := fmt.Sprintf("Error: in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		log.Fatal(errStr)
	}
}
