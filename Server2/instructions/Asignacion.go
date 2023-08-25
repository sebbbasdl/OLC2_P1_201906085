package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"strconv"
)

type Asignacion struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func esNumero(valor interface{}) (bool, int) {
	if str, ok := valor.(string); ok {
		numero, err := strconv.Atoi(str)
		if err == nil {
			return true, numero
		}
	}
	return false, 0
}

func NewAsignacion(lin int, col int, id string, val interfaces.Expression) Asignacion {
	instr := Asignacion{lin, col, id, val}
	return instr
}

func (p Asignacion) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	var result environment.Symbol
	result = p.Expresion.Ejecutar(ast, env)

	//verificacion de tipo

	if p.Expresion.Ejecutar(ast, env).Tipo == result.Tipo {
		//cambio de valor en la tabla de simbolos
		env.(environment.Environment).SetVariable(p.Id, result)

	}

	return nil
}
