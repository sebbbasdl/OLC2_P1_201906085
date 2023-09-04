package environment

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       string
	Switchexpr   string
	SwitchBool   bool
	Breakbool    bool
	ContinueBool bool
	Ant_str      string
	Tabla        [][]string
	Constantes   []string
}

func NewAST(inst []interface{}, print string, err string, valorRaiz string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() string {
	return a.Errors
}

func (a *AST) SetError(ToErr string) {
	a.Errors = a.Errors + ToErr + "\n"
}

func modificarValorEnTabla(tabla [][]string, dato []string) {

	for i := 0; i < len(tabla); i++ {
		if tabla[i][4] == dato[4] && tabla[i][5] == dato[5] {
			//modifica valor
			tabla[i] = dato
		}
	}

}
