package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type If struct {
	Lin         int
	Col         int
	Expresion   interfaces.Expression
	Bloque      []interface{}
	ElseIfBlock []interface{}
	ElseBlock   []interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque []interface{}, elseifblock []interface{}, elseblock []interface{}) If {
	ifInstr := If{lin, col, condition, bloque, elseifblock, elseblock}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var condicion environment.Symbol
	condicion = p.Expresion.Ejecutar(ast, env)
	//Validando tipo
	if condicion.Tipo != environment.BOOLEAN {
		ast.SetError("El tipo de variable es incorrecto para un If")
		return nil
	}
	var ifEnv environment.Environment
	ifEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
	//Ejecutando if
	if condicion.Valor == true {

		//ejecución
		for _, inst := range p.Bloque {
			inst.(interfaces.Instruction).Ejecutar(ast, ifEnv)
		}
		//print("if")
		//ast.Tabla_str += ifEnv.FormatSymbolTable()

		return nil
	} else {
		//agregar condiciones para else
	}

	if p.ElseIfBlock != nil {
		println("EN GO ELSEIF")
		var ifEnv environment.Environment
		ifEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
		//ejecución
		for _, inst := range p.ElseIfBlock {
			inst.(interfaces.Instruction).Ejecutar(ast, ifEnv)
		}
		return nil
	}

	if p.ElseBlock != nil {
		var ifEnv environment.Environment
		ifEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
		//ejecución
		for _, inst := range p.ElseBlock {
			inst.(interfaces.Instruction).Ejecutar(ast, ifEnv)
		}
		return nil
	}

	return nil
}
