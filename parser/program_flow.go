package parser

import (
	"log"
	"strings"
)

func (p *Parser) isProgramFlow(line string) bool {
	splitLine := strings.Split(line, " ")
	switch splitLine[0] {
	case "label", "goto", "if-goto":
		return true
	}
	return false
}

func (p *Parser) programFlow(line string) (lines []string) {
	splitLine := strings.Split(line, " ")
	switch splitLine[0] {
	case "label":
		lines = p.label(splitLine[1])
	case "goto":
		lines = p.gotoFunc(splitLine[1])
	case "if-goto":
		lines = p.ifGoto(splitLine[1])
	default:
		log.Panicln("unknown program flow: " + splitLine[0])
	}
	return
}

func (p *Parser) label(name string) (lines []string) {
	lines = append(lines, "("+name+")")
	return
}

func (p *Parser) gotoFunc(name string) (lines []string) {
	lines = append(lines, "@"+name)
	lines = append(lines, "0;JMP")
	return
}

func (p *Parser) ifGoto(name string) (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "@"+name)
	lines = append(lines, "D;JNE")
	return
}
