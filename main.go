package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PhilipGeil/vm-parser/parser"
)

func main() {
	// get path from arg
	args := os.Args
	if len(args) < 2 {
		log.Panicf("no path provided")
	}
	path := args[1]

	// check if path is file or directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Panicf("path error: %v", err)
	}

	var lines []string
	var outputFileName string
	parser := parser.NewParser()

	lines = append(lines, parser.Init()...)

	if fileInfo.IsDir() {
		// load files from directory
		files, err := loadFiles(path)
		if err != nil {
			log.Panicf("load files error: %v", err)
		}
		outputFileName = path + "/" + path
		for _, file := range files {
			fileLines, err := loadFile(path + "/" + file)
			if err != nil {
				log.Panicf("load file error: %v", err)
			}
			// lines = append(lines, fileLines...)
			parser.FileName = strings.TrimSuffix(file, filepath.Ext(file))
			lines = append(lines, parser.Parse(fileLines)...)
		}
		writeToFile(lines, outputFileName)
	} else {
		// load file
		lines, err = loadFile(path)
		if err != nil {
			log.Panicf("load file error: %v", err)
		}
		fileName := strings.TrimSuffix(path, filepath.Ext(path))
		outputFileName = fileName + "/" + fileName
		parser.FileName = fileName
		writeToFile(parser.Parse(lines), outputFileName)
	}

}

// load files from directory
func loadFiles(path string) ([]string, error) {
	var lines []string

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".vm" {
			lines = append(lines, file.Name())
		}
	}

	return lines, nil
}

// load file
func loadFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\r\n")
	return lines, nil
}

func writeToFile(lines []string, name string) {
	// Open file using os.Create()
	file, err := os.Create(name + ".asm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write each line to the file
	for _, line := range lines {
		_, err := writer.WriteString(line + "\r\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the buffer to ensure everything is written to the file
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
