package environment

type funcSymbol struct {
	Line      int
	Col       int
	Tipo      TipoExpresion
	Id        string
	ParamList []interface{}
	InstList  interface{}
}

/*func NewFuncSymbol() funcSymbol {
	return funcSymbol{}
}

func NewFuncSymbolWithParams(line, col int, tipo TipoExpresion, id string, param []interface{}, inst interface{}) funcSymbol {
	return funcSymbol{
		Line:      line,
		Col:       col,
		Tipo:      tipo,
		Id:        id,
		ParamList: param,
		InstList:  inst,
	}
}
*/
