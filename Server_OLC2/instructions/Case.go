package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Case struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewCase(lin int, col int, expr1 interfaces.Expression, bloque []interface{}) Case {
	caseInstr := Case{lin, col, expr1, bloque}
	return caseInstr
}

func (p Case) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	if ast.SwitchBool == true {
		for _, inst := range p.Bloque {
			inst.(interfaces.Instruction).Ejecutar(ast, env)
		}
	}
	return nil

}
