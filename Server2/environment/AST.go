package environment

import (
	"fmt"
	"strings"
)

type Nodo struct {
	Hijos []*Nodo
	Padre *Nodo
	Valor string
}

func NewNodo(valor string) *Nodo {
	return &Nodo{Valor: valor, Padre: nil}
}

func (n *Nodo) GenerateProduction(labels []string) {

	if len(labels) > 0 {
		combinedValue := ""
		for _, label := range labels {
			combinedValue += label
		}
		n.Hijos = append(n.Hijos, NewNodo(combinedValue))
	}
}

func (n *Nodo) GetValor() string {
	return n.Valor
}

func (n *Nodo) SetValor(valor string) {
	n.Valor = valor
}

func (n *Nodo) SetHijos(hijos []*Nodo) {
	n.Hijos = hijos
}

func (n *Nodo) SetPadre(padre *Nodo) {
	n.Padre = padre
}

func (n *Nodo) GetPadre() *Nodo {
	return n.Padre
}

func (n *Nodo) GetHijos() []*Nodo {
	return n.Hijos
}

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       string
	Raiz         *Nodo
	Switchexpr   string
	SwitchBool   bool
}

func NewAST(inst []interface{}, print string, err string, valorRaiz string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err, Raiz: NewNodo(valorRaiz)}
	return ast
}

func (a *AST) GetRaiz() *Nodo {
	return a.Raiz
}

func (a *AST) SetRaiz(nodoRaiz *Nodo) {
	a.Raiz = nodoRaiz
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

func (a *AST) GenerateDot() string {
	dotBuilder := strings.Builder{}
	dotBuilder.WriteString("digraph AST {\n")

	nodeIDMap := make(map[*Nodo]int) // Para asignar IDs únicos a los nodos
	nodeCount := 0

	// Función auxiliar para asignar IDs a los nodos
	var assignNodeID func(*Nodo) int
	assignNodeID = func(node *Nodo) int {
		if _, exists := nodeIDMap[node]; !exists {
			nodeCount++
			nodeIDMap[node] = nodeCount
		}
		return nodeIDMap[node]
	}

	// Recorrido en profundidad (DFS) para generar conexiones DOT
	var buildDot func(*Nodo)
	buildDot = func(node *Nodo) {
		nodeID := assignNodeID(node)
		escapedLabel := strings.ReplaceAll(node.Valor, "\"", "\\\"")
		dotBuilder.WriteString(fmt.Sprintf("  %d [label=\"%s\"];\n", nodeID, escapedLabel))

		for _, child := range node.Hijos {
			childID := assignNodeID(child)
			dotBuilder.WriteString(fmt.Sprintf("  %d -> %d;\n", nodeID, childID))
			buildDot(child)
		}
	}

	buildDot(a.Raiz)
	dotBuilder.WriteString("}")

	return dotBuilder.String()
}

/*
func (a *AST) BuildTree(padre, nodoPadre *Nodo) {
	nodos := padre.GetHijos()
	for _, nodo := range nodos {
		node := NewNodo(nodo.GetValor())
		node.SetPadre(nodoPadre)

		nodoPadre.Hijos = append(nodoPadre.Hijos, node)
		a.BuildTree(nodo, node)
	}
}

func (a *AST) GenerateDotString() string {
	dotString := "digraph G {\n"
	contador := 0
	nodos := []*Nodo{a.Raiz}
	for len(nodos) > 0 {
		nodoActual := nodos[len(nodos)-1]
		nodos = nodos[:len(nodos)-1]

		dotString += "  " + strconv.FormatUint(uint64(uintptr(unsafe.Pointer(nodoActual))), 10) + " [label=\"" + nodoActual.GetValor() + "\"];\n"

		hijos := nodoActual.GetHijos()
		contador += 1
		fmt.Println(nodoActual.Valor, contador)
		for _, hijo := range hijos {
			nodos = append(nodos, hijo)
			dotString += "  " + strconv.FormatUint(uint64(uintptr(unsafe.Pointer(nodoActual))), 10) + " -> " + strconv.FormatUint(uint64(uintptr(unsafe.Pointer(hijo))), 10) + ";\n"
		}
	}

	dotString += "}"

	return dotString
}*/
