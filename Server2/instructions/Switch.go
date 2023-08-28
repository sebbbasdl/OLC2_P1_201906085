package instructions

import (
	"Server2/environment"
	"Server2/interfaces"
)

type Switch struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	CaseBlock []interface{}
	Block     []interface{}
}

func NewSwitch(lin int, col int, expr1 interfaces.Expression, casebloque []interface{}, block []interface{}) Switch {
	switchInstr := Switch{lin, col, expr1, casebloque, block}
	return switchInstr
}

func (p Switch) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var expr1 environment.Symbol
	expr1 = p.Expresion.Ejecutar(ast, env)
	var switchEnv environment.Environment
	switchEnv = environment.NewEnvironment(env.(environment.Environment), "Switch")

	for _, inst := range p.CaseBlock {
		if caseInstr, ok := inst.(Case); ok {
			vari := caseInstr.Expresion
			//fmt.Println(vari)
			var expr2 interfaces.Expression
			expr2 = vari
			if expr1.Valor == expr2.Ejecutar(ast, switchEnv).Valor {
				ast.SwitchBool = true
				inst.(interfaces.Instruction).Ejecutar(ast, switchEnv)
				ast.SwitchBool = false
				return nil
			}
		}
	}

	for _, inst := range p.Block {
		inst.(interfaces.Instruction).Ejecutar(ast, switchEnv)
	}

	return nil
}
