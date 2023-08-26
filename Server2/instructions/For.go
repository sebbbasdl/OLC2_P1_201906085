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
	Bloque     []interface{}
}

func NewFor(lin int, col int, id string, expr1 interfaces.Expression, expr2 interfaces.Expression, bloque []interface{}) Fors {
	forInstr := Fors{lin, col, id, expr1, expr2, bloque}
	return forInstr
}

func (p Fors) Ejecutar(ast *environment.AST, env interface{}) interface{} {
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

	for i := num1Value; i <= num2Value; i++ {
		var forEnv environment.Environment
		forEnv = environment.NewEnvironment(env.(environment.Environment), "For")
		iSymbol := environment.Symbol{Valor: i, Tipo: environment.INTEGER} // Crear un nuevo Symbol con el valor 'i'
		env.(environment.Environment).SetVariable(p.Id, iSymbol)
		//ejecución
		for _, inst := range p.Bloque {
			inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
		}

	}

	return nil

}
