package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
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

		if contieneString(ast.Constantes, p.Id) == false {
			//cambio de valor en la tabla de simbolos
			env.(environment.Environment).SetVariable(p.Id, result)
			datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(int)), p.Id}
			modificarValorEnTabla(ast.Tabla, datos)

		} else {
			println("Es una variable constante, no se puede modificar")
		}

	}

	return nil
}

func modificarValorEnTabla(tabla [][]string, dato []string) {

	for i := 0; i < len(tabla); i++ {
		if tabla[i][7] == dato[7] {
			dato[3] = tabla[i][3]
			tabla[i] = dato
		}
	}

}

func contieneString(arr []string, elemento string) bool {
	for _, valor := range arr {
		if valor == elemento {
			return true
		}
	}
	return false
}
