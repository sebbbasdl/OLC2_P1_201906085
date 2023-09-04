package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
	"strconv"
)

type Declaration struct {
	Lin       int
	Col       int
	Id        string
	Tipo      environment.TipoExpresion
	Expresion interfaces.Expression
	Global    bool
}

func NewDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, val interfaces.Expression, global bool) Declaration {
	instr := Declaration{lin, col, id, tipo, val, global}
	return instr
}

func (p Declaration) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	//Traer simbolo
	var result environment.Symbol
	result = p.Expresion.Ejecutar(ast, env)
	//validar tipos
	if result.Tipo == environment.ARRAY {
		if p.ArrayValidation(result) {
			env.(environment.Environment).SaveVariable(p.Id, result)

			var valorString string
			if val, ok := result.Valor.(string); ok {
				valorString = val
			} else {
				valorString = fmt.Sprint(result.Valor)
			}
			valorString = arrayToString(result.Valor.([]interface{}))
			datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), valorString, p.Id}
			ast.Tabla = append(ast.Tabla, datos)
		} else {
			ast.SetError("La estructura del array es incorrecta")
		}

	} else if result.Tipo == p.Tipo {
		if p.Tipo == environment.INTEGER {

			env.(environment.Environment).SaveVariable(p.Id, result)
			if p.Global == true {
				ast.Constantes = append(ast.Constantes, p.Id)
			}

			datos := []string{"1", "Variable", "Int", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(int)), p.Id}
			ast.Tabla = append(ast.Tabla, datos)

		} else if p.Tipo == environment.FLOAT {
			env.(environment.Environment).SaveVariable(p.Id, result)
			datos := []string{"1", "Variable", "Float", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(float64)), p.Id}
			ast.Tabla = append(ast.Tabla, datos)
		} else if p.Tipo == environment.STRING {
			env.(environment.Environment).SaveVariable(p.Id, result)
			datos := []string{"1", "Variable", "String", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(string)), p.Id}
			ast.Tabla = append(ast.Tabla, datos)
		} else if p.Tipo == environment.BOOLEAN {
			env.(environment.Environment).SaveVariable(p.Id, result)
			datos := []string{"1", "Variable", "Bool", env.(environment.Environment).Id, strconv.Itoa(result.Col), strconv.Itoa(result.Lin), fmt.Sprint(result.Valor.(bool)), p.Id}
			ast.Tabla = append(ast.Tabla, datos)
		}

	} else {
		ast.SetError("Los tipos de datos son incorrectos")
	}
	return nil
}

func (p Declaration) ArrayValidation(result environment.Symbol) bool {
	return true
}
