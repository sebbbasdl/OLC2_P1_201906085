package expressions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Array struct {
	Lin     int
	Col     int
	ListExp []interface{}
}

func NewArray(lin int, col int, list []interface{}) Array {
	exp := Array{lin, col, list}
	return exp
}

func (p Array) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	int_flag := false
	str_flag := false
	float_flag := false
	arr_flag := false
	bool_flag := false
	var tempExp []interface{}
	for _, s := range p.ListExp {
		if s.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.INTEGER {
			int_flag = true

			//tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		}

		if s.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.STRING {

			str_flag = true

			//tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		}
		if s.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.FLOAT {

			float_flag = true

			//tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		}

		if s.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.ARRAY {
			arr_flag = true
			//tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		}

		if s.(interfaces.Expression).Ejecutar(ast, env).Tipo == environment.BOOLEAN {
			bool_flag = true
			//tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		}

		if int_flag == true && str_flag == false && float_flag == false && arr_flag == false && bool_flag == false {
			tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		} else if int_flag == false && str_flag == true && float_flag == false && arr_flag == false && bool_flag == false {
			tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		} else if int_flag == false && str_flag == false && float_flag == true && arr_flag == false && bool_flag == false {
			tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		} else if int_flag == false && str_flag == false && float_flag == false && arr_flag == true && bool_flag == false {
			tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		} else if int_flag == false && str_flag == false && float_flag == false && arr_flag == false && bool_flag == true {
			tempExp = append(tempExp, s.(interfaces.Expression).Ejecutar(ast, env))
		} else {
			ast.SetError("Error, los tipos de datos son incorrectos")
			break
		}

	}

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Tipo:  environment.ARRAY,
		Valor: tempExp,
	}
}
