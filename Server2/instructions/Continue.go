package instructions

import (
	"Server2/environment"
)

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	continstr := Continue{lin, col}
	return continstr
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	ast.ContinueBool = true
	return nil

}
