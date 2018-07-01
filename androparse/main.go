package main

import (
	"flag"
	"fmt"
	"github.com/AndroParse/androparse/controller"
	"github.com/AndroParse/androparse/utils"
	"log"
	"os"
	"strings"
)

type ParserList []string

func (i *ParserList) String() string {
	return ""
}

func (i *ParserList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	configFlag := flag.String("config", "", "Location to YAML config file.")
	cleanFlag := flag.Bool("clean", false, "Move all apk files to their SHA256 value.")
	vtFlag := flag.Bool("vt", false, "Scan all files through Virus Total.")
	appendFlag := flag.Bool("append", false, "Append new feature extractor data.")
	var parsers ParserList
	flag.Var(&parsers, "parser", "List a plugin to run.")
	flag.Parse()

	if len(*configFlag) == 0 {
		printUsage()
		os.Exit(1)
	}

	config := utils.ReadConfig(*configFlag)
	config.Clean = *cleanFlag
	config.VtApiCheck = *vtFlag
	config.Append = *appendFlag
	config.Parsers = make([]string, 0)

	if len(parsers) != 0 {
		for _, parser := range parsers {
			config.Parsers = append(config.Parsers, strings.TrimSpace(strings.Trim(parser, ",")))
		}
	}

	log.Printf("apkDir: " + config.ApkDir)
	log.Printf("outputDir: " + config.OutputDir)
	log.Printf("codeDir: " + config.CodeDir)
	log.Printf("clean: %t", config.Clean)
	log.Printf("vt: %t", config.VtApiCheck)
	log.Printf("append: %t", config.Append)
	log.Printf("parsers: %v", config.Parsers)

	controller.Runner(config)
}

func printUsage() {
	fmt.Println(`
Syntax:
	>androparse -config <YAML config file> [-clean] [-vtFlag] [-append] [-parser <plugin name>]

Example:
	>androparse -config ./androparse.yaml
		`)
}
