package compiler

import (
	"go-hacks/code"
	"go/ast"
)

type Compiler struct {
	instructions code.Instructions
	constants []ast.Object
}

