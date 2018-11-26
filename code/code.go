package code

import (
	"fmt"
	"encoding/binary"
)

type Instructions []byte

type OpCode byte

const (
	OpConstant OpCode = iota
)

type Definition struct {
	Name string
	OperandWidths []int
}

var definitions = map[OpCode]*Definition {
	OpConstant: {"OpConstant",  []int{2}},
}

func LookUp(op byte) (*Definition, error){
	def, ok := definitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op OpCode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)

	instruction[0] = byte(op)

	offset := 1

	for i, o := range operands {
		w := def.OperandWidths[i]
		switch w {
			case 2: binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += w
	}

	return instruction
}








