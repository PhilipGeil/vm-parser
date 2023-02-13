package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (p *Parser) push(line string) []string {
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the segment
	segment := splittedLines[1]
	switch segment {
	case "local":
		return p.pushLocal(line)
	case "argument":
		return p.pushArgument(line)
	case "this":
		return p.pushThis(line)
	case "that":
		return p.pushThat(line)
	case "constant":
		return p.pushConstant(line)
	case "static":
		return p.pushStatic(line)
	case "pointer":
		return p.pushPointer(line)
	case "temp":
		return p.pushTemp(line)
	default:
		log.Panicf("unknown segment: %s", segment)
		return nil
	}

}

func (p *Parser) pushLocal(line string) (lines []string) {
	// push local 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@LCL")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "A=D+A")
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushArgument(line string) (lines []string) {
	// push argument 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@ARG")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "A=D+A")
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushThis(line string) (lines []string) {
	// push this 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@THIS")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "A=D+A")
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushThat(line string) (lines []string) {
	// push that 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@THAT")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "A=D+A")
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushConstant(line string) (lines []string) {
	// push constant 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@"+index)
	lines = append(lines, "D=A")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushStatic(line string) (lines []string) {
	// push static 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	// convert index to int
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		log.Panicf("failed to convert index to int: %s", err)
	}

	fmt.Println("push", p.FileName)

	lines = append(lines, "@"+p.FileName+"."+strconv.Itoa(indexInt))
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}

func (p *Parser) pushPointer(line string) (lines []string) {
	// push pointer 0
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	if index == "0" {
		lines = append(lines, "@THIS")
		lines = append(lines, "D=M")
		lines = append(lines, "@SP")
		lines = append(lines, "A=M")
		lines = append(lines, "M=D")
		lines = append(lines, incSP()...)
	} else if index == "1" {
		lines = append(lines, "@THAT")
		lines = append(lines, "D=M")
		lines = append(lines, "@SP")
		lines = append(lines, "A=M")
		lines = append(lines, "M=D")
		lines = append(lines, incSP()...)
	} else {
		log.Panicf("unknown index: %s", index)
	}
	return
}

func (p *Parser) pushTemp(line string) (lines []string) {
	// push temp 0
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	// convert index to int
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		log.Panicf("failed to convert index to int: %s", err)
	}

	lines = append(lines, "@"+strconv.Itoa(indexInt+p.Temp))
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	return
}
