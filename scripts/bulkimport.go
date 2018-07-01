package main

import (
	"github.com/olivere/elastic"
	"encoding/json"
	"io/ioutil"
	"log"
	"context"
	"os"
	"path/filepath"
	"strings"
	"sort"
)

type ApkData struct {
	Apis		[]string
	Date		string
	FileSize	int64
	Intents		[]string
	Malicious	string
	Md5		string
	PackageName	string
	PackageVersion	string
	Sha1		string
	Sha256		string
	Strings		[]string
	Permissions	[]string
}

func GetPaths(dir string, Containing string) []string {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileInfo, errS := os.Stat(path)
		if errS != nil {
			log.Fatal(errS)
		}
		if strings.Contains(path, Containing) && fileInfo.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return fileList
}

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
    return list[i] < list[j]
}

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	jsonData := GetPaths("/iscsi/output/", ".json")
	sort.Sort(Alphabetic(jsonData))
	jsonData = jsonData[:250]
	for i, file := range jsonData {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var apkData ApkData
		err = json.Unmarshal(data, &apkData)
		if err != nil {
			log.Fatal(err)
		}
		apkData.Date = strings.Replace(apkData.Date, " ", "T", -1)
		_, err = client.Index().Index("apks").Type("_doc").Id(string(i)).BodyJson(apkData).Do(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(i)
	}
	_, err = client.Flush().Index("apks").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
