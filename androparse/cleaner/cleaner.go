package cleaner

import (
	"crypto/sha256"
	"fmt"
	"github.com/AndroParse/androparse/utils"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func Sha256File(path string) string {
	f, err := os.Open(path)
	utils.Check(err)
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	utils.Check(err)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CleanDirectory(config utils.ConfigData) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, runtime.NumCPU())
	fileList := utils.GetPaths(config.ApkDir, ".apk")
	for _, file := range fileList {
		wg.Add(1)
		go func(file string) {
			sem <- struct{}{}
			defer func() { <-sem }()
			defer wg.Done()
			hash := Sha256File(file)
			newPath := ""
			if strings.Contains(file, "benign") {
				newPath = config.ApkDir + "/benign/" + hash + ".apk"
			} else {
				newPath = config.ApkDir + "/malware/" + hash + ".apk"
			}
			if file == newPath {
				log.Printf("Skipping: %v already cleaned", file)
				return
			}
			err := os.Rename(file, newPath)
			utils.Check(err)
			log.Printf("Cleaned: %v -> %v", file, newPath)
		}(file)
	}
	wg.Wait()
	close(sem)
}
