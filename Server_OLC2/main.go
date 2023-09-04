package main

import (
	"Server2/environment"
	"Server2/interfaces"
	"Server2/parser"
	"fmt"
	"os/exec"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"os"
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

func generarDot(dotData string, rutaArchivo string, nombreArchivo string) error {
	// Crear el archivo DOT en la ubicación especificada
	archivo, err := os.Create(rutaArchivo + "/" + nombreArchivo + ".dot")
	if err != nil {
		return err
	}
	defer archivo.Close()

	// Escribir los datos DOT en el archivo
	_, err = archivo.WriteString(dotData)
	if err != nil {
		return err
	}

	return nil
}

func generarImagen(rutaDot string, rutaImagen string, nombreArchivo string) error {
	// Comando para generar la imagen SVG desde el archivo DOT
	cmd := exec.Command("dot", "-Tsvg", rutaDot, "-o", rutaImagen+"/"+nombreArchivo+".svg")

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
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

func handleInterpreter() error {
	var message Message

	/*if err := c.BodyParser(&message); err != nil {
		return err
	}*/
	message.Content = " let vec1: Int = 1 \n vec1=5 \n  var vec2: Int = 5 \n vec2=100 print(vec2)  "

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
	println("------------------------------------------------")

	//Para generar arbol
	dotCST := TreesStringTreeToDOT(tree, p.GetRuleNames(), nil)
	rutaImagen := "C:\\Users\\sebas\\go\\bin\\client\\src\\pages\\imagenes"
	nombreArchivo := "arbol"
	if err := generarDot(dotCST, rutaImagen, nombreArchivo); err != nil {
		fmt.Println("Error al generar el archivo DOT:", err)
		return nil
	}

	// Llama a la función generarImagen para crear la imagen PNG a partir del archivo DOT
	if err := generarImagen(rutaImagen+"/"+nombreArchivo+".dot", rutaImagen, nombreArchivo); err != nil {
		fmt.Println("Error al generar la imagen:", err)
		return nil
	}

	//Para generar Tabla de simbolos
	dotTable := GenerateTableDOT(Ast.Tabla)
	rutaImagenTabla := "C:\\Users\\sebas\\go\\bin\\client\\src\\pages\\imagenes"
	nombreArchivoTabla := "tabla"

	if err := generarDot(dotTable, rutaImagenTabla, nombreArchivoTabla); err != nil {
		fmt.Println("Error al generar el archivo DOT:", err)
		return nil
	}

	// Llama a la función generarImagen para crear la imagen PNG a partir del archivo DOT
	if err := generarImagen(rutaImagenTabla+"/"+nombreArchivoTabla+".dot", rutaImagenTabla, nombreArchivoTabla); err != nil {
		fmt.Println("Error al generar la imagen:", err)
		return nil
	}
	//return c.Status(fiber.StatusOK).JSON(response)
	//return ConsoleOut
	return nil
}

func main() {
	/*	app := fiber.New()
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
