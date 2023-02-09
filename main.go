package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PhilipGeil/vm-parser/parser"
)

func main() {
	filePath := "StackTest.vm"
	// load file
	lines, err := loadFile(filePath)
	if err != nil {
		log.Panicf("load file error: %v", err)
	}

	parser := parser.NewParser()
	writeToFile(parser.Parse(lines))
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

func writeToFile(lines []string) {
	// Open file using os.Create()
	file, err := os.Create("StackTest.asm")
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
