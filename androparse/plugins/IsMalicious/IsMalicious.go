package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/AndroParse/androparse/utils"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type VirusTotal struct {
	apikey string
}

type VirusTotalResponse struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"verbose_msg"`
}

type ScanResponse struct {
	VirusTotalResponse

	ScanId    string `json:"scan_id"`
	Sha1      string `json:"sha1"`
	Resource  string `json:"resource"`
	Sha256    string `json:"sha256"`
	Permalink string `json:"permalink"`
	Md5       string `json:"md5"`
}

type FileScan struct {
	Detected bool   `json:"detected"`
	Version  string `json:"version"`
	Result   string `json:"result"`
	Update   string `json:"update"`
}

type ReportResponse struct {
	VirusTotalResponse
	Resource  string              `json:"resource"`
	ScanId    string              `json:"scan_id"`
	Sha1      string              `json:"sha1"`
	Sha256    string              `json:"sha256"`
	Md5       string              `json:"md5"`
	Scandate  string              `json:"scan_date"`
	Positives int                 `json:"positives"`
	Total     int                 `json:"total"`
	Permalink string              `json:"permalink"`
	Scans     map[string]FileScan `json:"scans"`
}

func NeedLock() bool { return false }

func GetKey() string { return "Malicious" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	malicious := ""
	if !config.VtApiCheck {
		return fallbackMalicious(path), nil
	} else {
		vti, err := NewVirusTotal(config.VtApiKey)
		utils.Check(err)
		_, hash := filepath.Split(path)
		hash = hash[:len(hash)-4]
		rr, err := vti.checkReport(hash)
		if err != nil || (rr.Md5 == "" && rr.Positives == 0) {
			log.Printf("Hash not found: " + hash + " scanning now...")
			err := vti.scanApk(path)
			if err != nil {
				return fallbackMalicious(path), err
			}
			time.Sleep(2 * time.Minute)
			rr, err = vti.checkReport(hash)
			if err != nil {
				return fallbackMalicious(path), err
			}
		} else {
			if rr.Positives > 4 {
				malicious = "true"
			} else {
				malicious = "false"
			}
		}
	}
	return malicious, nil
}

func fallbackMalicious(apkPath string) string {
	if strings.Contains(apkPath, "benign") {
		return "false"
	} else if strings.Contains(apkPath, "malware") {
		return "true"
	} else {
		return "unknown"
	}
}

func (vt *VirusTotal) scanApk(apkPath string) error {
	file, err := os.Open(apkPath)
	utils.Check(err)
	_, errApi := vt.Scan(apkPath, file)
	for errApi != nil {
		if errApi != nil && errApi.Error() == "API Limit Reached" {
			log.Printf("API limit reached, sleeping for 1 minute...")
			time.Sleep(1*time.Minute + 5*time.Second)
			_, errApi = vt.Scan(apkPath, file)
		} else {
			return errors.New("Warning: Could not parse VirusTotal scan output of " + apkPath)
		}
	}
	file.Close()
	return nil
}

func (vt *VirusTotal) checkReport(hash string) (*ReportResponse, error) {
	var rr *ReportResponse
	rr, errApi := vt.Report(hash)
	for errApi != nil {
		if errApi != nil && errApi.Error() == "API Limit Reached" {
			log.Printf("API limit reached, sleeping for 1 minute...")
			time.Sleep(1*time.Minute + 5*time.Second)
			rr, errApi = vt.Report(hash)
		} else {
			return nil, errors.New("Warning: Could not parse VirusTotal report output of " + hash)
		}
	}
	return rr, nil
}

func NewVirusTotal(apikey string) (*VirusTotal, error) {
	vt := &VirusTotal{apikey: apikey}
	return vt, nil
}

func (vt *VirusTotal) Report(resource string) (*ReportResponse, error) {
	u, err := url.Parse("https://www.virustotal.com/vtapi/v2/file/report")

	params := url.Values{"apikey": {vt.apikey}, "resource": {resource}}

	resp, err := http.PostForm(u.String(), params)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 204 {
		return nil, errors.New("API Limit Reached")
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var reportResponse = &ReportResponse{}

	err = json.Unmarshal(contents, &reportResponse)

	return reportResponse, err
}

func (vt *VirusTotal) Scan(path string, file io.Reader) (*ScanResponse, error) {
	params := map[string]string{
		"apikey": vt.apikey,
	}

	request, err := newfileUploadRequest("http://www.virustotal.com/vtapi/v2/file/scan", params, path, file)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 204 {
		return nil, errors.New("API Limit Reached")
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var scanResponse = &ScanResponse{}
	err = json.Unmarshal(contents, &scanResponse)

	return scanResponse, err
}

func newfileUploadRequest(uri string, params map[string]string, path string, file io.Reader) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
