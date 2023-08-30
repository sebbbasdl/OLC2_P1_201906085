package expressions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
	"strconv"
)

type Conversion struct {
	Expr interfaces.Expression
	Tipo environment.TipoExpresion
}

func NewConversion(Expr interfaces.Expression, Tipo environment.TipoExpresion) Conversion {
	exp := Conversion{Expr: Expr, Tipo: Tipo}
	return exp
}

func (o Conversion) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	if o.Tipo == environment.INTEGER {
		if o.Expr.Ejecutar(ast, env).Tipo == environment.STRING {
			stringValue := o.Expr.Ejecutar(ast, env).Valor.(string) // Obtén el valor como string

			intValue, err := strconv.Atoi(stringValue) // Convierte el string a int
			if err != nil {
				// Maneja el error si la conversión falla
				fmt.Println("Error al convertir a int:", err)
				return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: nil}
			}

			return environment.Symbol{Lin: 0, Col: 0, Tipo: o.Tipo, Valor: intValue}
		} else if o.Expr.Ejecutar(ast, env).Tipo == environment.FLOAT {
			floatValue := o.Expr.Ejecutar(ast, env).Valor.(float64) // Obtén el valor como float

			intValue := int(floatValue) // Convierte el float a int directamente

			return environment.Symbol{Lin: 0, Col: 0, Tipo: o.Tipo, Valor: intValue}
		}
	} else if o.Tipo == environment.STRING {
		if o.Expr.Ejecutar(ast, env).Tipo == environment.INTEGER {
			intValue := o.Expr.Ejecutar(ast, env).Valor.(int) // Obtén el valor como int

			stringValue := strconv.Itoa(intValue) // Convierte el int a string

			return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.STRING, Valor: stringValue}
		}else if o.Expr.Ejecutar(ast, env).Tipo == environment.FLOAT {
			floatValue := o.Expr.Ejecutar(ast, env).Valor.(float64) // Obtén el valor como float
		
			stringValue := strconv.FormatFloat(floatValue, 'f', -1, 64) // Convierte el float a string
		
			return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.STRING, Valor: stringValue}
		}
		

	} else if o.Tipo == environment.FLOAT {
		if o.Expr.Ejecutar(ast, env).Tipo == environment.INTEGER {
			intValue := o.Expr.Ejecutar(ast, env).Valor.(int) // Obtén el valor como int
		
			floatValue := float64(intValue) // Convierte el int a float
		
			return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.FLOAT, Valor: floatValue}
		}else if o.Expr.Ejecutar(ast, env).Tipo == environment.STRING {
			stringValue := o.Expr.Ejecutar(ast, env).Valor.(string) // Obtén el valor como string
		
			floatValue, err := strconv.ParseFloat(stringValue, 64) // Convierte el string a float64
			if err != nil {
				// Maneja el error si la conversión falla
				fmt.Println("Error al convertir a float:", err)
				return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: nil}
			}
		
			return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.FLOAT, Valor: floatValue}
		}
		
		
	}

	var result interface{}
	return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: result}
}
