package parser

import (
	"strings"
)

type Parser struct {
	SP                int
	ArithmeticCounter int
	Static            int
	Temp              int
}

func NewParser() *Parser {
	return &Parser{
		SP:                256,
		ArithmeticCounter: 0,
		Static:            16,
		Temp:              5,
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

	// arithmetic
	return p.arithmetic(line)
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
