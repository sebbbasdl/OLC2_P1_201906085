package expressions

import (
	"Server2/environment"
	"fmt"
)

type Count struct {
	Lin int
	Col int
	Id  string
}

func NewCount(lin int, col int, id string) Count {
	return Count{lin, col, id}
}

func (o Count) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	aux := environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
	sym := env.(environment.Environment).GetVariable(o.Id)

	if sym == aux {
		println("No existe el ID del vector")
	} else if sym.Tipo == environment.ARRAY {
		arrayValue, isArray := sym.Valor.([]interface{})
		if isArray {
			elementCount := len(arrayValue)
			fmt.Printf("El array tiene %d elementos.\n", elementCount)
			return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.INTEGER, Valor: elementCount}
		} else {
			println("Error! no es un arreglo")
		}
	} else {
		println("El ID no es un array")
	}

	return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: nil}
}
