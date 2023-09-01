package main

import (
	"Server2/environment"
	"Server2/interfaces"
	"Server2/parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Tabla de simbolos DOT
func GenerateTableDOT(data [][]string) string {
	contador := 0
	dotTable := "digraph Table {\n"
	dotTable += "  node [shape=none fontname=Helvetica]\n"

	// Encabezado de la tabla
	dotTable += "  n1 [label = <<table>\n"
	dotTable += "    <tr><td colspan=\"9\" bgcolor=\"Peru\">Tabla</td></tr>\n"
	dotTable += "    <tr><td bgcolor=\"orange\">ID</td><td bgcolor=\"orange\">Tipo Símbolo</td><td bgcolor=\"orange\">Tipo Dato</td><td bgcolor=\"orange\">Ámbito</td><td bgcolor=\"orange\">Línea</td><td bgcolor=\"orange\">Columna</td><td bgcolor=\"orange\">Valor</td><td bgcolor=\"orange\">Id</td></tr>"

	// Filas de datos
	for _, row := range data {
		if len(row) >= 7 {
			contador++
			contadorStr := fmt.Sprintf("%d", contador) // Convertir contador a cadena
			dotTable += fmt.Sprintf("    <tr><td bgcolor=\"#00bfff\">%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>\n", contadorStr, row[1], row[2], row[3], row[4], row[5], row[6], row[7])
		}
	}

	// Cierre de la tabla
	dotTable += "  </table>> ]\n"

	dotTable += "}\n"

	return dotTable
}

// Para generar Dot usando funciones de antlr
func TreesStringTreeToDOT(tree antlr.Tree, ruleNames []string, recog antlr.Recognizer) string {
	if recog != nil {
		ruleNames = recog.GetRuleNames()
	}

	nodeCounter := 0                    // Para asignar nombres únicos a los nodos
	nodeIDs := make(map[antlr.Tree]int) // Mapa para almacenar IDs de nodos

	var generateDOT func(antlr.Tree) (int, string)
	generateDOT = func(node antlr.Tree) (int, string) {
		nodeID, exists := nodeIDs[node]
		if !exists {
			nodeCounter++
			nodeID = nodeCounter
			nodeIDs[node] = nodeID
		}

		nodeText := antlr.TreesGetNodeText(node, ruleNames, nil)
		nodeText = strings.ReplaceAll(nodeText, "\"", "\\\"") // Escapar comillas

		dot := fmt.Sprintf("  %d [label=\"%s\"];\n", nodeID, nodeText)

		childCount := node.GetChildCount()
		childDot := ""
		for i := 0; i < childCount; i++ {
			child := node.GetChild(i)
			childID, childDotPart := generateDOT(child)
			childDot += fmt.Sprintf("  %d -> %d;\n", nodeID, childID) + childDotPart
		}

		return nodeID, dot + childDot
	}

	dotTree := "digraph AST {\n"
	_, dotTreePart := generateDOT(tree)
	dotTree += dotTreePart
	dotTree += "}\n"

	return dotTree
}

type ErrorType int

const (
	LexicalError ErrorType = iota
	SyntaxTypeError
)

type ErrorInfo struct {
	Type    ErrorType
	Line    int
	Column  int
	Message string
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []ErrorInfo
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

func (e *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	errorType := LexicalError

	if _, ok := offendingSymbol.(antlr.Token); ok {
		errorType = SyntaxTypeError
	}

	errorInfo := ErrorInfo{
		Type:    errorType,
		Line:    line,
		Column:  column,
		Message: msg,
	}
	e.Errors = append(e.Errors, errorInfo)
}

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

type Resp struct {
	Output  string
	Flag    bool
	Message string
}

type Message struct {
	Content string `json:"content"`
}

func handleInterpreter() interface{} {
	var message Message
	/*if err := c.BodyParser(&message); err != nil {
		return err
	}*/
	message.Content = " var x1 : Int = 9 \n while x1 <= 10 {   var x2 : Int = 9 if x2 == 9 {var x3 : Int = 9 \n x3 =999 print(x3+\"e\") }  print(x1) x1=x1+1 }"

	//Entrada
	code := message.Content
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	errorListener := NewErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.RemoveErrorListeners() // Remove the default error listener
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	tree := p.S()

	//dotRepresentation := TreesStringTreeToDOT(tree, p.GetRuleNames(), nil)
	//fmt.Println(dotRepresentation)
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST

	//create env
	var globalEnv environment.Environment = environment.NewEnvironment(nil, "GLOBAL")

	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Ejecutar(&Ast, globalEnv)
	}

	for _, errorInfo := range errorListener.Errors {
		errorType := ""
		switch errorInfo.Type {
		case LexicalError:
			errorType = "Lexical"
		case SyntaxTypeError:
			errorType = "Syntax Error"
		default:
			errorType = "Unknown"
		}

		fmt.Printf("%s error at line %d, column %d: %s\n", errorType, errorInfo.Line, errorInfo.Column, errorInfo.Message)
	}
	//Ast.BuildTree(rootNode, Ast.GetRaiz())
	// Generate DOT string from the AST

	// Print or use the generated DOT string as needed
	//println(Ast.GenerateDot())

	var ConsoleOut = ""
	if Ast.GetErrors() == "" {
		ConsoleOut = Ast.GetPrint()
	} else {
		ConsoleOut = Ast.GetErrors()
	}

	/*response := Resp{
		Output:  ConsoleOut,
		Flag:    true,
		Message: "<3 Ejecución realizada con éxito <3",
	}*/

	println("--------------------CONSOLA---------------------")
	println(ConsoleOut)
	println("-------------Tabla de Simbolos------------------")
	//Ast.Tabla_str += globalEnv.FormatSymbolTable()
	//println(Ast.Tabla_str)
	/*data := [][]string{
		{"1", "Variable", "Int", "Global", "10", "5"},
		{"2", "Función", "Void", "Local", "15", "3"},
		// ... más filas de datos
	}*/

	dotTable := GenerateTableDOT(Ast.Tabla)
	fmt.Println(dotTable)

	return ConsoleOut
}

func main() {
	/*app := fiber.New()
	app.Use(cors.New())
	app.Post("/Interpreter", handleInterpreter)
	app.Listen(":3002")*/
	handleInterpreter()

}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}
