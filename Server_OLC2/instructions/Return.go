package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Return struct {
	Lin int
	Col int
	Exp interfaces.Expression
}

func NewReturn(lin int, col int, exp interfaces.Expression) Return {
	returninstr := Return{lin, col, exp}
	return returninstr
}

func (p Return) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	ast.ReturnBool = true
	ast.ReturnLine = p.Lin
	ast.ReturnCol = p.Col
	result := p.Exp.Ejecutar(ast, env)

	// Verificar si el valor es de tipo int
	if intValue, ok := result.Valor.(int); ok {
		ast.ReturnVal = intValue
		ast.ReturnTipo = result.Tipo
	}
	println(ast.ReturnVal)
	return nil

}
