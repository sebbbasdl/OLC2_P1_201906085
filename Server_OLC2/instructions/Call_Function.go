package instructions

/*
import (
	"Server2/environment"
	"Server2/interfaces"
)

type call_func struct {
	Line    int
	Col     int
	Id      string
	ListExp []interface{}
}

func NewCallFunc(line, col int, id string, list []interface{}) call_func {
	return call_func{line, col, id, list}
}

func (c call_func) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var funcEnv environment.Environment
	funcEnv = environment.NewEnvironment(env.(environment.Environment), "Func")
	symFunc := environment.GetFunc(c.Id, env.(environment.Environment))
	var tempParams []interface{}
	tempParams = symFunc.ParamList
	var instr interfaces.Instruction

	if c.ListExp != nil {
		mapTemp := tempParams
		if len(c.ListExp) == len(mapTemp) {
			cont := 0
			for _, value := range mapTemp {
				sym := c.ListExp[cont].(environment.Symbol)
				if sym.Tipo == value.(environment.Symbol).Tipo {
					funcEnv.SaveVariable(c.Id, sym)
				} else {
					ast.SetError("Error: tipos de parámetros incompatibles")
				}
				cont++
			}

		} else {
			ast.SetError("Error: la cantidad de parámetros no es la correcta")
			return nil
		}
	}

	instr.Ejecutar(ast, funcEnv)

	return nil
}
*/
