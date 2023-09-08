package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type CallFunc struct {
	Lin    int
	Col    int
	Id     string
	Params []interface{}
}

func NewCall_Function(lin int, col int, id string, par []interface{}) CallFunc {
	exp := CallFunc{lin, col, id, par}
	return exp
}

func (p CallFunc) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	var funcSym environment.FunctionSymbol
	//obtener la funcion anteriormente guardada
	funcSym = env.(environment.Environment).GetFunction(p.Id)
	//definiendo nuevo entorno de funcion
	var funcEnv environment.Environment
	//si no es modulo se usa el global como anterior

	funcEnv = environment.NewEnvironment(p.GetGlobal(env), p.Id+"(FUNCTION)")
	//se agregan los parámetros como variables
	if len(p.Params) == len(funcSym.ListDec) {
		for i := 0; i < len(p.Params); i++ {
			var val environment.Symbol
			val = p.Params[i].(interfaces.Expression).Ejecutar(ast, env)
			if val.Tipo == funcSym.ListDec[i].(ParamsDeclaration).Tipo {
				funcEnv.SaveVariable(funcSym.ListDec[i].(ParamsDeclaration).Id, val)
			} else {
				ast.SetError("El tipo de parámetro es incorrecto")
				return result
			}
		}
	} else {
		ast.SetError("Faltan parámetros en la función")
		return result
	}
	//ejecutar bloque con entorno funcEnv
	for _, inst := range funcSym.Block {
		inst.(interfaces.Instruction).Ejecutar(ast, funcEnv)
	}
	return nil
}

func (p CallFunc) GetGlobal(env interface{}) environment.Environment {
	//obteniendo entorno global
	var tmpEnvGbl environment.Environment
	tmpEnvGbl = env.(environment.Environment)
	for {
		if tmpEnvGbl.Id == "GLOBAL" {
			return tmpEnvGbl
		}
		if tmpEnvGbl.Anterior == nil {
			break
		} else {
			tmpEnvGbl = tmpEnvGbl.Anterior.(environment.Environment)
		}
	}
	return tmpEnvGbl
}
