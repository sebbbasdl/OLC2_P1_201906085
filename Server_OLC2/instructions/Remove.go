package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
	"strconv"
)

type Remove struct {
	Lin   int
	Col   int
	Id    string
	Value interface{}
}

func NewRemove(lin int, col int, id string, val interface{}) Remove {
	return Remove{lin, col, id, val}
}

func (p Remove) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	result = p.Value.(interfaces.Expression).Ejecutar(ast, env)

	if p.Value.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.INTEGER {
		aux := environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
		sym := env.(environment.Environment).GetVariable(p.Id)

		if sym == aux {
			println("No existe el ID del vector")
		} else if sym.Tipo == environment.ARRAY {
			indexToRemove := result.Valor.(int)
			arrayValue, isArray := sym.Valor.([]interface{})

			if isArray {
				if indexToRemove >= 0 && indexToRemove < len(arrayValue) {
					// Eliminar el elemento en la posición especificada
					sym.Valor = append(arrayValue[:indexToRemove], arrayValue[indexToRemove+1:]...)
					fmt.Println(sym.Valor)
					// Actualizar la variable en el entorno
					env.(environment.Environment).SetVariable(p.Id, sym)

					// Resto del código aquí...

					var valorString string
					if val, ok := sym.Valor.(string); ok {
						valorString = val
					} else {
						valorString = fmt.Sprint(sym.Valor)
					}

					valorString = arrayToString(sym.Valor.([]interface{}))
					datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(p.Col), strconv.Itoa(p.Lin), valorString, p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				} else {
					println("La posición especificada está fuera de los límites del array.")
				}
			} else {
				println("El ID no es un array")
			}
		} else {
			println("El ID no es un array")
		}
	}

	return nil
}
