package main

import (
	"github.com/AndroParse/androparse/utils"
	"time"
)

func NeedLock() bool { return false }

func GetKey() string { return "Date" }

func GetValue(path string, config utils.ConfigData) (interface{}, error) {
	t := time.Now().UTC()
	return t.Format("2006-01-02 15:04:05"), nil
}
