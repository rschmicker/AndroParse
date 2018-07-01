package controller

import (
	"encoding/json"
	"fmt"
	"github.com/AndroParse/androparse/cleaner"
	"github.com/AndroParse/androparse/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"sync"
)

type WorkerData struct {
	Sem    chan struct{}
	Config utils.ConfigData
	Count  *int
	Length int
}

var mutex = &sync.Mutex{}
var countMutex = &sync.Mutex{}
var wg sync.WaitGroup

func Runner(config utils.ConfigData) {
	count := 0
	files := Cleaner(config)
	wd := &WorkerData{
		Sem:    make(chan struct{}, runtime.NumCPU()),
		Config: config,
		Count:  &count,
		Length: len(files),
	}
	for _, file := range files {
		wg.Add(1)
		wd.Sem <- struct{}{}
		go worker(wd, file)
	}
	wg.Wait()
	close(wd.Sem)
}

func worker(wd *WorkerData, apk string) {
	log.Printf("Working on APK: " + apk)
	defer func() { <-wd.Sem }()
	defer wg.Done()
	err := extractFeatures(apk, wd.Config)
	if err != nil {
		log.Printf("Warning: " + apk + " is not a valid APK file")
		countMutex.Lock()
		*wd.Count++
		countMutex.Unlock()
		log.Println(err)
		return
	}
	countMutex.Lock()
	*wd.Count++
	countMutex.Unlock()
	percent := (float64(*wd.Count) / float64(wd.Length) * float64(100))
	_, name := filepath.Split(apk)
	log.Printf("(%.2f%%) Completed: "+name, percent)
}

func Cleaner(config utils.ConfigData) []string {
	if config.Clean {
		cleaner.CleanDirectory(config)
	}
	return utils.GetPaths(config.ApkDir, ".apk")
}

func getParsedJson(path string) (map[string]interface{}, error) {
	if _, err := os.Stat(path); err == nil {
		f, err := ioutil.ReadFile(path)
		utils.Check(err)
		var objs interface{}
		json.Unmarshal(f, &objs)
		features, ok := objs.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Warning: Could not unmarshal existing json file %v... Parsing all features", path)
		}
		return features, nil
	} else {
		return nil, fmt.Errorf("Warning: Could not unmarshal existing json file %v... Parsing all features", path)
	}
}

func createJsonBuilder(path string, config utils.ConfigData) map[string]interface{} {
	jsonBuilder := make(map[string]interface{})
	var err error
	if config.Append {
		jsonPath := config.OutputDir
		_, newPath := filepath.Split(path)
		jsonPath += "/" + newPath[:len(newPath)-4] + ".json"
		jsonBuilder, err = getParsedJson(jsonPath)
		if err != nil {
			log.Println(err.Error())
			jsonBuilder = make(map[string]interface{})
		}
	}
	return jsonBuilder
}

func createPlugins(config utils.ConfigData) []string {
	plugins := []string{}
	if len(config.Parsers) == 0 {
		plugins = utils.GetPaths(config.CodeDir+"/plugins/", ".so")
	} else {
		needSha256 := true
		for _, parser := range config.Parsers {
			needSha256 = !(parser == "Sha256")
			path := config.CodeDir + "/plugins/" + parser + ".so"
			if _, err := os.Stat(path); os.IsNotExist(err) {
				log.Println("Warning: Could not find plugin specified " + path)
			} else {
				plugins = append(plugins, path)
			}
		}
		if needSha256 {
			plugins = append(plugins, config.CodeDir+"/plugins/Sha256.so")
		}
	}
	return plugins
}

func inParsers(key string, config utils.ConfigData) bool {
	status := false
	for _, parser := range config.Parsers {
		if parser == key {
			status = true
			break
		}
	}
	return status
}

func extractFeatures(path string, config utils.ConfigData) error {
	jsonBuilder := createJsonBuilder(path, config)
	plugins := createPlugins(config)
	for _, plug := range plugins {
		p, err := plugin.Open(plug)
		utils.Check(err)
		needLock, err := p.Lookup("NeedLock")
		utils.Check(err)
		needLockFunc, ok := needLock.(func() bool)
		if !ok {
			log.Fatal("Error: Malformed NeedLock function in " + plug)
		}

		k, err := p.Lookup("GetKey")
		utils.Check(err)
		keyfunc, ok := k.(func() string)
		if !ok {
			log.Fatal("Error: Malformed GetKey function in " + plug)
		}
		key := keyfunc()

		if _, ok := jsonBuilder[key]; ok && !inParsers(key, config) {
			continue
		}

		v, err := p.Lookup("GetValue")
		utils.Check(err)
		valuefunc, ok := v.(func(string, utils.ConfigData) (interface{}, error))
		if !ok {
			log.Fatal("Error: Malformed GetValue function in " + plug)
		}
		if needLockFunc() {
			mutex.Lock()
		}
		result, err := valuefunc(path, config)
		if needLockFunc() {
			mutex.Unlock()
		}
		if err != nil {
			return err
		}
		jsonBuilder[key] = result
	}
	WriteJSON(jsonBuilder, config.OutputDir)
	return nil
}

func WriteJSON(toWrite map[string]interface{}, OutputDir string) {
	data, err := json.Marshal(toWrite)
	utils.Check(err)
	Sha256, ok := toWrite["Sha256"].(string)
	if !ok {
		log.Println(toWrite)
		log.Fatal("Error: Count not validate Sha256 value as a string")
	}
	outputFile := OutputDir + "/" + Sha256 + ".json"
	fo, err := os.Create(outputFile)
	utils.Check(err)
	fo.Write(data)
	fo.Close()
}
