package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Whiles struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewWhiles(lin int, col int, condition interfaces.Expression, bloque []interface{}) Whiles {
	whileInstr := Whiles{lin, col, condition, bloque}
	return whileInstr
}

func (p Whiles) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var condicion environment.Symbol
	condicion = p.Expresion.Ejecutar(ast, env)
	//Validando tipo
	if condicion.Tipo != environment.BOOLEAN {
		ast.SetError("El tipo de variable es incorrecto para un If")
		return nil
	}
	var whileEnv environment.Environment
	for condicion.Valor == true {
		//fmt.Println(condicion.Valor)

		whileEnv = environment.NewEnvironment(env.(environment.Environment), "While")
		//ejecución
		for _, inst := range p.Bloque {
			if ast.Breakbool == true {

				break
			} else if ast.ContinueBool == true {
				//ast.ContinueBool = false
				//inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
				continue
			} else {
				inst.(interfaces.Instruction).Ejecutar(ast, whileEnv)

			}

		}
		condicion = p.Expresion.Ejecutar(ast, env)
		if ast.Breakbool == true {
			break
		} else if ast.ContinueBool == true {
			ast.ContinueBool = false
			//inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
			continue
		}

	}
	print("while")
	ast.Tabla_str = whileEnv.FormatSymbolTable()
	ast.Breakbool = false
	ast.ContinueBool = false

	return nil
}
