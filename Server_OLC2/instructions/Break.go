package instructions

import (
	"Server2/environment"
)

type Break struct {
	Lin int
	Col int
}

func NewBreak(lin int, col int) Break {
	breakinstr := Break{lin, col}
	return breakinstr
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	ast.Breakbool = true
	return nil

}
