package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type Append struct {
	Lin   int
	Col   int
	Id    string
	Value interface{}
}

func NewAppend(lin int, col int, id string, val interface{}) Append {
	return Append{lin, col, id, val}
}

func (p Append) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	result = p.Value.(interfaces.Expression).Ejecutar(ast, env)

	if p.Value.(interfaces.Expression).Ejecutar(ast, env).Tipo == result.Tipo {
		aux := environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
		sym := env.(environment.Environment).GetVariable(p.Id)

		if sym == aux {
			println("No existe el ID del vector")
		} else if sym.Tipo == environment.ARRAY {
			// Agregar el nuevo valor al array
			sym.Valor = append(sym.Valor.([]interface{}), result)
			fmt.Println(sym.Valor)
			// Actualizar la variable en el entorno
			env.(environment.Environment).SetVariable(p.Id, sym)
			var valorString string
			if val, ok := sym.Valor.(string); ok {
				valorString = val
			} else {
				valorString = fmt.Sprint(sym.Valor)
			}

			valorString = arrayToString(sym.Valor.([]interface{}))
			datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), valorString, p.Id}
			modificarValorEnTabla(ast.Tabla, datos)

		} else {
			println("El ID no es un array")
		}
	}

	return nil
}

func arrayToString(arr []interface{}) string {
	values := []string{}
	for _, v := range arr {
		if elem, ok := v.(environment.Symbol); ok {
			values = append(values, fmt.Sprintf("%v", elem.Valor))
		}
	}
	return "[" + strings.Join(values, ",") + "]"
}

func retornaTipo(arr []interface{}) int {
	var tipo int // Declarar la variable tipo fuera del bloque if
	for _, v := range arr {
		if elem, ok := v.(environment.Symbol); ok {
			tipo = int(elem.Tipo)
			return tipo
		}
	}
	return tipo // Devolver el valor de tipo fuera del bucle
}
