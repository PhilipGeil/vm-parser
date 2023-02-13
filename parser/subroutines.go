package parser

import (
	"log"
	"strconv"
	"strings"
)

func (p *Parser) subroutines(line string) (lines []string) {
	splitLine := strings.Split(line, " ")
	switch splitLine[0] {
	case "call":
		lines = p.call(splitLine[1:])
	case "function":
		lines = p.function(splitLine[1:])
	case "return":
		lines = p.returnFunc(splitLine[1:])
	default:
		log.Panicln("unknown subroutine: " + splitLine[0])
	}

	return
}

func (p *Parser) isSubroutine(line string) bool {
	return strings.HasPrefix(line, "call") || strings.HasPrefix(line, "function") || strings.HasPrefix(line, "return")
}

func (p *Parser) call(args []string) (lines []string) {
	slots := []string{"LCL", "ARG", "THIS", "THAT"}

	lines = append(lines, "@RETURN_ADDRESS"+strconv.Itoa(p.CallCounter))
	lines = append(lines, "D=A")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)
	for _, slot := range slots {
		lines = append(lines, p.pushSlot(slot)...)
	}
	lines = append(lines, "@SP")
	lines = append(lines, "D=M")
	lines = append(lines, "@5")
	lines = append(lines, "D=D-A")
	lines = append(lines, "@"+args[1])
	lines = append(lines, "D=D-A")
	lines = append(lines, "@ARG")
	lines = append(lines, "M=D")
	lines = append(lines, "@SP")
	lines = append(lines, "D=M")
	lines = append(lines, "@LCL")
	lines = append(lines, "M=D")
	lines = append(lines, "@"+args[0])
	lines = append(lines, "0;JMP")
	lines = append(lines, "(RETURN_ADDRESS"+strconv.Itoa(p.CallCounter)+")")

	p.CallCounter++
	return
}

func (p *Parser) function(args []string) (lines []string) {
	lines = append(lines, "("+args[0]+")")

	// Convert nArgs to int
	nArgs, err := strconv.Atoi(args[1])
	if err != nil {
		log.Panicln("function: " + err.Error())
	}
	for i := 0; i < nArgs; i++ {
		lines = append(lines, p.push("push constant 0")...)
	}
	return
}

func (p *Parser) returnFunc(args []string) (lines []string) {
	slots := []string{"THAT", "THIS", "ARG", "LCL"}

	lines = append(lines, "@LCL")
	lines = append(lines, "D=M")
	lines = append(lines, "@R11")
	lines = append(lines, "M=D")
	lines = append(lines, "@5")
	lines = append(lines, "A=D-A")
	lines = append(lines, "D=M")
	lines = append(lines, "@R12")
	lines = append(lines, "M=D")
	lines = append(lines, p.pop("pop argument 0")...)
	lines = append(lines, "@ARG")
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "M=D+1")

	// Set slots to endframe - n
	for _, slot := range slots {
		lines = append(lines, p.setSlotToFrameMinusN(slot)...)
	}

	lines = append(lines, "@R12")
	lines = append(lines, "A=M")
	lines = append(lines, "0;JMP")
	return
}

// Set slot to endframe - n
func (p *Parser) setSlotToFrameMinusN(slot string) (lines []string) {
	lines = append(lines, "@R11")
	lines = append(lines, "D=M-1")
	lines = append(lines, "AM=D")
	lines = append(lines, "D=M")
	lines = append(lines, "@"+slot)
	lines = append(lines, "M=D")

	return
}

func (p *Parser) pushSlot(slot string) (lines []string) {
	lines = append(lines, "@"+slot)
	lines = append(lines, "D=M")
	lines = append(lines, "@SP")
	lines = append(lines, "A=M")
	lines = append(lines, "M=D")
	lines = append(lines, incSP()...)

	return
}
