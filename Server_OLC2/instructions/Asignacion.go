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
	Op        string
	
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

func NewAsignacion(lin int, col int, id string, val interfaces.Expression, op string) Asignacion {
	instr := Asignacion{lin, col, id, val, op}
	return instr
}

func (p Asignacion) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	var result environment.Symbol
	result = p.Expresion.Ejecutar(ast, env)

	//verificacion de tipo
	if p.Op == "=" {
		if p.Expresion.Ejecutar(ast, env).Tipo == result.Tipo {

			if contieneString(ast.Constantes, p.Id) == false {
				//cambio de valor en la tabla de simbolos
				if result.Tipo == environment.INTEGER {
					env.(environment.Environment).SetVariable(p.Id, result)
					datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(int)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)

				} else if result.Tipo == environment.FLOAT {
					env.(environment.Environment).SetVariable(p.Id, result)
					datos := []string{"1", "Variable", "Float", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(float64)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				} else if result.Tipo == environment.STRING {
					env.(environment.Environment).SetVariable(p.Id, result)
					datos := []string{"1", "Variable", "String", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(string)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				} else if result.Tipo == environment.BOOLEAN {
					env.(environment.Environment).SetVariable(p.Id, result)
					datos := []string{"1", "Variable", "Bool", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(bool)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				}

			} else {
				ast.SetError("Es una variable constante, no se puede modificar")
			}

		}

		return nil

	} else if p.Op == "+=" {
		
		if p.Expresion.Ejecutar(ast, env).Tipo == result.Tipo {
		
			if contieneString(ast.Constantes, p.Id) == false {
				//cambio de valor en la tabla de simbolos
		
				sym := env.(environment.Environment).GetVariable(p.Id)
		
				if result.Tipo == environment.INTEGER {
					aux := environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: result.Tipo, Valor: sym.Valor.(int) + result.Valor.(int)}
					env.(environment.Environment).SetVariable(p.Id, aux)
					datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(int)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)

				} else if result.Tipo == environment.FLOAT {
					aux := environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: result.Tipo, Valor: sym.Valor.(float64) + result.Valor.(float64)}
					env.(environment.Environment).SetVariable(p.Id, aux)
					datos := []string{"1", "Variable", "Float", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(float64)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				} else if result.Tipo == environment.STRING {
					aux := environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: result.Tipo, Valor: sym.Valor.(string) + result.Valor.(string)}
					env.(environment.Environment).SetVariable(p.Id, aux)
					datos := []string{"1", "Variable", "String", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(string)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				} else if result.Tipo == environment.BOOLEAN {
					env.(environment.Environment).SetVariable(p.Id, result)
					datos := []string{"1", "Variable", "Bool", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(bool)), p.Id}
					modificarValorEnTabla(ast.Tabla, datos)
				}

			} else {
				ast.SetError("Es una variable constante, no se puede modificar")
			}

		}

		return nil

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
