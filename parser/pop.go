package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (p *Parser) pop(line string) []string {
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the segment
	segment := splittedLines[1]
	switch segment {
	case "local":
		return p.popLocal(line)
	case "argument":
		return p.popArgument(line)
	case "this":
		return p.popThis(line)
	case "that":
		return p.popThat(line)
	case "static":
		return p.popStatic(line)
	case "temp":
		return p.popTemp(line)
	case "pointer":
		return p.popPointer(line)
	default:
		log.Panicf("unknown segment: %s", segment)
	}
	return nil
}

func (p *Parser) popLocal(line string) (lines []string) {
	// pop local 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@LCL")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "D=D+A")
	lines = append(lines, "@R13")
	lines = append(lines, "M=D")
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@R13")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popArgument(line string) (lines []string) {
	// pop argument 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@ARG")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "D=D+A")
	lines = append(lines, "@R13")
	lines = append(lines, "M=D")
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@R13")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popThis(line string) (lines []string) {
	// pop this 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@THIS")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "D=D+A")
	lines = append(lines, "@R13")
	lines = append(lines, "M=D")
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@R13")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popThat(line string) (lines []string) {
	// pop that 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	lines = append(lines, "@THAT")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+index)
	lines = append(lines, "D=D+A")
	lines = append(lines, "@R13")
	lines = append(lines, "M=D")
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@R13")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popStatic(line string) (lines []string) {
	// pop static 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	// convert index to int
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		log.Panicf("failed to convert index to int: %s", err)
	}

	fmt.Println("pop", p.FileName)
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@"+p.FileName+"."+strconv.Itoa(indexInt))
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popTemp(line string) (lines []string) {
	// pop temp 2
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	// convert index to int
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		log.Panicf("failed to convert index to int: %s", err)
	}

	lines = append(lines, "@5")
	lines = append(lines, "D=A")
	lines = append(lines, "@"+strconv.Itoa(indexInt))
	lines = append(lines, "D=D+A")
	lines = append(lines, "@R13")
	lines = append(lines, "M=D")
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "@R13")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")

	return
}

func (p *Parser) popPointer(line string) (lines []string) {
	// pop pointer 0
	// split the lines by spaces
	splittedLines := strings.Split(line, " ")

	// get the index
	index := splittedLines[2]

	if index == "0" {
		lines = append(lines, "@THIS")
		lines = append(lines, "D=A")
		lines = append(lines, "@R13")
		lines = append(lines, "M=D")
		lines = append(lines, decSP()...)
		lines = append(lines, "D=M")
		lines = append(lines, "@R13")
		lines = append(lines, "A=M")
		lines = append(lines, "M=D")
	} else if index == "1" {
		lines = append(lines, "@THAT")
		lines = append(lines, "D=A")
		lines = append(lines, "@R13")
		lines = append(lines, "M=D")
		lines = append(lines, decSP()...)
		lines = append(lines, "D=M")
		lines = append(lines, "@R13")
		lines = append(lines, "A=M")
		lines = append(lines, "M=D")
	} else {
		log.Panicf("invalid index for pop pointer: %s", index)
	}

	return
}
