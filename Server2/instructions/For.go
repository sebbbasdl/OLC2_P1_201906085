package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
	"fmt"
)

type Fors struct {
	Lin        int
	Col        int
	Id         string
	Expresion  interfaces.Expression
	Expresion2 interfaces.Expression
	Str        string
	Bloque     []interface{}
}

func NewFor(lin int, col int, id string, expr1 interfaces.Expression, expr2 interfaces.Expression, str string, bloque []interface{}) Fors {
	forInstr := Fors{lin, col, id, expr1, expr2, str, bloque}
	return forInstr
}

func (p Fors) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	if p.Expresion != nil && p.Expresion2 != nil {
		var num1 environment.Symbol
		num1 = p.Expresion.Ejecutar(ast, env)
		var num2 environment.Symbol
		num2 = p.Expresion2.Ejecutar(ast, env)
		env.(environment.Environment).SaveVariable(p.Id, num1)

		num1Value, ok1 := num1.Valor.(int)
		num2Value, ok2 := num2.Valor.(int)

		if !ok1 || !ok2 {
			fmt.Println("Error al obtener valores numéricos de environment.Symbol")
			return nil
		}

		var forEnv environment.Environment
		forEnv = environment.NewEnvironment(env.(environment.Environment), "For")

		for i := num1Value; i <= num2Value; i++ {

			iSymbol := environment.Symbol{Valor: i, Tipo: environment.INTEGER} // Crear un nuevo Symbol con el valor 'i'
			env.(environment.Environment).SetVariable(p.Id, iSymbol)
			//ejecución
			for _, inst := range p.Bloque {
				inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
			}

		}

	} else {

		var str1 string
		str1 = p.Str
		//Es string entonces
		if str1[0] == '"' {
			content := str1[1 : len(str1)-1]
			iSymbol0 := environment.Symbol{Valor: string(content), Tipo: environment.STRING}
			env.(environment.Environment).SaveVariable(p.Id, iSymbol0)
			if content != "nil" {
				var forEnv environment.Environment
				forEnv = environment.NewEnvironment(env.(environment.Environment), "For")
				for _, c := range content {
					iSymbol := environment.Symbol{Valor: string(c), Tipo: environment.STRING} // Crear un nuevo Symbol con el valor 'i'
					env.(environment.Environment).SetVariable(p.Id, iSymbol)
					for _, inst := range p.Bloque {
						inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
					}
				}

			}

		} else {
			//Es ID entonces
			var id1 environment.Symbol
			id1 = env.(environment.Environment).GetVariable(str1)
			if id1.Tipo == environment.STRING {
				var forEnv environment.Environment
				forEnv = environment.NewEnvironment(env.(environment.Environment), "For")
				iSymbol0 := environment.Symbol{Valor: id1.Valor, Tipo: environment.STRING}
				env.(environment.Environment).SaveVariable(p.Id, iSymbol0)
				for _, c := range id1.Valor.(string) {
					iSymbol := environment.Symbol{Valor: string(c), Tipo: environment.STRING} // Crear un nuevo Symbol con el valor 'i'
					env.(environment.Environment).SetVariable(p.Id, iSymbol)
					for _, inst := range p.Bloque {
						inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
					}
				}

			} else if id1.Tipo == environment.ARRAY {
				var forEnv environment.Environment
				forEnv = environment.NewEnvironment(env.(environment.Environment), "For")
				iSymbol0 := environment.Symbol{Valor: 0, Tipo: environment.INTEGER}
				env.(environment.Environment).SaveVariable(p.Id, iSymbol0)
				aux := id1.Valor.([]interface{})
				count := len(aux)

				for i := 0; i < count; i++ {
					iSymbol := environment.Symbol{Valor: i, Tipo: environment.INTEGER} // Crear un nuevo Symbol con el valor 'i'
					env.(environment.Environment).SetVariable(p.Id, iSymbol)
					for _, inst := range p.Bloque {
						inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
					}

				}

			}
		}

	}

	return nil

}
