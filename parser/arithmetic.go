package parser

import (
	"strconv"
	"strings"
)

func (p *Parser) arithmetic(line string) (lines []string) {
	splittedLines := strings.Split(line, " ")
	switch splittedLines[0] {
	case "add":
		lines = p.add()
	case "sub":
		lines = p.sub()
	case "neg":
		lines = p.neg()
	case "eq":
		lines = p.eq()
	case "gt":
		lines = p.gt()
	case "lt":
		lines = p.lt()
	case "and":
		lines = p.and()
	case "or":
		lines = p.or()
	case "not":
		lines = p.not()
	default:
		panic("unknown arithmetic command: " + line)
	}
	return
}

func (p *Parser) add() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "M=D+M")
	return
}

func (p *Parser) sub() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "M=M-D")
	return
}

func (p *Parser) neg() (lines []string) {
	lines = append(lines, "D=0")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=D-M")
	return
}

func (p *Parser) eq() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "D=M-D")
	lines = append(lines, "@FALSE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "D;JNE")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=-1")
	lines = append(lines, "@TRUE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "0;JMP")
	lines = append(lines, "(FALSE."+strconv.Itoa(p.ArithmeticCounter)+")")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=0")
	lines = append(lines, "(TRUE."+strconv.Itoa(p.ArithmeticCounter)+")")

	p.ArithmeticCounter++
	return
}

func (p *Parser) gt() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "D=M-D")
	lines = append(lines, "@FALSE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "D;JLE")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=-1")
	lines = append(lines, "@TRUE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "0;JMP")
	lines = append(lines, "(FALSE."+strconv.Itoa(p.ArithmeticCounter)+")")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=0")
	lines = append(lines, "(TRUE."+strconv.Itoa(p.ArithmeticCounter)+")")

	p.ArithmeticCounter++
	return
}

func (p *Parser) lt() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "D=M-D")
	lines = append(lines, "@FALSE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "D;JGE")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=-1")
	lines = append(lines, "@TRUE."+strconv.Itoa(p.ArithmeticCounter)+"")
	lines = append(lines, "0;JMP")
	lines = append(lines, "(FALSE."+strconv.Itoa(p.ArithmeticCounter)+")")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=0")
	lines = append(lines, "(TRUE."+strconv.Itoa(p.ArithmeticCounter)+")")

	p.ArithmeticCounter++
	return
}

func (p *Parser) and() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "M=M&D")
	return
}

func (p *Parser) or() (lines []string) {
	lines = append(lines, decSP()...)
	lines = append(lines, "D=M")
	lines = append(lines, "A=A-1")
	lines = append(lines, "M=M|D")
	return
}

func (p *Parser) not() (lines []string) {
	lines = append(lines, "@SP")
	lines = append(lines, "A=M-1")
	lines = append(lines, "M=!M")
	return
}
