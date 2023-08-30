package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Guard struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, bloque []interface{}) Guard {
	guardInstr := Guard{lin, col, condition, bloque}
	return guardInstr
}

func (p Guard) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var condicion environment.Symbol
	condicion = p.Expresion.Ejecutar(ast, env)

	if condicion.Tipo != environment.BOOLEAN {
		ast.SetError("El tipo de variable es incorrecto para un If")
		return nil
	}

	if condicion.Valor != true {
		var guardEnv environment.Environment
		guardEnv = environment.NewEnvironment(env.(environment.Environment), "GUARD")
		//ejecución
		for _, inst := range p.Bloque {
			if ast.Breakbool == true {
				break
			} else if ast.ContinueBool == true {
				//ast.ContinueBool = false
				//inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
				continue
			} else {
				inst.(interfaces.Instruction).Ejecutar(ast, guardEnv)
			}

		}

	}
	ast.Breakbool = false
	//ast.ContinueBool = false

	return nil
}
