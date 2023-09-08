package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
)

type Print struct {
	Lin    int
	Col    int
	Value  interface{}
	Value2 interface{}
}

func NewPrint(lin int, col int, val interface{}, val2 interface{}) Print {
	return Print{lin, col, val, val2}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	if p.Value2 != nil {
		valueToPrint := p.Value.(interfaces.Expression).Ejecutar(ast, env)
		valueToPrint2 := p.Value2.(interfaces.Expression).Ejecutar(ast, env)
		consoleOut := fmt.Sprintf("%v"+"%v", valueToPrint.Valor, valueToPrint2.Valor)
		ast.SetPrint(consoleOut + "\n")
	} else {
		valueToPrint := p.Value.(interfaces.Expression).Ejecutar(ast, env)
		consoleOut := fmt.Sprintf("%v", valueToPrint.Valor)
		ast.SetPrint(consoleOut + "\n")
	}

	return nil
}
