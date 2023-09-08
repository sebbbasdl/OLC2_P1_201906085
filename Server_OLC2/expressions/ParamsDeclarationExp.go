package expressions

import (
	"Server2/environment"
)

type ParamsDeclarationFunc struct {
	Lin  int
	Col  int
	Id   string
	Tipo environment.TipoExpresion
}

func NewParamsDeclarationFunc(lin int, col int, id string, tipo environment.TipoExpresion) ParamsDeclarationFunc {
	instr := ParamsDeclarationFunc{lin, col, id, tipo}
	return instr
}

func (p ParamsDeclarationFunc) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result environment.Symbol

	result = environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Tipo:  p.Tipo,
		Valor: 0,
	}

	env.(environment.Environment).SaveVariable(p.Id, result)

	return result
}
