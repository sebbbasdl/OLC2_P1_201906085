package main

import (
	"Server2/environment"
	"Server2/interfaces"
	"Server2/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

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
	message.Content = " var i : Int  = 2 \n while (i <= 5) {if i == 3  { print(\"e\") \n i = i + 1 \n var prueba : Int  = 10\n continue}print(i)i = i + 1}"

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

	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST

	rootNode := environment.NewNodo("Root")
	Ast.SetRaiz(rootNode)
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
	//println("-------------Tabla de Simbolos------------------")
	//globalEnv.PrintChain()

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
