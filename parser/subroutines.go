package parser

import (
	"log"
	"strings"
)

func (p *Parser) subroutines(line string) (lines []string) {
	splitLine := strings.Split(line, " ")
	switch splitLine[0] {
	case "call":
		lines = p.call()
	case "function":
		lines = p.function()
	case "return":
		lines = p.returnFunc()
	default:
		log.Panicln("unknown subroutine: " + splitLine[0])
	}

	return
}

func (p *Parser) call() (lines []string) {
	return
}

func (p *Parser) function() (lines []string) {
	return
}

func (p *Parser) returnFunc() (lines []string) {
	return
}
