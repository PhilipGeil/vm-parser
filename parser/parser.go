package parser

import (
	"strings"
)

type Parser struct {
	SP                int
	ArithmeticCounter int
	Static            int
	Temp              int
	CallCounter       int
	FileName          string
}

func NewParser() *Parser {
	return &Parser{
		SP:                256,
		ArithmeticCounter: 0,
		Static:            16,
		Temp:              5,
		CallCounter:       0,
	}
}

func (p *Parser) Parse(lines []string) (parsed []string) {
	for _, line := range lines {
		// remove comments
		line = strings.Split(line, "//")[0]
		// skip empty line
		if line == "" {
			continue
		}

		// parse line
		parsed = append(parsed, p.parseLine(line)...)
	}
	return
}

func (p *Parser) parseLine(line string) (lines []string) {
	// push local 0

	// Check if push
	if strings.HasPrefix(line, "push") {
		return p.push(line)
	}

	// Check if pop
	if strings.HasPrefix(line, "pop") {
		return p.pop(line)
	}

	// Check if program flow
	if p.isProgramFlow(line) {
		return p.programFlow(line)
	}

	// Check if subroutine
	if p.isSubroutine(line) {
		return p.subroutines(line)
	}

	// arithmetic
	return p.arithmetic(line)
}

func (p *Parser) Init() []string {
	return []string{
		"@256",
		"D=A",
		"@SP",
		"M=D",
		"@RETURN_LABEL0",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"@LCL",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"@ARG",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"@THIS",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"@THAT",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"@SP",
		"D=M",
		"@5",
		"D=D-A",
		"@0",
		"D=D-A",
		"@ARG",
		"M=D",
		"@SP",
		"D=M",
		"@LCL",
		"M=D",
		"@Sys.init",
		"0;JMP",
		"(RETURN_LABEL0)",
	}
}

func incSP() []string {
	return []string{
		"@SP",
		"M=M+1",
	}
}

func decSP() []string {
	return []string{
		"@SP",
		"AM=M-1",
	}
}
