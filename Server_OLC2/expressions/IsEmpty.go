package expressions

import (
	"Server2/environment"
	"fmt"
)

type IsEmpty struct {
	Lin int
	Col int
	Id  string
}

func NewIsEmpty(lin int, col int, id string) IsEmpty {
	return IsEmpty{lin, col, id}
}

func (o IsEmpty) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	aux := environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
	sym := env.(environment.Environment).GetVariable(o.Id)

	if sym == aux {
		println("No existe el ID del vector")
	} else if sym.Tipo == environment.ARRAY {
		arrayValue, isArray := sym.Valor.([]interface{})
		if isArray {
			if len(arrayValue) == 0 {
				fmt.Println("El array está vacío.")
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: true}
			} else {
				fmt.Println("El array no está vacío.")
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: false}
			}
		} else {
			println("Error! no es un arreglo")
		}
	} else {
		println("El ID no es un array")
	}

	return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: nil}
}
