package instructions

import (
	"Server2/environment"
	"fmt"
	"strconv"
)

type RemoveLast struct {
	Lin int
	Col int
	Id  string
}

func NewRemoveLast(lin int, col int, id string) RemoveLast {
	return RemoveLast{lin, col, id}
}

func (p RemoveLast) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	aux := environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
	sym := env.(environment.Environment).GetVariable(p.Id)

	if sym == aux {
		println("No existe el ID del vector")
	} else if sym.Tipo == environment.ARRAY {
		// Verificar si el array no está vacío antes de eliminar el último elemento
		if len(sym.Valor.([]interface{})) > 0 {
			// Eliminar el último elemento del array
			sym.Valor = sym.Valor.([]interface{})[:len(sym.Valor.([]interface{}))-1]
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
			println("El array está vacío, no se puede eliminar el último elemento.")
		}
	} else {
		println("El ID no es un array")
	}
	return nil
}
