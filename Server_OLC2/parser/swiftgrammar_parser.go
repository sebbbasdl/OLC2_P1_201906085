// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Server2/interfaces"
import "Server2/environment"
import "Server2/expressions"
import "Server2/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", "'print'",
		"'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", "'default'",
		"'in'", "'var'", "'let'", "'break'", "'continue'", "'return'", "'guard'",
		"'append'", "'removeLast'", "'remove'", "'count'", "'isEmpty'", "'at'",
		"'func'", "'struct'", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'",
		"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('",
		"')'", "'{'", "'}'", "':'", "'['", "']'", "','", "'->'", "'.'", "'%'",
		"'_'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "LET", "BREAK",
		"CONTINUE", "RETURN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", "COUNT",
		"ISEMPTY", "AT", "FUNC", "STRUCT", "NUMBER", "STRING", "ID", "DIF",
		"IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR",
		"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"D_PTS", "CORIZQ", "CORDER", "COMA", "ARROW1", "PUNTO", "MODULO", "GUIONB",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "returntmt", "function", "listParamsFunc",
		"structCreation", "listStructDec", "removetmt", "removelastmt", "appendtmt",
		"guardtmt", "breaktmt", "continuetmt", "printstmt", "forstmt", "switchtmt",
		"caselist", "case", "ifstmt", "elseiflist", "elseif", "whilestmt", "declarationstmt",
		"asignacion", "types", "expr", "listParams", "listArray", "callFunctionInst",
		"callFunction", "listParamsCall", "listStructExp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 664, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 72, 8, 1, 11, 1,
		12, 1, 73, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 129, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 157, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 166, 8,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 175, 8, 5, 10, 5, 12,
		5, 178, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 195, 8, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 205, 8, 7, 10, 7, 12, 7, 208, 9, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 263,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 300, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 320, 8, 17, 10, 17, 12, 17, 323,
		9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 3, 19, 369, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 5, 20, 379, 8, 20, 10, 20, 12, 20, 382, 9, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 439, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 452, 8, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 472, 8, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 3, 26, 539, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 576, 8, 26, 10,
		26, 12, 26, 579, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 5, 27, 590, 8, 27, 10, 27, 12, 27, 593, 9, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 5, 28, 609, 8, 28, 10, 28, 12, 28, 612, 9, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 631, 8, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 5, 31, 638, 8, 31, 10, 31, 12, 31, 641, 9, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 650, 8, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 659, 8, 32, 10, 32, 12, 32,
		662, 9, 32, 1, 32, 0, 9, 10, 14, 34, 40, 52, 54, 56, 62, 64, 33, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 6, 1, 0, 31, 32, 2,
		0, 43, 44, 57, 57, 1, 0, 45, 46, 2, 0, 39, 39, 41, 41, 2, 0, 40, 40, 42,
		42, 1, 0, 33, 34, 701, 0, 66, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 128, 1,
		0, 0, 0, 6, 130, 1, 0, 0, 0, 8, 156, 1, 0, 0, 0, 10, 165, 1, 0, 0, 0, 12,
		179, 1, 0, 0, 0, 14, 194, 1, 0, 0, 0, 16, 209, 1, 0, 0, 0, 18, 219, 1,
		0, 0, 0, 20, 226, 1, 0, 0, 0, 22, 234, 1, 0, 0, 0, 24, 242, 1, 0, 0, 0,
		26, 245, 1, 0, 0, 0, 28, 262, 1, 0, 0, 0, 30, 299, 1, 0, 0, 0, 32, 301,
		1, 0, 0, 0, 34, 311, 1, 0, 0, 0, 36, 324, 1, 0, 0, 0, 38, 368, 1, 0, 0,
		0, 40, 370, 1, 0, 0, 0, 42, 383, 1, 0, 0, 0, 44, 391, 1, 0, 0, 0, 46, 438,
		1, 0, 0, 0, 48, 451, 1, 0, 0, 0, 50, 471, 1, 0, 0, 0, 52, 538, 1, 0, 0,
		0, 54, 580, 1, 0, 0, 0, 56, 594, 1, 0, 0, 0, 58, 613, 1, 0, 0, 0, 60, 619,
		1, 0, 0, 0, 62, 630, 1, 0, 0, 0, 64, 649, 1, 0, 0, 0, 66, 67, 3, 2, 1,
		0, 67, 68, 5, 0, 0, 1, 68, 69, 6, 0, -1, 0, 69, 1, 1, 0, 0, 0, 70, 72,
		3, 4, 2, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 6, 1, -1, 0, 76, 3, 1,
		0, 0, 0, 77, 78, 3, 28, 14, 0, 78, 79, 6, 2, -1, 0, 79, 129, 1, 0, 0, 0,
		80, 81, 3, 38, 19, 0, 81, 82, 6, 2, -1, 0, 82, 129, 1, 0, 0, 0, 83, 84,
		3, 46, 23, 0, 84, 85, 6, 2, -1, 0, 85, 129, 1, 0, 0, 0, 86, 87, 3, 48,
		24, 0, 87, 88, 6, 2, -1, 0, 88, 129, 1, 0, 0, 0, 89, 90, 3, 44, 22, 0,
		90, 91, 6, 2, -1, 0, 91, 129, 1, 0, 0, 0, 92, 93, 3, 30, 15, 0, 93, 94,
		6, 2, -1, 0, 94, 129, 1, 0, 0, 0, 95, 96, 3, 32, 16, 0, 96, 97, 6, 2, -1,
		0, 97, 129, 1, 0, 0, 0, 98, 99, 3, 24, 12, 0, 99, 100, 6, 2, -1, 0, 100,
		129, 1, 0, 0, 0, 101, 102, 3, 26, 13, 0, 102, 103, 6, 2, -1, 0, 103, 129,
		1, 0, 0, 0, 104, 105, 3, 22, 11, 0, 105, 106, 6, 2, -1, 0, 106, 129, 1,
		0, 0, 0, 107, 108, 3, 20, 10, 0, 108, 109, 6, 2, -1, 0, 109, 129, 1, 0,
		0, 0, 110, 111, 3, 18, 9, 0, 111, 112, 6, 2, -1, 0, 112, 129, 1, 0, 0,
		0, 113, 114, 3, 16, 8, 0, 114, 115, 6, 2, -1, 0, 115, 129, 1, 0, 0, 0,
		116, 117, 3, 8, 4, 0, 117, 118, 6, 2, -1, 0, 118, 129, 1, 0, 0, 0, 119,
		120, 3, 58, 29, 0, 120, 121, 6, 2, -1, 0, 121, 129, 1, 0, 0, 0, 122, 123,
		3, 12, 6, 0, 123, 124, 6, 2, -1, 0, 124, 129, 1, 0, 0, 0, 125, 126, 3,
		6, 3, 0, 126, 127, 6, 2, -1, 0, 127, 129, 1, 0, 0, 0, 128, 77, 1, 0, 0,
		0, 128, 80, 1, 0, 0, 0, 128, 83, 1, 0, 0, 0, 128, 86, 1, 0, 0, 0, 128,
		89, 1, 0, 0, 0, 128, 92, 1, 0, 0, 0, 128, 95, 1, 0, 0, 0, 128, 98, 1, 0,
		0, 0, 128, 101, 1, 0, 0, 0, 128, 104, 1, 0, 0, 0, 128, 107, 1, 0, 0, 0,
		128, 110, 1, 0, 0, 0, 128, 113, 1, 0, 0, 0, 128, 116, 1, 0, 0, 0, 128,
		119, 1, 0, 0, 0, 128, 122, 1, 0, 0, 0, 128, 125, 1, 0, 0, 0, 129, 5, 1,
		0, 0, 0, 130, 131, 5, 20, 0, 0, 131, 132, 3, 52, 26, 0, 132, 133, 6, 3,
		-1, 0, 133, 7, 1, 0, 0, 0, 134, 135, 5, 28, 0, 0, 135, 136, 5, 32, 0, 0,
		136, 137, 5, 47, 0, 0, 137, 138, 3, 10, 5, 0, 138, 139, 5, 48, 0, 0, 139,
		140, 5, 49, 0, 0, 140, 141, 3, 2, 1, 0, 141, 142, 5, 50, 0, 0, 142, 143,
		6, 4, -1, 0, 143, 157, 1, 0, 0, 0, 144, 145, 5, 28, 0, 0, 145, 146, 5,
		32, 0, 0, 146, 147, 5, 47, 0, 0, 147, 148, 3, 10, 5, 0, 148, 149, 5, 48,
		0, 0, 149, 150, 5, 55, 0, 0, 150, 151, 3, 50, 25, 0, 151, 152, 5, 49, 0,
		0, 152, 153, 3, 2, 1, 0, 153, 154, 5, 50, 0, 0, 154, 155, 6, 4, -1, 0,
		155, 157, 1, 0, 0, 0, 156, 134, 1, 0, 0, 0, 156, 144, 1, 0, 0, 0, 157,
		9, 1, 0, 0, 0, 158, 159, 6, 5, -1, 0, 159, 160, 5, 32, 0, 0, 160, 161,
		5, 51, 0, 0, 161, 162, 3, 50, 25, 0, 162, 163, 6, 5, -1, 0, 163, 166, 1,
		0, 0, 0, 164, 166, 6, 5, -1, 0, 165, 158, 1, 0, 0, 0, 165, 164, 1, 0, 0,
		0, 166, 176, 1, 0, 0, 0, 167, 168, 10, 3, 0, 0, 168, 169, 5, 54, 0, 0,
		169, 170, 5, 32, 0, 0, 170, 171, 5, 51, 0, 0, 171, 172, 3, 50, 25, 0, 172,
		173, 6, 5, -1, 0, 173, 175, 1, 0, 0, 0, 174, 167, 1, 0, 0, 0, 175, 178,
		1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 11, 1, 0,
		0, 0, 178, 176, 1, 0, 0, 0, 179, 180, 5, 29, 0, 0, 180, 181, 5, 32, 0,
		0, 181, 182, 5, 49, 0, 0, 182, 183, 3, 14, 7, 0, 183, 184, 5, 50, 0, 0,
		184, 185, 6, 6, -1, 0, 185, 13, 1, 0, 0, 0, 186, 187, 6, 7, -1, 0, 187,
		188, 5, 16, 0, 0, 188, 189, 5, 32, 0, 0, 189, 190, 5, 51, 0, 0, 190, 191,
		3, 50, 25, 0, 191, 192, 6, 7, -1, 0, 192, 195, 1, 0, 0, 0, 193, 195, 6,
		7, -1, 0, 194, 186, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 206, 1, 0, 0,
		0, 196, 197, 10, 3, 0, 0, 197, 198, 5, 54, 0, 0, 198, 199, 5, 16, 0, 0,
		199, 200, 5, 32, 0, 0, 200, 201, 5, 51, 0, 0, 201, 202, 3, 50, 25, 0, 202,
		203, 6, 7, -1, 0, 203, 205, 1, 0, 0, 0, 204, 196, 1, 0, 0, 0, 205, 208,
		1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 15, 1, 0,
		0, 0, 208, 206, 1, 0, 0, 0, 209, 210, 5, 32, 0, 0, 210, 211, 5, 56, 0,
		0, 211, 212, 5, 24, 0, 0, 212, 213, 5, 47, 0, 0, 213, 214, 5, 27, 0, 0,
		214, 215, 5, 51, 0, 0, 215, 216, 3, 52, 26, 0, 216, 217, 5, 48, 0, 0, 217,
		218, 6, 8, -1, 0, 218, 17, 1, 0, 0, 0, 219, 220, 5, 32, 0, 0, 220, 221,
		5, 56, 0, 0, 221, 222, 5, 23, 0, 0, 222, 223, 5, 47, 0, 0, 223, 224, 5,
		48, 0, 0, 224, 225, 6, 9, -1, 0, 225, 19, 1, 0, 0, 0, 226, 227, 5, 32,
		0, 0, 227, 228, 5, 56, 0, 0, 228, 229, 5, 22, 0, 0, 229, 230, 5, 47, 0,
		0, 230, 231, 3, 52, 26, 0, 231, 232, 5, 48, 0, 0, 232, 233, 6, 10, -1,
		0, 233, 21, 1, 0, 0, 0, 234, 235, 5, 21, 0, 0, 235, 236, 3, 52, 26, 0,
		236, 237, 5, 9, 0, 0, 237, 238, 5, 49, 0, 0, 238, 239, 3, 2, 1, 0, 239,
		240, 5, 50, 0, 0, 240, 241, 6, 11, -1, 0, 241, 23, 1, 0, 0, 0, 242, 243,
		5, 18, 0, 0, 243, 244, 6, 12, -1, 0, 244, 25, 1, 0, 0, 0, 245, 246, 5,
		19, 0, 0, 246, 247, 6, 13, -1, 0, 247, 27, 1, 0, 0, 0, 248, 249, 5, 7,
		0, 0, 249, 250, 5, 47, 0, 0, 250, 251, 3, 52, 26, 0, 251, 252, 5, 48, 0,
		0, 252, 253, 6, 14, -1, 0, 253, 263, 1, 0, 0, 0, 254, 255, 5, 7, 0, 0,
		255, 256, 5, 47, 0, 0, 256, 257, 3, 52, 26, 0, 257, 258, 5, 54, 0, 0, 258,
		259, 3, 52, 26, 0, 259, 260, 5, 48, 0, 0, 260, 261, 6, 14, -1, 0, 261,
		263, 1, 0, 0, 0, 262, 248, 1, 0, 0, 0, 262, 254, 1, 0, 0, 0, 263, 29, 1,
		0, 0, 0, 264, 265, 5, 11, 0, 0, 265, 266, 5, 32, 0, 0, 266, 267, 5, 15,
		0, 0, 267, 268, 3, 52, 26, 0, 268, 269, 5, 56, 0, 0, 269, 270, 5, 56, 0,
		0, 270, 271, 5, 56, 0, 0, 271, 272, 3, 52, 26, 0, 272, 273, 5, 49, 0, 0,
		273, 274, 3, 2, 1, 0, 274, 275, 5, 50, 0, 0, 275, 276, 6, 15, -1, 0, 276,
		300, 1, 0, 0, 0, 277, 278, 5, 11, 0, 0, 278, 279, 5, 58, 0, 0, 279, 280,
		5, 15, 0, 0, 280, 281, 3, 52, 26, 0, 281, 282, 5, 56, 0, 0, 282, 283, 5,
		56, 0, 0, 283, 284, 5, 56, 0, 0, 284, 285, 3, 52, 26, 0, 285, 286, 5, 49,
		0, 0, 286, 287, 3, 2, 1, 0, 287, 288, 5, 50, 0, 0, 288, 289, 6, 15, -1,
		0, 289, 300, 1, 0, 0, 0, 290, 291, 5, 11, 0, 0, 291, 292, 5, 32, 0, 0,
		292, 293, 5, 15, 0, 0, 293, 294, 7, 0, 0, 0, 294, 295, 5, 49, 0, 0, 295,
		296, 3, 2, 1, 0, 296, 297, 5, 50, 0, 0, 297, 298, 6, 15, -1, 0, 298, 300,
		1, 0, 0, 0, 299, 264, 1, 0, 0, 0, 299, 277, 1, 0, 0, 0, 299, 290, 1, 0,
		0, 0, 300, 31, 1, 0, 0, 0, 301, 302, 5, 12, 0, 0, 302, 303, 3, 52, 26,
		0, 303, 304, 5, 49, 0, 0, 304, 305, 3, 34, 17, 0, 305, 306, 5, 14, 0, 0,
		306, 307, 5, 51, 0, 0, 307, 308, 3, 2, 1, 0, 308, 309, 5, 50, 0, 0, 309,
		310, 6, 16, -1, 0, 310, 33, 1, 0, 0, 0, 311, 312, 6, 17, -1, 0, 312, 313,
		3, 36, 18, 0, 313, 314, 6, 17, -1, 0, 314, 321, 1, 0, 0, 0, 315, 316, 10,
		2, 0, 0, 316, 317, 3, 36, 18, 0, 317, 318, 6, 17, -1, 0, 318, 320, 1, 0,
		0, 0, 319, 315, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0,
		321, 322, 1, 0, 0, 0, 322, 35, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 325,
		5, 13, 0, 0, 325, 326, 3, 52, 26, 0, 326, 327, 5, 51, 0, 0, 327, 328, 3,
		2, 1, 0, 328, 329, 6, 18, -1, 0, 329, 37, 1, 0, 0, 0, 330, 331, 5, 8, 0,
		0, 331, 332, 3, 52, 26, 0, 332, 333, 5, 49, 0, 0, 333, 334, 3, 2, 1, 0,
		334, 335, 5, 50, 0, 0, 335, 336, 6, 19, -1, 0, 336, 369, 1, 0, 0, 0, 337,
		338, 5, 8, 0, 0, 338, 339, 3, 52, 26, 0, 339, 340, 5, 49, 0, 0, 340, 341,
		3, 2, 1, 0, 341, 342, 5, 50, 0, 0, 342, 343, 5, 9, 0, 0, 343, 344, 5, 49,
		0, 0, 344, 345, 3, 2, 1, 0, 345, 346, 5, 50, 0, 0, 346, 347, 6, 19, -1,
		0, 347, 369, 1, 0, 0, 0, 348, 349, 5, 8, 0, 0, 349, 350, 3, 52, 26, 0,
		350, 351, 5, 49, 0, 0, 351, 352, 3, 2, 1, 0, 352, 353, 5, 50, 0, 0, 353,
		354, 3, 40, 20, 0, 354, 355, 6, 19, -1, 0, 355, 369, 1, 0, 0, 0, 356, 357,
		5, 8, 0, 0, 357, 358, 3, 52, 26, 0, 358, 359, 5, 49, 0, 0, 359, 360, 3,
		2, 1, 0, 360, 361, 5, 50, 0, 0, 361, 362, 3, 40, 20, 0, 362, 363, 5, 9,
		0, 0, 363, 364, 5, 49, 0, 0, 364, 365, 3, 2, 1, 0, 365, 366, 5, 50, 0,
		0, 366, 367, 6, 19, -1, 0, 367, 369, 1, 0, 0, 0, 368, 330, 1, 0, 0, 0,
		368, 337, 1, 0, 0, 0, 368, 348, 1, 0, 0, 0, 368, 356, 1, 0, 0, 0, 369,
		39, 1, 0, 0, 0, 370, 371, 6, 20, -1, 0, 371, 372, 3, 42, 21, 0, 372, 373,
		6, 20, -1, 0, 373, 380, 1, 0, 0, 0, 374, 375, 10, 2, 0, 0, 375, 376, 3,
		42, 21, 0, 376, 377, 6, 20, -1, 0, 377, 379, 1, 0, 0, 0, 378, 374, 1, 0,
		0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0,
		381, 41, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 384, 5, 9, 0, 0, 384, 385,
		5, 8, 0, 0, 385, 386, 3, 52, 26, 0, 386, 387, 5, 49, 0, 0, 387, 388, 3,
		2, 1, 0, 388, 389, 5, 50, 0, 0, 389, 390, 6, 21, -1, 0, 390, 43, 1, 0,
		0, 0, 391, 392, 5, 10, 0, 0, 392, 393, 3, 52, 26, 0, 393, 394, 5, 49, 0,
		0, 394, 395, 3, 2, 1, 0, 395, 396, 5, 50, 0, 0, 396, 397, 6, 22, -1, 0,
		397, 45, 1, 0, 0, 0, 398, 399, 5, 16, 0, 0, 399, 400, 5, 32, 0, 0, 400,
		401, 5, 51, 0, 0, 401, 402, 3, 50, 25, 0, 402, 403, 5, 38, 0, 0, 403, 404,
		3, 52, 26, 0, 404, 405, 6, 23, -1, 0, 405, 439, 1, 0, 0, 0, 406, 407, 5,
		16, 0, 0, 407, 408, 5, 32, 0, 0, 408, 409, 5, 38, 0, 0, 409, 410, 3, 52,
		26, 0, 410, 411, 6, 23, -1, 0, 411, 439, 1, 0, 0, 0, 412, 413, 5, 17, 0,
		0, 413, 414, 5, 32, 0, 0, 414, 415, 5, 38, 0, 0, 415, 416, 3, 52, 26, 0,
		416, 417, 6, 23, -1, 0, 417, 439, 1, 0, 0, 0, 418, 419, 5, 17, 0, 0, 419,
		420, 5, 32, 0, 0, 420, 421, 5, 51, 0, 0, 421, 422, 3, 50, 25, 0, 422, 423,
		5, 38, 0, 0, 423, 424, 3, 52, 26, 0, 424, 425, 6, 23, -1, 0, 425, 439,
		1, 0, 0, 0, 426, 427, 5, 16, 0, 0, 427, 428, 5, 32, 0, 0, 428, 429, 5,
		38, 0, 0, 429, 430, 3, 52, 26, 0, 430, 431, 6, 23, -1, 0, 431, 439, 1,
		0, 0, 0, 432, 433, 5, 17, 0, 0, 433, 434, 5, 32, 0, 0, 434, 435, 5, 38,
		0, 0, 435, 436, 3, 52, 26, 0, 436, 437, 6, 23, -1, 0, 437, 439, 1, 0, 0,
		0, 438, 398, 1, 0, 0, 0, 438, 406, 1, 0, 0, 0, 438, 412, 1, 0, 0, 0, 438,
		418, 1, 0, 0, 0, 438, 426, 1, 0, 0, 0, 438, 432, 1, 0, 0, 0, 439, 47, 1,
		0, 0, 0, 440, 441, 5, 32, 0, 0, 441, 442, 5, 38, 0, 0, 442, 443, 3, 52,
		26, 0, 443, 444, 6, 24, -1, 0, 444, 452, 1, 0, 0, 0, 445, 446, 5, 32, 0,
		0, 446, 447, 5, 45, 0, 0, 447, 448, 5, 38, 0, 0, 448, 449, 3, 52, 26, 0,
		449, 450, 6, 24, -1, 0, 450, 452, 1, 0, 0, 0, 451, 440, 1, 0, 0, 0, 451,
		445, 1, 0, 0, 0, 452, 49, 1, 0, 0, 0, 453, 454, 5, 1, 0, 0, 454, 472, 6,
		25, -1, 0, 455, 456, 5, 2, 0, 0, 456, 472, 6, 25, -1, 0, 457, 458, 5, 4,
		0, 0, 458, 472, 6, 25, -1, 0, 459, 460, 5, 3, 0, 0, 460, 472, 6, 25, -1,
		0, 461, 462, 5, 52, 0, 0, 462, 463, 3, 50, 25, 0, 463, 464, 5, 53, 0, 0,
		464, 465, 6, 25, -1, 0, 465, 472, 1, 0, 0, 0, 466, 467, 5, 52, 0, 0, 467,
		468, 5, 53, 0, 0, 468, 472, 6, 25, -1, 0, 469, 470, 5, 32, 0, 0, 470, 472,
		6, 25, -1, 0, 471, 453, 1, 0, 0, 0, 471, 455, 1, 0, 0, 0, 471, 457, 1,
		0, 0, 0, 471, 459, 1, 0, 0, 0, 471, 461, 1, 0, 0, 0, 471, 466, 1, 0, 0,
		0, 471, 469, 1, 0, 0, 0, 472, 51, 1, 0, 0, 0, 473, 474, 6, 26, -1, 0, 474,
		475, 5, 46, 0, 0, 475, 476, 3, 52, 26, 23, 476, 477, 6, 26, -1, 0, 477,
		539, 1, 0, 0, 0, 478, 479, 5, 35, 0, 0, 479, 480, 3, 52, 26, 15, 480, 481,
		6, 26, -1, 0, 481, 539, 1, 0, 0, 0, 482, 483, 3, 60, 30, 0, 483, 484, 6,
		26, -1, 0, 484, 539, 1, 0, 0, 0, 485, 486, 5, 32, 0, 0, 486, 487, 5, 47,
		0, 0, 487, 488, 3, 64, 32, 0, 488, 489, 5, 48, 0, 0, 489, 490, 6, 26, -1,
		0, 490, 539, 1, 0, 0, 0, 491, 492, 5, 47, 0, 0, 492, 493, 3, 52, 26, 0,
		493, 494, 5, 48, 0, 0, 494, 495, 6, 26, -1, 0, 495, 539, 1, 0, 0, 0, 496,
		497, 3, 56, 28, 0, 497, 498, 6, 26, -1, 0, 498, 539, 1, 0, 0, 0, 499, 500,
		5, 52, 0, 0, 500, 501, 3, 54, 27, 0, 501, 502, 5, 53, 0, 0, 502, 503, 6,
		26, -1, 0, 503, 539, 1, 0, 0, 0, 504, 505, 5, 1, 0, 0, 505, 506, 5, 47,
		0, 0, 506, 507, 3, 52, 26, 0, 507, 508, 5, 48, 0, 0, 508, 509, 6, 26, -1,
		0, 509, 539, 1, 0, 0, 0, 510, 511, 5, 4, 0, 0, 511, 512, 5, 47, 0, 0, 512,
		513, 3, 52, 26, 0, 513, 514, 5, 48, 0, 0, 514, 515, 6, 26, -1, 0, 515,
		539, 1, 0, 0, 0, 516, 517, 5, 2, 0, 0, 517, 518, 5, 47, 0, 0, 518, 519,
		3, 52, 26, 0, 519, 520, 5, 48, 0, 0, 520, 521, 6, 26, -1, 0, 521, 539,
		1, 0, 0, 0, 522, 523, 5, 32, 0, 0, 523, 524, 5, 56, 0, 0, 524, 525, 5,
		25, 0, 0, 525, 539, 6, 26, -1, 0, 526, 527, 5, 32, 0, 0, 527, 528, 5, 56,
		0, 0, 528, 529, 5, 26, 0, 0, 529, 539, 6, 26, -1, 0, 530, 531, 5, 30, 0,
		0, 531, 539, 6, 26, -1, 0, 532, 533, 5, 31, 0, 0, 533, 539, 6, 26, -1,
		0, 534, 535, 5, 5, 0, 0, 535, 539, 6, 26, -1, 0, 536, 537, 5, 6, 0, 0,
		537, 539, 6, 26, -1, 0, 538, 473, 1, 0, 0, 0, 538, 478, 1, 0, 0, 0, 538,
		482, 1, 0, 0, 0, 538, 485, 1, 0, 0, 0, 538, 491, 1, 0, 0, 0, 538, 496,
		1, 0, 0, 0, 538, 499, 1, 0, 0, 0, 538, 504, 1, 0, 0, 0, 538, 510, 1, 0,
		0, 0, 538, 516, 1, 0, 0, 0, 538, 522, 1, 0, 0, 0, 538, 526, 1, 0, 0, 0,
		538, 530, 1, 0, 0, 0, 538, 532, 1, 0, 0, 0, 538, 534, 1, 0, 0, 0, 538,
		536, 1, 0, 0, 0, 539, 577, 1, 0, 0, 0, 540, 541, 10, 22, 0, 0, 541, 542,
		7, 1, 0, 0, 542, 543, 3, 52, 26, 23, 543, 544, 6, 26, -1, 0, 544, 576,
		1, 0, 0, 0, 545, 546, 10, 21, 0, 0, 546, 547, 7, 2, 0, 0, 547, 548, 3,
		52, 26, 22, 548, 549, 6, 26, -1, 0, 549, 576, 1, 0, 0, 0, 550, 551, 10,
		20, 0, 0, 551, 552, 7, 3, 0, 0, 552, 553, 3, 52, 26, 21, 553, 554, 6, 26,
		-1, 0, 554, 576, 1, 0, 0, 0, 555, 556, 10, 19, 0, 0, 556, 557, 7, 4, 0,
		0, 557, 558, 3, 52, 26, 20, 558, 559, 6, 26, -1, 0, 559, 576, 1, 0, 0,
		0, 560, 561, 10, 18, 0, 0, 561, 562, 7, 5, 0, 0, 562, 563, 3, 52, 26, 19,
		563, 564, 6, 26, -1, 0, 564, 576, 1, 0, 0, 0, 565, 566, 10, 17, 0, 0, 566,
		567, 5, 37, 0, 0, 567, 568, 3, 52, 26, 18, 568, 569, 6, 26, -1, 0, 569,
		576, 1, 0, 0, 0, 570, 571, 10, 16, 0, 0, 571, 572, 5, 36, 0, 0, 572, 573,
		3, 52, 26, 17, 573, 574, 6, 26, -1, 0, 574, 576, 1, 0, 0, 0, 575, 540,
		1, 0, 0, 0, 575, 545, 1, 0, 0, 0, 575, 550, 1, 0, 0, 0, 575, 555, 1, 0,
		0, 0, 575, 560, 1, 0, 0, 0, 575, 565, 1, 0, 0, 0, 575, 570, 1, 0, 0, 0,
		576, 579, 1, 0, 0, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578,
		53, 1, 0, 0, 0, 579, 577, 1, 0, 0, 0, 580, 581, 6, 27, -1, 0, 581, 582,
		3, 52, 26, 0, 582, 583, 6, 27, -1, 0, 583, 591, 1, 0, 0, 0, 584, 585, 10,
		2, 0, 0, 585, 586, 5, 54, 0, 0, 586, 587, 3, 52, 26, 0, 587, 588, 6, 27,
		-1, 0, 588, 590, 1, 0, 0, 0, 589, 584, 1, 0, 0, 0, 590, 593, 1, 0, 0, 0,
		591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 55, 1, 0, 0, 0, 593, 591,
		1, 0, 0, 0, 594, 595, 6, 28, -1, 0, 595, 596, 5, 32, 0, 0, 596, 597, 6,
		28, -1, 0, 597, 610, 1, 0, 0, 0, 598, 599, 10, 3, 0, 0, 599, 600, 5, 52,
		0, 0, 600, 601, 3, 52, 26, 0, 601, 602, 5, 53, 0, 0, 602, 603, 6, 28, -1,
		0, 603, 609, 1, 0, 0, 0, 604, 605, 10, 2, 0, 0, 605, 606, 5, 56, 0, 0,
		606, 607, 5, 32, 0, 0, 607, 609, 6, 28, -1, 0, 608, 598, 1, 0, 0, 0, 608,
		604, 1, 0, 0, 0, 609, 612, 1, 0, 0, 0, 610, 608, 1, 0, 0, 0, 610, 611,
		1, 0, 0, 0, 611, 57, 1, 0, 0, 0, 612, 610, 1, 0, 0, 0, 613, 614, 5, 32,
		0, 0, 614, 615, 5, 47, 0, 0, 615, 616, 3, 54, 27, 0, 616, 617, 5, 48, 0,
		0, 617, 618, 6, 29, -1, 0, 618, 59, 1, 0, 0, 0, 619, 620, 5, 32, 0, 0,
		620, 621, 5, 47, 0, 0, 621, 622, 3, 62, 31, 0, 622, 623, 5, 48, 0, 0, 623,
		624, 6, 30, -1, 0, 624, 61, 1, 0, 0, 0, 625, 626, 6, 31, -1, 0, 626, 627,
		3, 52, 26, 0, 627, 628, 6, 31, -1, 0, 628, 631, 1, 0, 0, 0, 629, 631, 6,
		31, -1, 0, 630, 625, 1, 0, 0, 0, 630, 629, 1, 0, 0, 0, 631, 639, 1, 0,
		0, 0, 632, 633, 10, 3, 0, 0, 633, 634, 5, 54, 0, 0, 634, 635, 3, 52, 26,
		0, 635, 636, 6, 31, -1, 0, 636, 638, 1, 0, 0, 0, 637, 632, 1, 0, 0, 0,
		638, 641, 1, 0, 0, 0, 639, 637, 1, 0, 0, 0, 639, 640, 1, 0, 0, 0, 640,
		63, 1, 0, 0, 0, 641, 639, 1, 0, 0, 0, 642, 643, 6, 32, -1, 0, 643, 644,
		5, 32, 0, 0, 644, 645, 5, 51, 0, 0, 645, 646, 3, 52, 26, 0, 646, 647, 6,
		32, -1, 0, 647, 650, 1, 0, 0, 0, 648, 650, 6, 32, -1, 0, 649, 642, 1, 0,
		0, 0, 649, 648, 1, 0, 0, 0, 650, 660, 1, 0, 0, 0, 651, 652, 10, 3, 0, 0,
		652, 653, 5, 54, 0, 0, 653, 654, 5, 32, 0, 0, 654, 655, 5, 51, 0, 0, 655,
		656, 3, 52, 26, 0, 656, 657, 6, 32, -1, 0, 657, 659, 1, 0, 0, 0, 658, 651,
		1, 0, 0, 0, 659, 662, 1, 0, 0, 0, 660, 658, 1, 0, 0, 0, 660, 661, 1, 0,
		0, 0, 661, 65, 1, 0, 0, 0, 662, 660, 1, 0, 0, 0, 25, 73, 128, 156, 165,
		176, 194, 206, 262, 299, 321, 368, 380, 438, 451, 471, 538, 575, 577, 591,
		608, 610, 630, 639, 649, 660,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserBOOL         = 3
	SwiftGrammarParserSTR          = 4
	SwiftGrammarParserTRU          = 5
	SwiftGrammarParserFAL          = 6
	SwiftGrammarParserPRINT        = 7
	SwiftGrammarParserIF           = 8
	SwiftGrammarParserELSE         = 9
	SwiftGrammarParserWHILE        = 10
	SwiftGrammarParserFOR          = 11
	SwiftGrammarParserSWITCH       = 12
	SwiftGrammarParserCASE         = 13
	SwiftGrammarParserDEFAULT      = 14
	SwiftGrammarParserIN           = 15
	SwiftGrammarParserVAR          = 16
	SwiftGrammarParserLET          = 17
	SwiftGrammarParserBREAK        = 18
	SwiftGrammarParserCONTINUE     = 19
	SwiftGrammarParserRETURN       = 20
	SwiftGrammarParserGUARD        = 21
	SwiftGrammarParserAPPEND       = 22
	SwiftGrammarParserREMOVELAST   = 23
	SwiftGrammarParserREMOVE       = 24
	SwiftGrammarParserCOUNT        = 25
	SwiftGrammarParserISEMPTY      = 26
	SwiftGrammarParserAT           = 27
	SwiftGrammarParserFUNC         = 28
	SwiftGrammarParserSTRUCT       = 29
	SwiftGrammarParserNUMBER       = 30
	SwiftGrammarParserSTRING       = 31
	SwiftGrammarParserID           = 32
	SwiftGrammarParserDIF          = 33
	SwiftGrammarParserIG_IG        = 34
	SwiftGrammarParserNOT          = 35
	SwiftGrammarParserOR           = 36
	SwiftGrammarParserAND          = 37
	SwiftGrammarParserIG           = 38
	SwiftGrammarParserMAY_IG       = 39
	SwiftGrammarParserMEN_IG       = 40
	SwiftGrammarParserMAYOR        = 41
	SwiftGrammarParserMENOR        = 42
	SwiftGrammarParserMUL          = 43
	SwiftGrammarParserDIV          = 44
	SwiftGrammarParserADD          = 45
	SwiftGrammarParserSUB          = 46
	SwiftGrammarParserPARIZQ       = 47
	SwiftGrammarParserPARDER       = 48
	SwiftGrammarParserLLAVEIZQ     = 49
	SwiftGrammarParserLLAVEDER     = 50
	SwiftGrammarParserD_PTS        = 51
	SwiftGrammarParserCORIZQ       = 52
	SwiftGrammarParserCORDER       = 53
	SwiftGrammarParserCOMA         = 54
	SwiftGrammarParserARROW1       = 55
	SwiftGrammarParserPUNTO        = 56
	SwiftGrammarParserMODULO       = 57
	SwiftGrammarParserGUIONB       = 58
	SwiftGrammarParserWHITESPACE   = 59
	SwiftGrammarParserCOMMENT      = 60
	SwiftGrammarParserLINE_COMMENT = 61
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                = 0
	SwiftGrammarParserRULE_block            = 1
	SwiftGrammarParserRULE_instruction      = 2
	SwiftGrammarParserRULE_returntmt        = 3
	SwiftGrammarParserRULE_function         = 4
	SwiftGrammarParserRULE_listParamsFunc   = 5
	SwiftGrammarParserRULE_structCreation   = 6
	SwiftGrammarParserRULE_listStructDec    = 7
	SwiftGrammarParserRULE_removetmt        = 8
	SwiftGrammarParserRULE_removelastmt     = 9
	SwiftGrammarParserRULE_appendtmt        = 10
	SwiftGrammarParserRULE_guardtmt         = 11
	SwiftGrammarParserRULE_breaktmt         = 12
	SwiftGrammarParserRULE_continuetmt      = 13
	SwiftGrammarParserRULE_printstmt        = 14
	SwiftGrammarParserRULE_forstmt          = 15
	SwiftGrammarParserRULE_switchtmt        = 16
	SwiftGrammarParserRULE_caselist         = 17
	SwiftGrammarParserRULE_case             = 18
	SwiftGrammarParserRULE_ifstmt           = 19
	SwiftGrammarParserRULE_elseiflist       = 20
	SwiftGrammarParserRULE_elseif           = 21
	SwiftGrammarParserRULE_whilestmt        = 22
	SwiftGrammarParserRULE_declarationstmt  = 23
	SwiftGrammarParserRULE_asignacion       = 24
	SwiftGrammarParserRULE_types            = 25
	SwiftGrammarParserRULE_expr             = 26
	SwiftGrammarParserRULE_listParams       = 27
	SwiftGrammarParserRULE_listArray        = 28
	SwiftGrammarParserRULE_callFunctionInst = 29
	SwiftGrammarParserRULE_callFunction     = 30
	SwiftGrammarParserRULE_listParamsCall   = 31
	SwiftGrammarParserRULE_listStructExp    = 32
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(67)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(70)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_declarationstmt returns the _declarationstmt rule contexts.
	Get_declarationstmt() IDeclarationstmtContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_switchtmt returns the _switchtmt rule contexts.
	Get_switchtmt() ISwitchtmtContext

	// Get_breaktmt returns the _breaktmt rule contexts.
	Get_breaktmt() IBreaktmtContext

	// Get_continuetmt returns the _continuetmt rule contexts.
	Get_continuetmt() IContinuetmtContext

	// Get_guardtmt returns the _guardtmt rule contexts.
	Get_guardtmt() IGuardtmtContext

	// Get_appendtmt returns the _appendtmt rule contexts.
	Get_appendtmt() IAppendtmtContext

	// Get_removelastmt returns the _removelastmt rule contexts.
	Get_removelastmt() IRemovelastmtContext

	// Get_removetmt returns the _removetmt rule contexts.
	Get_removetmt() IRemovetmtContext

	// Get_function returns the _function rule contexts.
	Get_function() IFunctionContext

	// Get_callFunctionInst returns the _callFunctionInst rule contexts.
	Get_callFunctionInst() ICallFunctionInstContext

	// Get_structCreation returns the _structCreation rule contexts.
	Get_structCreation() IStructCreationContext

	// Get_returntmt returns the _returntmt rule contexts.
	Get_returntmt() IReturntmtContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_declarationstmt sets the _declarationstmt rule contexts.
	Set_declarationstmt(IDeclarationstmtContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_switchtmt sets the _switchtmt rule contexts.
	Set_switchtmt(ISwitchtmtContext)

	// Set_breaktmt sets the _breaktmt rule contexts.
	Set_breaktmt(IBreaktmtContext)

	// Set_continuetmt sets the _continuetmt rule contexts.
	Set_continuetmt(IContinuetmtContext)

	// Set_guardtmt sets the _guardtmt rule contexts.
	Set_guardtmt(IGuardtmtContext)

	// Set_appendtmt sets the _appendtmt rule contexts.
	Set_appendtmt(IAppendtmtContext)

	// Set_removelastmt sets the _removelastmt rule contexts.
	Set_removelastmt(IRemovelastmtContext)

	// Set_removetmt sets the _removetmt rule contexts.
	Set_removetmt(IRemovetmtContext)

	// Set_function sets the _function rule contexts.
	Set_function(IFunctionContext)

	// Set_callFunctionInst sets the _callFunctionInst rule contexts.
	Set_callFunctionInst(ICallFunctionInstContext)

	// Set_structCreation sets the _structCreation rule contexts.
	Set_structCreation(IStructCreationContext)

	// Set_returntmt sets the _returntmt rule contexts.
	Set_returntmt(IReturntmtContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Declarationstmt() IDeclarationstmtContext
	Asignacion() IAsignacionContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Switchtmt() ISwitchtmtContext
	Breaktmt() IBreaktmtContext
	Continuetmt() IContinuetmtContext
	Guardtmt() IGuardtmtContext
	Appendtmt() IAppendtmtContext
	Removelastmt() IRemovelastmtContext
	Removetmt() IRemovetmtContext
	Function() IFunctionContext
	CallFunctionInst() ICallFunctionInstContext
	StructCreation() IStructCreationContext
	Returntmt() IReturntmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	inst              interfaces.Instruction
	_printstmt        IPrintstmtContext
	_ifstmt           IIfstmtContext
	_declarationstmt  IDeclarationstmtContext
	_asignacion       IAsignacionContext
	_whilestmt        IWhilestmtContext
	_forstmt          IForstmtContext
	_switchtmt        ISwitchtmtContext
	_breaktmt         IBreaktmtContext
	_continuetmt      IContinuetmtContext
	_guardtmt         IGuardtmtContext
	_appendtmt        IAppendtmtContext
	_removelastmt     IRemovelastmtContext
	_removetmt        IRemovetmtContext
	_function         IFunctionContext
	_callFunctionInst ICallFunctionInstContext
	_structCreation   IStructCreationContext
	_returntmt        IReturntmtContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_declarationstmt() IDeclarationstmtContext { return s._declarationstmt }

func (s *InstructionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_switchtmt() ISwitchtmtContext { return s._switchtmt }

func (s *InstructionContext) Get_breaktmt() IBreaktmtContext { return s._breaktmt }

func (s *InstructionContext) Get_continuetmt() IContinuetmtContext { return s._continuetmt }

func (s *InstructionContext) Get_guardtmt() IGuardtmtContext { return s._guardtmt }

func (s *InstructionContext) Get_appendtmt() IAppendtmtContext { return s._appendtmt }

func (s *InstructionContext) Get_removelastmt() IRemovelastmtContext { return s._removelastmt }

func (s *InstructionContext) Get_removetmt() IRemovetmtContext { return s._removetmt }

func (s *InstructionContext) Get_function() IFunctionContext { return s._function }

func (s *InstructionContext) Get_callFunctionInst() ICallFunctionInstContext {
	return s._callFunctionInst
}

func (s *InstructionContext) Get_structCreation() IStructCreationContext { return s._structCreation }

func (s *InstructionContext) Get_returntmt() IReturntmtContext { return s._returntmt }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_declarationstmt(v IDeclarationstmtContext) { s._declarationstmt = v }

func (s *InstructionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_switchtmt(v ISwitchtmtContext) { s._switchtmt = v }

func (s *InstructionContext) Set_breaktmt(v IBreaktmtContext) { s._breaktmt = v }

func (s *InstructionContext) Set_continuetmt(v IContinuetmtContext) { s._continuetmt = v }

func (s *InstructionContext) Set_guardtmt(v IGuardtmtContext) { s._guardtmt = v }

func (s *InstructionContext) Set_appendtmt(v IAppendtmtContext) { s._appendtmt = v }

func (s *InstructionContext) Set_removelastmt(v IRemovelastmtContext) { s._removelastmt = v }

func (s *InstructionContext) Set_removetmt(v IRemovetmtContext) { s._removetmt = v }

func (s *InstructionContext) Set_function(v IFunctionContext) { s._function = v }

func (s *InstructionContext) Set_callFunctionInst(v ICallFunctionInstContext) {
	s._callFunctionInst = v
}

func (s *InstructionContext) Set_structCreation(v IStructCreationContext) { s._structCreation = v }

func (s *InstructionContext) Set_returntmt(v IReturntmtContext) { s._returntmt = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Declarationstmt() IDeclarationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationstmtContext)
}

func (s *InstructionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Switchtmt() ISwitchtmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchtmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchtmtContext)
}

func (s *InstructionContext) Breaktmt() IBreaktmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreaktmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreaktmtContext)
}

func (s *InstructionContext) Continuetmt() IContinuetmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuetmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuetmtContext)
}

func (s *InstructionContext) Guardtmt() IGuardtmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardtmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardtmtContext)
}

func (s *InstructionContext) Appendtmt() IAppendtmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendtmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendtmtContext)
}

func (s *InstructionContext) Removelastmt() IRemovelastmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovelastmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovelastmtContext)
}

func (s *InstructionContext) Removetmt() IRemovetmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovetmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovetmtContext)
}

func (s *InstructionContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *InstructionContext) CallFunctionInst() ICallFunctionInstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionInstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionInstContext)
}

func (s *InstructionContext) StructCreation() IStructCreationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCreationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCreationContext)
}

func (s *InstructionContext) Returntmt() IReturntmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturntmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturntmtContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)

			var _x = p.Asignacion()

			localctx.(*InstructionContext)._asignacion = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignacion().GetAsign()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhiles()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(92)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFors()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(95)

			var _x = p.Switchtmt()

			localctx.(*InstructionContext)._switchtmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchtmt().GetSwtch()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(98)

			var _x = p.Breaktmt()

			localctx.(*InstructionContext)._breaktmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_breaktmt().GetBreak_()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(101)

			var _x = p.Continuetmt()

			localctx.(*InstructionContext)._continuetmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_continuetmt().GetContinue_()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(104)

			var _x = p.Guardtmt()

			localctx.(*InstructionContext)._guardtmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardtmt().GetGuard()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(107)

			var _x = p.Appendtmt()

			localctx.(*InstructionContext)._appendtmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_appendtmt().GetAppend_()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(110)

			var _x = p.Removelastmt()

			localctx.(*InstructionContext)._removelastmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removelastmt().GetRemovelast()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(113)

			var _x = p.Removetmt()

			localctx.(*InstructionContext)._removetmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removetmt().GetRemove()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(116)

			var _x = p.Function()

			localctx.(*InstructionContext)._function = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_function().GetFun()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(119)

			var _x = p.CallFunctionInst()

			localctx.(*InstructionContext)._callFunctionInst = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_callFunctionInst().GetCfi()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(122)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(125)

			var _x = p.Returntmt()

			localctx.(*InstructionContext)._returntmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_returntmt().GetRet()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturntmtContext is an interface to support dynamic dispatch.
type IReturntmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetRet returns the ret attribute.
	GetRet() interfaces.Instruction

	// SetRet sets the ret attribute.
	SetRet(interfaces.Instruction)

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturntmtContext differentiates from other interfaces.
	IsReturntmtContext()
}

type ReturntmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	ret     interfaces.Instruction
	_RETURN antlr.Token
	_expr   IExprContext
}

func NewEmptyReturntmtContext() *ReturntmtContext {
	var p = new(ReturntmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_returntmt
	return p
}

func InitEmptyReturntmtContext(p *ReturntmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_returntmt
}

func (*ReturntmtContext) IsReturntmtContext() {}

func NewReturntmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturntmtContext {
	var p = new(ReturntmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_returntmt

	return p
}

func (s *ReturntmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturntmtContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *ReturntmtContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *ReturntmtContext) Get_expr() IExprContext { return s._expr }

func (s *ReturntmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ReturntmtContext) GetRet() interfaces.Instruction { return s.ret }

func (s *ReturntmtContext) SetRet(v interfaces.Instruction) { s.ret = v }

func (s *ReturntmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *ReturntmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturntmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturntmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturntmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterReturntmt(s)
	}
}

func (s *ReturntmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitReturntmt(s)
	}
}

func (p *SwiftGrammarParser) Returntmt() (localctx IReturntmtContext) {
	localctx = NewReturntmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_returntmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)

		var _m = p.Match(SwiftGrammarParserRETURN)

		localctx.(*ReturntmtContext)._RETURN = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)

		var _x = p.expr(0)

		localctx.(*ReturntmtContext)._expr = _x
	}
	localctx.(*ReturntmtContext).ret = instructions.NewReturn((func() int {
		if localctx.(*ReturntmtContext).Get_RETURN() == nil {
			return 0
		} else {
			return localctx.(*ReturntmtContext).Get_RETURN().GetLine()
		}
	}()), (func() int {
		if localctx.(*ReturntmtContext).Get_RETURN() == nil {
			return 0
		} else {
			return localctx.(*ReturntmtContext).Get_RETURN().GetColumn()
		}
	}()), localctx.(*ReturntmtContext).Get_expr().GetE())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsFunc returns the _listParamsFunc rule contexts.
	Get_listParamsFunc() IListParamsFuncContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Set_listParamsFunc sets the _listParamsFunc rule contexts.
	Set_listParamsFunc(IListParamsFuncContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetFun returns the fun attribute.
	GetFun() interfaces.Instruction

	// SetFun sets the fun attribute.
	SetFun(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext
	PARDER() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	ARROW1() antlr.TerminalNode
	Types() ITypesContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	fun             interfaces.Instruction
	_FUNC           antlr.Token
	_ID             antlr.Token
	_listParamsFunc IListParamsFuncContext
	_block          IBlockContext
	_types          ITypesContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_listParamsFunc() IListParamsFuncContext { return s._listParamsFunc }

func (s *FunctionContext) Get_block() IBlockContext { return s._block }

func (s *FunctionContext) Get_types() ITypesContext { return s._types }

func (s *FunctionContext) Set_listParamsFunc(v IListParamsFuncContext) { s._listParamsFunc = v }

func (s *FunctionContext) Set_block(v IBlockContext) { s._block = v }

func (s *FunctionContext) Set_types(v ITypesContext) { s._types = v }

func (s *FunctionContext) GetFun() interfaces.Instruction { return s.fun }

func (s *FunctionContext) SetFun(v interfaces.Instruction) { s.fun = v }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FunctionContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *FunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FunctionContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FunctionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FunctionContext) ARROW1() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserARROW1, 0)
}

func (s *FunctionContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *SwiftGrammarParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_function)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(138)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(141)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*FunctionContext).fun = instructions.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), environment.NULL, localctx.(*FunctionContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(148)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(SwiftGrammarParserARROW1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)

			var _x = p.Types()

			localctx.(*FunctionContext)._types = _x
		}
		{
			p.SetState(151)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(153)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*FunctionContext).fun = instructions.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), localctx.(*FunctionContext).Get_types().GetTy(), localctx.(*FunctionContext).Get_block().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsFuncContext is an interface to support dynamic dispatch.
type IListParamsFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListParamsFuncContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListParamsFuncContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetLpf returns the lpf attribute.
	GetLpf() []interface{}

	// SetLpf sets the lpf attribute.
	SetLpf([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	COMA() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext

	// IsListParamsFuncContext differentiates from other interfaces.
	IsListParamsFuncContext()
}

type ListParamsFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lpf    []interface{}
	list   IListParamsFuncContext
	_ID    antlr.Token
	_types ITypesContext
}

func NewEmptyListParamsFuncContext() *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
	return p
}

func InitEmptyListParamsFuncContext(p *ListParamsFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
}

func (*ListParamsFuncContext) IsListParamsFuncContext() {}

func NewListParamsFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc

	return p
}

func (s *ListParamsFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsFuncContext) Get_ID() antlr.Token { return s._ID }

func (s *ListParamsFuncContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListParamsFuncContext) GetList() IListParamsFuncContext { return s.list }

func (s *ListParamsFuncContext) Get_types() ITypesContext { return s._types }

func (s *ListParamsFuncContext) SetList(v IListParamsFuncContext) { s.list = v }

func (s *ListParamsFuncContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListParamsFuncContext) GetLpf() []interface{} { return s.lpf }

func (s *ListParamsFuncContext) SetLpf(v []interface{}) { s.lpf = v }

func (s *ListParamsFuncContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListParamsFuncContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *ListParamsFuncContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListParamsFuncContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsFuncContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *ListParamsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsFunc(s)
	}
}

func (s *ListParamsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsFunc(s)
	}
}

func (p *SwiftGrammarParser) ListParamsFunc() (localctx IListParamsFuncContext) {
	return p.listParamsFunc(0)
}

func (p *SwiftGrammarParser) listParamsFunc(_p int) (localctx IListParamsFuncContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsFuncContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsFuncContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SwiftGrammarParserRULE_listParamsFunc, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(159)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)

			var _x = p.Types()

			localctx.(*ListParamsFuncContext)._types = _x
		}

		localctx.(*ListParamsFuncContext).lpf = []interface{}{}
		newParam := instructions.NewParamsDeclaration((func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetText()
			}
		}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
		localctx.(*ListParamsFuncContext).lpf = append(localctx.(*ListParamsFuncContext).lpf, newParam)

	case 2:
		localctx.(*ListParamsFuncContext).lpf = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
			localctx.(*ListParamsFuncContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
			p.SetState(167)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(168)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(169)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListParamsFuncContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(170)
				p.Match(SwiftGrammarParserD_PTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(171)

				var _x = p.Types()

				localctx.(*ListParamsFuncContext)._types = _x
			}

			var arr []interface{}
			newParam := instructions.NewParamsDeclaration((func() int {
				if localctx.(*ListParamsFuncContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
				}
			}()), (func() int {
				if localctx.(*ListParamsFuncContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
				}
			}()), (func() string {
				if localctx.(*ListParamsFuncContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListParamsFuncContext).Get_ID().GetText()
				}
			}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
			arr = append(localctx.(*ListParamsFuncContext).GetList().GetLpf(), newParam)
			localctx.(*ListParamsFuncContext).lpf = arr

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructCreationContext is an interface to support dynamic dispatch.
type IStructCreationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRUCT returns the _STRUCT token.
	Get_STRUCT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_STRUCT sets the _STRUCT token.
	Set_STRUCT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listStructDec returns the _listStructDec rule contexts.
	Get_listStructDec() IListStructDecContext

	// Set_listStructDec sets the _listStructDec rule contexts.
	Set_listStructDec(IListStructDecContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	ListStructDec() IListStructDecContext
	LLAVEDER() antlr.TerminalNode

	// IsStructCreationContext differentiates from other interfaces.
	IsStructCreationContext()
}

type StructCreationContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	dec            interfaces.Instruction
	_STRUCT        antlr.Token
	_ID            antlr.Token
	_listStructDec IListStructDecContext
}

func NewEmptyStructCreationContext() *StructCreationContext {
	var p = new(StructCreationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
	return p
}

func InitEmptyStructCreationContext(p *StructCreationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
}

func (*StructCreationContext) IsStructCreationContext() {}

func NewStructCreationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCreationContext {
	var p = new(StructCreationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structCreation

	return p
}

func (s *StructCreationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCreationContext) Get_STRUCT() antlr.Token { return s._STRUCT }

func (s *StructCreationContext) Get_ID() antlr.Token { return s._ID }

func (s *StructCreationContext) Set_STRUCT(v antlr.Token) { s._STRUCT = v }

func (s *StructCreationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructCreationContext) Get_listStructDec() IListStructDecContext { return s._listStructDec }

func (s *StructCreationContext) Set_listStructDec(v IListStructDecContext) { s._listStructDec = v }

func (s *StructCreationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *StructCreationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *StructCreationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *StructCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructCreationContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *StructCreationContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *StructCreationContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (p *SwiftGrammarParser) StructCreation() (localctx IStructCreationContext) {
	localctx = NewStructCreationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_structCreation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)

		var _m = p.Match(SwiftGrammarParserSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructCreationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(183)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructCreationContext).dec = instructions.NewStruct((func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructCreationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructCreationContext).Get_ID().GetText()
		}
	}()), localctx.(*StructCreationContext).Get_listStructDec().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructDecContext is an interface to support dynamic dispatch.
type IListStructDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructDecContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListStructDecContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	COMA() antlr.TerminalNode
	ListStructDec() IListStructDecContext

	// IsListStructDecContext differentiates from other interfaces.
	IsListStructDecContext()
}

type ListStructDecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructDecContext
	_ID    antlr.Token
	_types ITypesContext
}

func NewEmptyListStructDecContext() *ListStructDecContext {
	var p = new(ListStructDecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
	return p
}

func InitEmptyListStructDecContext(p *ListStructDecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
}

func (*ListStructDecContext) IsListStructDecContext() {}

func NewListStructDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructDecContext {
	var p = new(ListStructDecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec

	return p
}

func (s *ListStructDecContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructDecContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructDecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructDecContext) GetList() IListStructDecContext { return s.list }

func (s *ListStructDecContext) Get_types() ITypesContext { return s._types }

func (s *ListStructDecContext) SetList(v IListStructDecContext) { s.list = v }

func (s *ListStructDecContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListStructDecContext) GetL() []interface{} { return s.l }

func (s *ListStructDecContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *ListStructDecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListStructDecContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *ListStructDecContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListStructDecContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructDecContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *ListStructDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructDec(s)
	}
}

func (s *ListStructDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructDec(s)
	}
}

func (p *SwiftGrammarParser) ListStructDec() (localctx IListStructDecContext) {
	return p.listStructDec(0)
}

func (p *SwiftGrammarParser) listStructDec(_p int) (localctx IListStructDecContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructDecContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructDecContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, SwiftGrammarParserRULE_listStructDec, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)

			var _x = p.Types()

			localctx.(*ListStructDecContext)._types = _x
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructDecContext).Get_types().GetTy())
		arr = append(arr, newParams)
		localctx.(*ListStructDecContext).l = arr

	case 2:
		localctx.(*ListStructDecContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructDecContext(p, _parentctx, _parentState)
			localctx.(*ListStructDecContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
			p.SetState(196)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(197)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(198)
				p.Match(SwiftGrammarParserVAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(199)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructDecContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Match(SwiftGrammarParserD_PTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(201)

				var _x = p.Types()

				localctx.(*ListStructDecContext)._types = _x
			}

			var arr []interface{}
			newParams := environment.NewStructType((func() string {
				if localctx.(*ListStructDecContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructDecContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructDecContext).Get_types().GetTy())
			arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
			localctx.(*ListStructDecContext).l = arr

		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemovetmtContext is an interface to support dynamic dispatch.
type IRemovetmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetRemove returns the remove attribute.
	GetRemove() interfaces.Instruction

	// SetRemove sets the remove attribute.
	SetRemove(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AT() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsRemovetmtContext differentiates from other interfaces.
	IsRemovetmtContext()
}

type RemovetmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	remove interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyRemovetmtContext() *RemovetmtContext {
	var p = new(RemovetmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removetmt
	return p
}

func InitEmptyRemovetmtContext(p *RemovetmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removetmt
}

func (*RemovetmtContext) IsRemovetmtContext() {}

func NewRemovetmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovetmtContext {
	var p = new(RemovetmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removetmt

	return p
}

func (s *RemovetmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovetmtContext) Get_ID() antlr.Token { return s._ID }

func (s *RemovetmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemovetmtContext) Get_expr() IExprContext { return s._expr }

func (s *RemovetmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *RemovetmtContext) GetRemove() interfaces.Instruction { return s.remove }

func (s *RemovetmtContext) SetRemove(v interfaces.Instruction) { s.remove = v }

func (s *RemovetmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovetmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *RemovetmtContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *RemovetmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovetmtContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT, 0)
}

func (s *RemovetmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *RemovetmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RemovetmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovetmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovetmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovetmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovetmt(s)
	}
}

func (s *RemovetmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovetmt(s)
	}
}

func (p *SwiftGrammarParser) Removetmt() (localctx IRemovetmtContext) {
	localctx = NewRemovetmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_removetmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemovetmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(SwiftGrammarParserREMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(SwiftGrammarParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)

		var _x = p.expr(0)

		localctx.(*RemovetmtContext)._expr = _x
	}
	{
		p.SetState(216)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*RemovetmtContext).remove = instructions.NewRemove((func() int {
		if localctx.(*RemovetmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovetmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemovetmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovetmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemovetmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemovetmtContext).Get_ID().GetText()
		}
	}()), localctx.(*RemovetmtContext).Get_expr().GetE())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemovelastmtContext is an interface to support dynamic dispatch.
type IRemovelastmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetRemovelast returns the removelast attribute.
	GetRemovelast() interfaces.Instruction

	// SetRemovelast sets the removelast attribute.
	SetRemovelast(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsRemovelastmtContext differentiates from other interfaces.
	IsRemovelastmtContext()
}

type RemovelastmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	removelast interfaces.Instruction
	_ID        antlr.Token
}

func NewEmptyRemovelastmtContext() *RemovelastmtContext {
	var p = new(RemovelastmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastmt
	return p
}

func InitEmptyRemovelastmtContext(p *RemovelastmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastmt
}

func (*RemovelastmtContext) IsRemovelastmtContext() {}

func NewRemovelastmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovelastmtContext {
	var p = new(RemovelastmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removelastmt

	return p
}

func (s *RemovelastmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovelastmtContext) Get_ID() antlr.Token { return s._ID }

func (s *RemovelastmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemovelastmtContext) GetRemovelast() interfaces.Instruction { return s.removelast }

func (s *RemovelastmtContext) SetRemovelast(v interfaces.Instruction) { s.removelast = v }

func (s *RemovelastmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovelastmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *RemovelastmtContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVELAST, 0)
}

func (s *RemovelastmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovelastmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovelastmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovelastmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovelastmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovelastmt(s)
	}
}

func (s *RemovelastmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovelastmt(s)
	}
}

func (p *SwiftGrammarParser) Removelastmt() (localctx IRemovelastmtContext) {
	localctx = NewRemovelastmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_removelastmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemovelastmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(SwiftGrammarParserREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*RemovelastmtContext).removelast = instructions.NewRemoveLast((func() int {
		if localctx.(*RemovelastmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelastmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemovelastmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelastmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemovelastmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemovelastmtContext).Get_ID().GetText()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendtmtContext is an interface to support dynamic dispatch.
type IAppendtmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAppend_ returns the append_ attribute.
	GetAppend_() interfaces.Instruction

	// SetAppend_ sets the append_ attribute.
	SetAppend_(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsAppendtmtContext differentiates from other interfaces.
	IsAppendtmtContext()
}

type AppendtmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	append_ interfaces.Instruction
	_ID     antlr.Token
	_expr   IExprContext
}

func NewEmptyAppendtmtContext() *AppendtmtContext {
	var p = new(AppendtmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendtmt
	return p
}

func InitEmptyAppendtmtContext(p *AppendtmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendtmt
}

func (*AppendtmtContext) IsAppendtmtContext() {}

func NewAppendtmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendtmtContext {
	var p = new(AppendtmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_appendtmt

	return p
}

func (s *AppendtmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendtmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AppendtmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AppendtmtContext) Get_expr() IExprContext { return s._expr }

func (s *AppendtmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AppendtmtContext) GetAppend_() interfaces.Instruction { return s.append_ }

func (s *AppendtmtContext) SetAppend_(v interfaces.Instruction) { s.append_ = v }

func (s *AppendtmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AppendtmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *AppendtmtContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *AppendtmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *AppendtmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppendtmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *AppendtmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendtmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendtmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAppendtmt(s)
	}
}

func (s *AppendtmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAppendtmt(s)
	}
}

func (p *SwiftGrammarParser) Appendtmt() (localctx IAppendtmtContext) {
	localctx = NewAppendtmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_appendtmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AppendtmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(SwiftGrammarParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)

		var _x = p.expr(0)

		localctx.(*AppendtmtContext)._expr = _x
	}
	{
		p.SetState(231)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*AppendtmtContext).append_ = instructions.NewAppend((func() int {
		if localctx.(*AppendtmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendtmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AppendtmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendtmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AppendtmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AppendtmtContext).Get_ID().GetText()
		}
	}()), localctx.(*AppendtmtContext).Get_expr().GetE())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardtmtContext is an interface to support dynamic dispatch.
type IGuardtmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGuard returns the guard attribute.
	GetGuard() interfaces.Instruction

	// SetGuard sets the guard attribute.
	SetGuard(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardtmtContext differentiates from other interfaces.
	IsGuardtmtContext()
}

type GuardtmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	guard  interfaces.Instruction
	_GUARD antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyGuardtmtContext() *GuardtmtContext {
	var p = new(GuardtmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardtmt
	return p
}

func InitEmptyGuardtmtContext(p *GuardtmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardtmt
}

func (*GuardtmtContext) IsGuardtmtContext() {}

func NewGuardtmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardtmtContext {
	var p = new(GuardtmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardtmt

	return p
}

func (s *GuardtmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardtmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardtmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardtmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardtmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardtmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardtmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardtmtContext) GetGuard() interfaces.Instruction { return s.guard }

func (s *GuardtmtContext) SetGuard(v interfaces.Instruction) { s.guard = v }

func (s *GuardtmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardtmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardtmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardtmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardtmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardtmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardtmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardtmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardtmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardtmt(s)
	}
}

func (s *GuardtmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardtmt(s)
	}
}

func (p *SwiftGrammarParser) Guardtmt() (localctx IGuardtmtContext) {
	localctx = NewGuardtmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_guardtmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardtmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)

		var _x = p.expr(0)

		localctx.(*GuardtmtContext)._expr = _x
	}
	{
		p.SetState(236)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)

		var _x = p.Block()

		localctx.(*GuardtmtContext)._block = _x
	}
	{
		p.SetState(239)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardtmtContext).guard = instructions.NewGuard((func() int {
		if localctx.(*GuardtmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardtmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardtmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardtmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardtmtContext).Get_expr().GetE(), localctx.(*GuardtmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreaktmtContext is an interface to support dynamic dispatch.
type IBreaktmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// GetBreak_ returns the break_ attribute.
	GetBreak_() interfaces.Instruction

	// SetBreak_ sets the break_ attribute.
	SetBreak_(interfaces.Instruction)

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreaktmtContext differentiates from other interfaces.
	IsBreaktmtContext()
}

type BreaktmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	break_ interfaces.Instruction
	_BREAK antlr.Token
}

func NewEmptyBreaktmtContext() *BreaktmtContext {
	var p = new(BreaktmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breaktmt
	return p
}

func InitEmptyBreaktmtContext(p *BreaktmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breaktmt
}

func (*BreaktmtContext) IsBreaktmtContext() {}

func NewBreaktmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreaktmtContext {
	var p = new(BreaktmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_breaktmt

	return p
}

func (s *BreaktmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreaktmtContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *BreaktmtContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *BreaktmtContext) GetBreak_() interfaces.Instruction { return s.break_ }

func (s *BreaktmtContext) SetBreak_(v interfaces.Instruction) { s.break_ = v }

func (s *BreaktmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *BreaktmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreaktmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreaktmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBreaktmt(s)
	}
}

func (s *BreaktmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBreaktmt(s)
	}
}

func (p *SwiftGrammarParser) Breaktmt() (localctx IBreaktmtContext) {
	localctx = NewBreaktmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_breaktmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)

		var _m = p.Match(SwiftGrammarParserBREAK)

		localctx.(*BreaktmtContext)._BREAK = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*BreaktmtContext).break_ = instructions.NewBreak((func() int {
		if localctx.(*BreaktmtContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreaktmtContext).Get_BREAK().GetLine()
		}
	}()), (func() int {
		if localctx.(*BreaktmtContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreaktmtContext).Get_BREAK().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinuetmtContext is an interface to support dynamic dispatch.
type IContinuetmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetContinue_ returns the continue_ attribute.
	GetContinue_() interfaces.Instruction

	// SetContinue_ sets the continue_ attribute.
	SetContinue_(interfaces.Instruction)

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinuetmtContext differentiates from other interfaces.
	IsContinuetmtContext()
}

type ContinuetmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	continue_ interfaces.Instruction
	_CONTINUE antlr.Token
}

func NewEmptyContinuetmtContext() *ContinuetmtContext {
	var p = new(ContinuetmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuetmt
	return p
}

func InitEmptyContinuetmtContext(p *ContinuetmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuetmt
}

func (*ContinuetmtContext) IsContinuetmtContext() {}

func NewContinuetmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuetmtContext {
	var p = new(ContinuetmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_continuetmt

	return p
}

func (s *ContinuetmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuetmtContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *ContinuetmtContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *ContinuetmtContext) GetContinue_() interfaces.Instruction { return s.continue_ }

func (s *ContinuetmtContext) SetContinue_(v interfaces.Instruction) { s.continue_ = v }

func (s *ContinuetmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *ContinuetmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuetmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuetmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterContinuetmt(s)
	}
}

func (s *ContinuetmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitContinuetmt(s)
	}
}

func (p *SwiftGrammarParser) Continuetmt() (localctx IContinuetmtContext) {
	localctx = NewContinuetmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_continuetmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)

		var _m = p.Match(SwiftGrammarParserCONTINUE)

		localctx.(*ContinuetmtContext)._CONTINUE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ContinuetmtContext).continue_ = instructions.NewContinue((func() int {
		if localctx.(*ContinuetmtContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuetmtContext).Get_CONTINUE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ContinuetmtContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuetmtContext).Get_CONTINUE().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetExpr1 returns the expr1 rule contexts.
	GetExpr1() IExprContext

	// GetExpr2 returns the expr2 rule contexts.
	GetExpr2() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetExpr1 sets the expr1 rule contexts.
	SetExpr1(IExprContext)

	// SetExpr2 sets the expr2 rule contexts.
	SetExpr2(IExprContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	COMA() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prnt   interfaces.Instruction
	_PRINT antlr.Token
	_expr  IExprContext
	expr1  IExprContext
	expr2  IExprContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) GetExpr1() IExprContext { return s.expr1 }

func (s *PrintstmtContext) GetExpr2() IExprContext { return s.expr2 }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) SetExpr1(v IExprContext) { s.expr1 = v }

func (s *PrintstmtContext) SetExpr2(v IExprContext) { s.expr2 = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_printstmt)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)

			var _x = p.expr(0)

			localctx.(*PrintstmtContext)._expr = _x
		}
		{
			p.SetState(251)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).Get_expr().GetE(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)

			var _x = p.expr(0)

			localctx.(*PrintstmtContext).expr1 = _x
		}
		{
			p.SetState(257)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _x = p.expr(0)

			localctx.(*PrintstmtContext).expr2 = _x
		}
		{
			p.SetState(259)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).GetExpr1().GetE(), localctx.(*PrintstmtContext).GetExpr2().GetE())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_GUIONB returns the _GUIONB token.
	Get_GUIONB() antlr.Token

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_GUIONB sets the _GUIONB token.
	Set_GUIONB(antlr.Token)

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetFors returns the fors attribute.
	GetFors() interfaces.Instruction

	// SetFors sets the fors attribute.
	SetFors(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	IN() antlr.TerminalNode
	AllPUNTO() []antlr.TerminalNode
	PUNTO(i int) antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	GUIONB() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	fors    interfaces.Instruction
	_FOR    antlr.Token
	_ID     antlr.Token
	e1      IExprContext
	e2      IExprContext
	_block  IBlockContext
	_GUIONB antlr.Token
	ope     antlr.Token
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Get_GUIONB() antlr.Token { return s._GUIONB }

func (s *ForstmtContext) GetOpe() antlr.Token { return s.ope }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Set_GUIONB(v antlr.Token) { s._GUIONB = v }

func (s *ForstmtContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *ForstmtContext) GetE1() IExprContext { return s.e1 }

func (s *ForstmtContext) GetE2() IExprContext { return s.e2 }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ForstmtContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) GetFors() interfaces.Instruction { return s.fors }

func (s *ForstmtContext) SetFors(v interfaces.Instruction) { s.fors = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ForstmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPUNTO)
}

func (s *ForstmtContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, i)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForstmtContext) GUIONB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUIONB, 0)
}

func (s *ForstmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_forstmt)
	var _la int

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e1 = _x
		}
		{
			p.SetState(268)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e2 = _x
		}
		{
			p.SetState(272)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(274)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).fors = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ForstmtContext).GetE1().GetE(), localctx.(*ForstmtContext).GetE2().GetE(), "nil", localctx.(*ForstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)

			var _m = p.Match(SwiftGrammarParserGUIONB)

			localctx.(*ForstmtContext)._GUIONB = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e1 = _x
		}
		{
			p.SetState(281)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e2 = _x
		}
		{
			p.SetState(285)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(287)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).fors = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_GUIONB() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_GUIONB().GetText()
			}
		}()), localctx.(*ForstmtContext).GetE1().GetE(), localctx.(*ForstmtContext).GetE2().GetE(), "nil", localctx.(*ForstmtContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(290)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ForstmtContext).ope = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserSTRING || _la == SwiftGrammarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ForstmtContext).ope = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(294)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(296)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).fors = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), nil, nil, (func() string {
			if localctx.(*ForstmtContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).GetOpe().GetText()
			}
		}()), localctx.(*ForstmtContext).Get_block().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchtmtContext is an interface to support dynamic dispatch.
type ISwitchtmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_caselist returns the _caselist rule contexts.
	Get_caselist() ICaselistContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_caselist sets the _caselist rule contexts.
	Set_caselist(ICaselistContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetSwtch returns the swtch attribute.
	GetSwtch() interfaces.Instruction

	// SetSwtch sets the swtch attribute.
	SetSwtch(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Caselist() ICaselistContext
	DEFAULT() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsSwitchtmtContext differentiates from other interfaces.
	IsSwitchtmtContext()
}

type SwitchtmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	swtch     interfaces.Instruction
	_SWITCH   antlr.Token
	_expr     IExprContext
	_caselist ICaselistContext
	_block    IBlockContext
}

func NewEmptySwitchtmtContext() *SwitchtmtContext {
	var p = new(SwitchtmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchtmt
	return p
}

func InitEmptySwitchtmtContext(p *SwitchtmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchtmt
}

func (*SwitchtmtContext) IsSwitchtmtContext() {}

func NewSwitchtmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchtmtContext {
	var p = new(SwitchtmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchtmt

	return p
}

func (s *SwitchtmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchtmtContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchtmtContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchtmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchtmtContext) Get_caselist() ICaselistContext { return s._caselist }

func (s *SwitchtmtContext) Get_block() IBlockContext { return s._block }

func (s *SwitchtmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchtmtContext) Set_caselist(v ICaselistContext) { s._caselist = v }

func (s *SwitchtmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *SwitchtmtContext) GetSwtch() interfaces.Instruction { return s.swtch }

func (s *SwitchtmtContext) SetSwtch(v interfaces.Instruction) { s.swtch = v }

func (s *SwitchtmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchtmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchtmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchtmtContext) Caselist() ICaselistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaselistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaselistContext)
}

func (s *SwitchtmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *SwitchtmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *SwitchtmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchtmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchtmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchtmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchtmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchtmt(s)
	}
}

func (s *SwitchtmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchtmt(s)
	}
}

func (p *SwiftGrammarParser) Switchtmt() (localctx ISwitchtmtContext) {
	localctx = NewSwitchtmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_switchtmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)

		var _m = p.Match(SwiftGrammarParserSWITCH)

		localctx.(*SwitchtmtContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)

		var _x = p.expr(0)

		localctx.(*SwitchtmtContext)._expr = _x
	}
	{
		p.SetState(303)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)

		var _x = p.caselist(0)

		localctx.(*SwitchtmtContext)._caselist = _x
	}
	{
		p.SetState(305)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)

		var _x = p.Block()

		localctx.(*SwitchtmtContext)._block = _x
	}
	{
		p.SetState(308)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*SwitchtmtContext).swtch = instructions.NewSwitch((func() int {
		if localctx.(*SwitchtmtContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*SwitchtmtContext).Get_SWITCH().GetLine()
		}
	}()), (func() int {
		if localctx.(*SwitchtmtContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*SwitchtmtContext).Get_SWITCH().GetColumn()
		}
	}()), localctx.(*SwitchtmtContext).Get_expr().GetE(), localctx.(*SwitchtmtContext).Get_caselist().GetCases(), localctx.(*SwitchtmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaselistContext is an interface to support dynamic dispatch.
type ICaselistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetListcases returns the listcases rule contexts.
	GetListcases() ICaselistContext

	// Get_case returns the _case rule contexts.
	Get_case() ICaseContext

	// SetListcases sets the listcases rule contexts.
	SetListcases(ICaselistContext)

	// Set_case sets the _case rule contexts.
	Set_case(ICaseContext)

	// GetCases returns the cases attribute.
	GetCases() []interface{}

	// SetCases sets the cases attribute.
	SetCases([]interface{})

	// Getter signatures
	Case_() ICaseContext
	Caselist() ICaselistContext

	// IsCaselistContext differentiates from other interfaces.
	IsCaselistContext()
}

type CaselistContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	cases     []interface{}
	listcases ICaselistContext
	_case     ICaseContext
}

func NewEmptyCaselistContext() *CaselistContext {
	var p = new(CaselistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caselist
	return p
}

func InitEmptyCaselistContext(p *CaselistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caselist
}

func (*CaselistContext) IsCaselistContext() {}

func NewCaselistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaselistContext {
	var p = new(CaselistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_caselist

	return p
}

func (s *CaselistContext) GetParser() antlr.Parser { return s.parser }

func (s *CaselistContext) GetListcases() ICaselistContext { return s.listcases }

func (s *CaselistContext) Get_case() ICaseContext { return s._case }

func (s *CaselistContext) SetListcases(v ICaselistContext) { s.listcases = v }

func (s *CaselistContext) Set_case(v ICaseContext) { s._case = v }

func (s *CaselistContext) GetCases() []interface{} { return s.cases }

func (s *CaselistContext) SetCases(v []interface{}) { s.cases = v }

func (s *CaselistContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *CaselistContext) Caselist() ICaselistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaselistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaselistContext)
}

func (s *CaselistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaselistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaselistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCaselist(s)
	}
}

func (s *CaselistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCaselist(s)
	}
}

func (p *SwiftGrammarParser) Caselist() (localctx ICaselistContext) {
	return p.caselist(0)
}

func (p *SwiftGrammarParser) caselist(_p int) (localctx ICaselistContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCaselistContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICaselistContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, SwiftGrammarParserRULE_caselist, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)

		var _x = p.Case_()

		localctx.(*CaselistContext)._case = _x
	}

	localctx.(*CaselistContext).cases = []interface{}{}
	localctx.(*CaselistContext).cases = append(localctx.(*CaselistContext).cases, localctx.(*CaselistContext).Get_case().GetSwtch())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCaselistContext(p, _parentctx, _parentState)
			localctx.(*CaselistContext).listcases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_caselist)
			p.SetState(315)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(316)

				var _x = p.Case_()

				localctx.(*CaselistContext)._case = _x
			}
			var arr3 []interface{}
			arr3 = append(localctx.(*CaselistContext).GetListcases().GetCases(), localctx.(*CaselistContext).Get_case().GetSwtch())
			localctx.(*CaselistContext).cases = arr3

		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetSwtch returns the swtch attribute.
	GetSwtch() interfaces.Instruction

	// SetSwtch sets the swtch attribute.
	SetSwtch(interfaces.Instruction)

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	D_PTS() antlr.TerminalNode
	Block() IBlockContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	swtch  interfaces.Instruction
	_CASE  antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CaseContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CaseContext) Get_expr() IExprContext { return s._expr }

func (s *CaseContext) Get_block() IBlockContext { return s._block }

func (s *CaseContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CaseContext) Set_block(v IBlockContext) { s._block = v }

func (s *CaseContext) GetSwtch() interfaces.Instruction { return s.swtch }

func (s *CaseContext) SetSwtch(v interfaces.Instruction) { s.swtch = v }

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *CaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCase(s)
	}
}

func (p *SwiftGrammarParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CaseContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)

		var _x = p.expr(0)

		localctx.(*CaseContext)._expr = _x
	}
	{
		p.SetState(326)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)

		var _x = p.Block()

		localctx.(*CaseContext)._block = _x
	}
	localctx.(*CaseContext).swtch = instructions.NewCase((func() int {
		if localctx.(*CaseContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CaseContext).Get_CASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CaseContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CaseContext).Get_CASE().GetColumn()
		}
	}()), localctx.(*CaseContext).Get_expr().GetE(), localctx.(*CaseContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetB1 returns the b1 rule contexts.
	GetB1() IBlockContext

	// GetB2 returns the b2 rule contexts.
	GetB2() IBlockContext

	// Get_elseiflist returns the _elseiflist rule contexts.
	Get_elseiflist() IElseiflistContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetB1 sets the b1 rule contexts.
	SetB1(IBlockContext)

	// SetB2 sets the b2 rule contexts.
	SetB2(IBlockContext)

	// Set_elseiflist sets the _elseiflist rule contexts.
	Set_elseiflist(IElseiflistContext)

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Elseiflist() IElseiflistContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	ifinst      interfaces.Instruction
	_IF         antlr.Token
	_expr       IExprContext
	_block      IBlockContext
	b1          IBlockContext
	b2          IBlockContext
	_elseiflist IElseiflistContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) Get_block() IBlockContext { return s._block }

func (s *IfstmtContext) GetB1() IBlockContext { return s.b1 }

func (s *IfstmtContext) GetB2() IBlockContext { return s.b2 }

func (s *IfstmtContext) Get_elseiflist() IElseiflistContext { return s._elseiflist }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) SetB1(v IBlockContext) { s.b1 = v }

func (s *IfstmtContext) SetB2(v IBlockContext) { s.b2 = v }

func (s *IfstmtContext) Set_elseiflist(v IElseiflistContext) { s._elseiflist = v }

func (s *IfstmtContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *IfstmtContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) Elseiflist() IElseiflistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseiflistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseiflistContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_ifstmt)
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(332)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(334)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(339)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)

			var _x = p.Block()

			localctx.(*IfstmtContext).b1 = _x
		}
		{
			p.SetState(341)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)

			var _x = p.Block()

			localctx.(*IfstmtContext).b2 = _x
		}
		{
			p.SetState(345)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetB1().GetBlk(), nil, localctx.(*IfstmtContext).GetB2().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(350)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(352)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)

			var _x = p.elseiflist(0)

			localctx.(*IfstmtContext)._elseiflist = _x
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), localctx.(*IfstmtContext).Get_elseiflist().GetElse_if_list(), nil)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(356)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(358)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)

			var _x = p.Block()

			localctx.(*IfstmtContext).b1 = _x
		}
		{
			p.SetState(360)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)

			var _x = p.elseiflist(0)

			localctx.(*IfstmtContext)._elseiflist = _x
		}
		{
			p.SetState(362)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _x = p.Block()

			localctx.(*IfstmtContext).b2 = _x
		}
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetB1().GetBlk(), localctx.(*IfstmtContext).Get_elseiflist().GetElse_if_list(), localctx.(*IfstmtContext).GetB2().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseiflistContext is an interface to support dynamic dispatch.
type IElseiflistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetListif returns the listif rule contexts.
	GetListif() IElseiflistContext

	// Get_elseif returns the _elseif rule contexts.
	Get_elseif() IElseifContext

	// SetListif sets the listif rule contexts.
	SetListif(IElseiflistContext)

	// Set_elseif sets the _elseif rule contexts.
	Set_elseif(IElseifContext)

	// GetElse_if_list returns the else_if_list attribute.
	GetElse_if_list() []interface{}

	// SetElse_if_list sets the else_if_list attribute.
	SetElse_if_list([]interface{})

	// Getter signatures
	Elseif() IElseifContext
	Elseiflist() IElseiflistContext

	// IsElseiflistContext differentiates from other interfaces.
	IsElseiflistContext()
}

type ElseiflistContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	else_if_list []interface{}
	listif       IElseiflistContext
	_elseif      IElseifContext
}

func NewEmptyElseiflistContext() *ElseiflistContext {
	var p = new(ElseiflistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseiflist
	return p
}

func InitEmptyElseiflistContext(p *ElseiflistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseiflist
}

func (*ElseiflistContext) IsElseiflistContext() {}

func NewElseiflistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseiflistContext {
	var p = new(ElseiflistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elseiflist

	return p
}

func (s *ElseiflistContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseiflistContext) GetListif() IElseiflistContext { return s.listif }

func (s *ElseiflistContext) Get_elseif() IElseifContext { return s._elseif }

func (s *ElseiflistContext) SetListif(v IElseiflistContext) { s.listif = v }

func (s *ElseiflistContext) Set_elseif(v IElseifContext) { s._elseif = v }

func (s *ElseiflistContext) GetElse_if_list() []interface{} { return s.else_if_list }

func (s *ElseiflistContext) SetElse_if_list(v []interface{}) { s.else_if_list = v }

func (s *ElseiflistContext) Elseif() IElseifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifContext)
}

func (s *ElseiflistContext) Elseiflist() IElseiflistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseiflistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseiflistContext)
}

func (s *ElseiflistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseiflistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseiflistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElseiflist(s)
	}
}

func (s *ElseiflistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElseiflist(s)
	}
}

func (p *SwiftGrammarParser) Elseiflist() (localctx IElseiflistContext) {
	return p.elseiflist(0)
}

func (p *SwiftGrammarParser) elseiflist(_p int) (localctx IElseiflistContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewElseiflistContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElseiflistContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, SwiftGrammarParserRULE_elseiflist, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)

		var _x = p.Elseif()

		localctx.(*ElseiflistContext)._elseif = _x
	}

	localctx.(*ElseiflistContext).else_if_list = []interface{}{}
	localctx.(*ElseiflistContext).else_if_list = append(localctx.(*ElseiflistContext).else_if_list, localctx.(*ElseiflistContext).Get_elseif().GetIfinst())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElseiflistContext(p, _parentctx, _parentState)
			localctx.(*ElseiflistContext).listif = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_elseiflist)
			p.SetState(374)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(375)

				var _x = p.Elseif()

				localctx.(*ElseiflistContext)._elseif = _x
			}
			var arr2 []interface{}
			arr2 = append(localctx.(*ElseiflistContext).GetListif().GetElse_if_list(), localctx.(*ElseiflistContext).Get_elseif().GetIfinst())
			localctx.(*ElseiflistContext).else_if_list = arr2

		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseifContext is an interface to support dynamic dispatch.
type IElseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElseifContext differentiates from other interfaces.
	IsElseifContext()
}

type ElseifContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ifinst interfaces.Instruction
	_IF    antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyElseifContext() *ElseifContext {
	var p = new(ElseifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseif
	return p
}

func InitEmptyElseifContext(p *ElseifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseif
}

func (*ElseifContext) IsElseifContext() {}

func NewElseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifContext {
	var p = new(ElseifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elseif

	return p
}

func (s *ElseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifContext) Get_IF() antlr.Token { return s._IF }

func (s *ElseifContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *ElseifContext) Get_expr() IExprContext { return s._expr }

func (s *ElseifContext) Get_block() IBlockContext { return s._block }

func (s *ElseifContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ElseifContext) Set_block(v IBlockContext) { s._block = v }

func (s *ElseifContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *ElseifContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *ElseifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElseifContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *ElseifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseifContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElseifContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseifContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElseif(s)
	}
}

func (s *ElseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElseif(s)
	}
}

func (p *SwiftGrammarParser) Elseif() (localctx IElseifContext) {
	localctx = NewElseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_elseif)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)

		var _m = p.Match(SwiftGrammarParserIF)

		localctx.(*ElseifContext)._IF = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)

		var _x = p.expr(0)

		localctx.(*ElseifContext)._expr = _x
	}
	{
		p.SetState(386)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)

		var _x = p.Block()

		localctx.(*ElseifContext)._block = _x
	}
	{
		p.SetState(388)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ElseifContext).ifinst = instructions.NewIf((func() int {
		if localctx.(*ElseifContext).Get_IF() == nil {
			return 0
		} else {
			return localctx.(*ElseifContext).Get_IF().GetLine()
		}
	}()), (func() int {
		if localctx.(*ElseifContext).Get_IF() == nil {
			return 0
		} else {
			return localctx.(*ElseifContext).Get_IF().GetColumn()
		}
	}()), localctx.(*ElseifContext).Get_expr().GetE(), localctx.(*ElseifContext).Get_block().GetBlk(), nil, nil)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhiles returns the whiles attribute.
	GetWhiles() interfaces.Instruction

	// SetWhiles sets the whiles attribute.
	SetWhiles(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	whiles interfaces.Instruction
	_WHILE antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhiles() interfaces.Instruction { return s.whiles }

func (s *WhilestmtContext) SetWhiles(v interfaces.Instruction) { s.whiles = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(393)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(395)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).whiles = instructions.NewWhiles((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationstmtContext is an interface to support dynamic dispatch.
type IDeclarationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	LET() antlr.TerminalNode

	// IsDeclarationstmtContext differentiates from other interfaces.
	IsDeclarationstmtContext()
}

type DeclarationstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    interfaces.Instruction
	_VAR   antlr.Token
	_ID    antlr.Token
	_types ITypesContext
	_expr  IExprContext
	_LET   antlr.Token
}

func NewEmptyDeclarationstmtContext() *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
	return p
}

func InitEmptyDeclarationstmtContext(p *DeclarationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
}

func (*DeclarationstmtContext) IsDeclarationstmtContext() {}

func NewDeclarationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt

	return p
}

func (s *DeclarationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclarationstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarationstmtContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclarationstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarationstmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclarationstmtContext) Get_types() ITypesContext { return s._types }

func (s *DeclarationstmtContext) Get_expr() IExprContext { return s._expr }

func (s *DeclarationstmtContext) Set_types(v ITypesContext) { s._types = v }

func (s *DeclarationstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclarationstmtContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarationstmtContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarationstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *DeclarationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DeclarationstmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *DeclarationstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *DeclarationstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DeclarationstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *DeclarationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclarationstmt(s)
	}
}

func (s *DeclarationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclarationstmt(s)
	}
}

func (p *SwiftGrammarParser) Declarationstmt() (localctx IDeclarationstmtContext) {
	localctx = NewDeclarationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(402)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE(), false))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*DeclarationstmtContext).Get_expr().GetE(), false))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*DeclarationstmtContext).Get_expr().GetE(), true))

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(418)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(422)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE(), true))

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(426)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.STRUCT, localctx.(*DeclarationstmtContext).Get_expr().GetE(), false))

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(432)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).SetDec(instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.STRUCT, localctx.(*DeclarationstmtContext).Get_expr().GetE(), true))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_IG returns the _IG token.
	Get_IG() antlr.Token

	// Get_ADD returns the _ADD token.
	Get_ADD() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_IG sets the _IG token.
	Set_IG(antlr.Token)

	// Set_ADD sets the _ADD token.
	Set_ADD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAsign returns the asign attribute.
	GetAsign() interfaces.Instruction

	// SetAsign sets the asign attribute.
	SetAsign(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	ADD() antlr.TerminalNode

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	asign  interfaces.Instruction
	_ID    antlr.Token
	_IG    antlr.Token
	_expr  IExprContext
	_ADD   antlr.Token
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Get_IG() antlr.Token { return s._IG }

func (s *AsignacionContext) Get_ADD() antlr.Token { return s._ADD }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Set_IG(v antlr.Token) { s._IG = v }

func (s *AsignacionContext) Set_ADD(v antlr.Token) { s._ADD = v }

func (s *AsignacionContext) Get_expr() IExprContext { return s._expr }

func (s *AsignacionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AsignacionContext) GetAsign() interfaces.Instruction { return s.asign }

func (s *AsignacionContext) SetAsign(v interfaces.Instruction) { s.asign = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignacionContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AsignacionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *SwiftGrammarParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_asignacion)
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignacionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)

			var _m = p.Match(SwiftGrammarParserIG)

			localctx.(*AsignacionContext)._IG = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)

			var _x = p.expr(0)

			localctx.(*AsignacionContext)._expr = _x
		}
		localctx.(*AsignacionContext).asign = instructions.NewAsignacion((func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignacionContext).Get_expr().GetE(), (func() string {
			if localctx.(*AsignacionContext).Get_IG() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_IG().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(445)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignacionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)

			var _m = p.Match(SwiftGrammarParserADD)

			localctx.(*AsignacionContext)._ADD = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)

			var _m = p.Match(SwiftGrammarParserIG)

			localctx.(*AsignacionContext)._IG = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)

			var _x = p.expr(0)

			localctx.(*AsignacionContext)._expr = _x
		}
		localctx.(*AsignacionContext).asign = instructions.NewAsignacion((func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignacionContext).Get_expr().GetE(), (func() string {
			if localctx.(*AsignacionContext).Get_ADD() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_ADD().GetText()
			}
		}())+(func() string {
			if localctx.(*AsignacionContext).Get_IG() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_IG().GetText()
			}
		}()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty attribute.
	GetTy() environment.TipoExpresion

	// SetTy sets the ty attribute.
	SetTy(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Types() ITypesContext
	CORDER() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     environment.TipoExpresion
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) GetTy() environment.TipoExpresion { return s.ty }

func (s *TypesContext) SetTy(v environment.TipoExpresion) { s.ty = v }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TypesContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TypesContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *TypesContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TypesContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *TypesContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *SwiftGrammarParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_types)
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(457)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRING

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(461)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Types()
		}
		{
			p.SetState(463)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(466)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(469)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRUCT

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_STR returns the _STR token.
	Get_STR() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_STR sets the _STR token.
	Set_STR(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_callFunction returns the _callFunction rule contexts.
	Get_callFunction() ICallFunctionContext

	// Get_listStructExp returns the _listStructExp rule contexts.
	Get_listStructExp() IListStructExpContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_callFunction sets the _callFunction rule contexts.
	Set_callFunction(ICallFunctionContext)

	// Set_listStructExp sets the _listStructExp rule contexts.
	Set_listStructExp(IListStructExpContext)

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	SUB() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	CallFunction() ICallFunctionContext
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListStructExp() IListStructExpContext
	PARDER() antlr.TerminalNode
	ListArray() IListArrayContext
	CORIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	CORDER() antlr.TerminalNode
	INT() antlr.TerminalNode
	STR() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MODULO() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	e              interfaces.Expression
	left           IExprContext
	op             antlr.Token
	right          IExprContext
	_expr          IExprContext
	_callFunction  ICallFunctionContext
	_ID            antlr.Token
	_listStructExp IListStructExpContext
	list           IListArrayContext
	_CORIZQ        antlr.Token
	_listParams    IListParamsContext
	_INT           antlr.Token
	_STR           antlr.Token
	_FLOAT         antlr.Token
	_NUMBER        antlr.Token
	_STRING        antlr.Token
	_TRU           antlr.Token
	_FAL           antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *ExprContext) Get_INT() antlr.Token { return s._INT }

func (s *ExprContext) Get_STR() antlr.Token { return s._STR }

func (s *ExprContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *ExprContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *ExprContext) Set_STR(v antlr.Token) { s._STR = v }

func (s *ExprContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_callFunction() ICallFunctionContext { return s._callFunction }

func (s *ExprContext) Get_listStructExp() IListStructExpContext { return s._listStructExp }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_callFunction(v ICallFunctionContext) { s._callFunction = v }

func (s *ExprContext) Set_listStructExp(v IListStructExpContext) { s._listStructExp = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) CallFunction() ICallFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ExprContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ExprContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *ExprContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *ExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *ExprContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *ExprContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserISEMPTY, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMODULO, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(474)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)

			var _x = p.expr(23)

			localctx.(*ExprContext).right = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation(localctx.(*ExprContext).GetOp().GetLine(), localctx.(*ExprContext).GetOp().GetColumn(), localctx.(*ExprContext).GetRight().GetE(), "neg", localctx.(*ExprContext).GetRight().GetE())

	case 2:
		{
			p.SetState(478)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)

			var _x = p.expr(15)

			localctx.(*ExprContext).right = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation(localctx.(*ExprContext).GetOp().GetLine(), localctx.(*ExprContext).GetOp().GetColumn(), localctx.(*ExprContext).GetRight().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()), localctx.(*ExprContext).GetRight().GetE())

	case 3:
		{
			p.SetState(482)

			var _x = p.CallFunction()

			localctx.(*ExprContext)._callFunction = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_callFunction().GetCf()

	case 4:
		{
			p.SetState(485)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)

			var _x = p.listStructExp(0)

			localctx.(*ExprContext)._listStructExp = _x
		}
		{
			p.SetState(488)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewStructExp((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()), localctx.(*ExprContext).Get_listStructExp().GetL())

	case 5:
		{
			p.SetState(491)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(493)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 6:
		{
			p.SetState(496)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case 7:
		{
			p.SetState(499)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(501)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listParams().GetL())

	case 8:
		{
			p.SetState(504)

			var _m = p.Match(SwiftGrammarParserINT)

			localctx.(*ExprContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(507)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewConversion(localctx.(*ExprContext).Get_INT().GetLine(), localctx.(*ExprContext).Get_INT().GetColumn(), localctx.(*ExprContext).Get_expr().GetE(), environment.INTEGER)

	case 9:
		{
			p.SetState(510)

			var _m = p.Match(SwiftGrammarParserSTR)

			localctx.(*ExprContext)._STR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(513)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewConversion(localctx.(*ExprContext).Get_STR().GetLine(), localctx.(*ExprContext).Get_STR().GetColumn(), localctx.(*ExprContext).Get_expr().GetE(), environment.STRING)

	case 10:
		{
			p.SetState(516)

			var _m = p.Match(SwiftGrammarParserFLOAT)

			localctx.(*ExprContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(519)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewConversion(localctx.(*ExprContext).Get_FLOAT().GetLine(), localctx.(*ExprContext).Get_FLOAT().GetColumn(), localctx.(*ExprContext).Get_expr().GetE(), environment.FLOAT)

	case 11:
		{
			p.SetState(522)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCount((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 12:
		{
			p.SetState(526)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Match(SwiftGrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewIsEmpty((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 13:
		{
			p.SetState(530)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 14:
		{
			p.SetState(532)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 15:
		{
			p.SetState(534)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 16:
		{
			p.SetState(536)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(575)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(540)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(541)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144141576354922496) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(542)

					var _x = p.expr(23)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(545)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(546)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(547)

					var _x = p.expr(22)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(550)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(551)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(552)

					var _x = p.expr(21)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(555)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(556)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(557)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(560)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(561)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(562)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(565)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(566)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(567)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(570)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(571)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(572)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *SwiftGrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *SwiftGrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(584)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(585)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(586)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext
	PUNTO() antlr.TerminalNode

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expression
	list   IListArrayContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
	return p
}

func InitEmptyListArrayContext(p *ListArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *SwiftGrammarParser) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *SwiftGrammarParser) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(595)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ListArrayContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ListArrayContext).p = expressions.NewCallVar((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(608)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(598)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(599)
					p.Match(SwiftGrammarParserCORIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(600)

					var _x = p.expr(0)

					localctx.(*ListArrayContext)._expr = _x
				}
				{
					p.SetState(601)
					p.Match(SwiftGrammarParserCORDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expr().GetE())

			case 2:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(604)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(605)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(606)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListArrayContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewStructAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), (func() string {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetText()
					}
				}()))

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(612)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFunctionInstContext is an interface to support dynamic dispatch.
type ICallFunctionInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetCfi returns the cfi attribute.
	GetCfi() interfaces.Instruction

	// SetCfi sets the cfi attribute.
	SetCfi(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	PARDER() antlr.TerminalNode

	// IsCallFunctionInstContext differentiates from other interfaces.
	IsCallFunctionInstContext()
}

type CallFunctionInstContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	cfi         interfaces.Instruction
	_ID         antlr.Token
	_listParams IListParamsContext
}

func NewEmptyCallFunctionInstContext() *CallFunctionInstContext {
	var p = new(CallFunctionInstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunctionInst
	return p
}

func InitEmptyCallFunctionInstContext(p *CallFunctionInstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunctionInst
}

func (*CallFunctionInstContext) IsCallFunctionInstContext() {}

func NewCallFunctionInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionInstContext {
	var p = new(CallFunctionInstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callFunctionInst

	return p
}

func (s *CallFunctionInstContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionInstContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFunctionInstContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFunctionInstContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *CallFunctionInstContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *CallFunctionInstContext) GetCfi() interfaces.Instruction { return s.cfi }

func (s *CallFunctionInstContext) SetCfi(v interfaces.Instruction) { s.cfi = v }

func (s *CallFunctionInstContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallFunctionInstContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallFunctionInstContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *CallFunctionInstContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallFunctionInstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionInstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunctionInstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallFunctionInst(s)
	}
}

func (s *CallFunctionInstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallFunctionInst(s)
	}
}

func (p *SwiftGrammarParser) CallFunctionInst() (localctx ICallFunctionInstContext) {
	localctx = NewCallFunctionInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_callFunctionInst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFunctionInstContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(614)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(615)

		var _x = p.listParams(0)

		localctx.(*CallFunctionInstContext)._listParams = _x
	}
	{
		p.SetState(616)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallFunctionInstContext).cfi = instructions.NewCall_Function((func() int {
		if localctx.(*CallFunctionInstContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionInstContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallFunctionInstContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionInstContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallFunctionInstContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallFunctionInstContext).Get_ID().GetText()
		}
	}()), localctx.(*CallFunctionInstContext).Get_listParams().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFunctionContext is an interface to support dynamic dispatch.
type ICallFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCf returns the cf attribute.
	GetCf() interfaces.Expression

	// SetCf sets the cf attribute.
	SetCf(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext
	PARDER() antlr.TerminalNode

	// IsCallFunctionContext differentiates from other interfaces.
	IsCallFunctionContext()
}

type CallFunctionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	cf              interfaces.Expression
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallFunctionContext() *CallFunctionContext {
	var p = new(CallFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunction
	return p
}

func InitEmptyCallFunctionContext(p *CallFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunction
}

func (*CallFunctionContext) IsCallFunctionContext() {}

func NewCallFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionContext {
	var p = new(CallFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callFunction

	return p
}

func (s *CallFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFunctionContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallFunctionContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallFunctionContext) GetCf() interfaces.Expression { return s.cf }

func (s *CallFunctionContext) SetCf(v interfaces.Expression) { s.cf = v }

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallFunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallFunctionContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallFunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallFunction(s)
	}
}

func (s *CallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallFunction(s)
	}
}

func (p *SwiftGrammarParser) CallFunction() (localctx ICallFunctionContext) {
	localctx = NewCallFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_callFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(619)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFunctionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(620)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(621)

		var _x = p.listParamsCall(0)

		localctx.(*CallFunctionContext)._listParamsCall = _x
	}
	{
		p.SetState(622)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallFunctionContext).cf = expressions.NewCallExp((func() int {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetText()
		}
	}()), localctx.(*CallFunctionContext).Get_listParamsCall().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsCallContext is an interface to support dynamic dispatch.
type IListParamsCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsCallContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsCallContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext

	// IsListParamsCallContext differentiates from other interfaces.
	IsListParamsCallContext()
}

type ListParamsCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsCallContext
	_expr  IExprContext
}

func NewEmptyListParamsCallContext() *ListParamsCallContext {
	var p = new(ListParamsCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
	return p
}

func InitEmptyListParamsCallContext(p *ListParamsCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
}

func (*ListParamsCallContext) IsListParamsCallContext() {}

func NewListParamsCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsCallContext {
	var p = new(ListParamsCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall

	return p
}

func (s *ListParamsCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsCallContext) GetList() IListParamsCallContext { return s.list }

func (s *ListParamsCallContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsCallContext) SetList(v IListParamsCallContext) { s.list = v }

func (s *ListParamsCallContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsCallContext) GetL() []interface{} { return s.l }

func (s *ListParamsCallContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsCallContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsCallContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *ListParamsCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsCall(s)
	}
}

func (s *ListParamsCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsCall(s)
	}
}

func (p *SwiftGrammarParser) ListParamsCall() (localctx IListParamsCallContext) {
	return p.listParamsCall(0)
}

func (p *SwiftGrammarParser) listParamsCall(_p int) (localctx IListParamsCallContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsCallContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsCallContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, SwiftGrammarParserRULE_listParamsCall, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(630)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(626)

			var _x = p.expr(0)

			localctx.(*ListParamsCallContext)._expr = _x
		}

		localctx.(*ListParamsCallContext).l = []interface{}{}
		localctx.(*ListParamsCallContext).l = append(localctx.(*ListParamsCallContext).l, localctx.(*ListParamsCallContext).Get_expr().GetE())

	case 2:

		localctx.(*ListParamsCallContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsCallContext(p, _parentctx, _parentState)
			localctx.(*ListParamsCallContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsCall)
			p.SetState(632)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(633)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(634)

				var _x = p.expr(0)

				localctx.(*ListParamsCallContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsCallContext).GetList().GetL(), localctx.(*ListParamsCallContext).Get_expr().GetE())
			localctx.(*ListParamsCallContext).l = arr

		}
		p.SetState(641)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructExpContext is an interface to support dynamic dispatch.
type IListStructExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructExpContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListStructExpContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListStructExp() IListStructExpContext

	// IsListStructExpContext differentiates from other interfaces.
	IsListStructExpContext()
}

type ListStructExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructExpContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListStructExpContext() *ListStructExpContext {
	var p = new(ListStructExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
	return p
}

func InitEmptyListStructExpContext(p *ListStructExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
}

func (*ListStructExpContext) IsListStructExpContext() {}

func NewListStructExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructExpContext {
	var p = new(ListStructExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp

	return p
}

func (s *ListStructExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructExpContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructExpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructExpContext) GetList() IListStructExpContext { return s.list }

func (s *ListStructExpContext) Get_expr() IExprContext { return s._expr }

func (s *ListStructExpContext) SetList(v IListStructExpContext) { s.list = v }

func (s *ListStructExpContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListStructExpContext) GetL() []interface{} { return s.l }

func (s *ListStructExpContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructExpContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListStructExpContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *ListStructExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListStructExpContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructExpContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ListStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructExp(s)
	}
}

func (s *ListStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructExp(s)
	}
}

func (p *SwiftGrammarParser) ListStructExp() (localctx IListStructExpContext) {
	return p.listStructExp(0)
}

func (p *SwiftGrammarParser) listStructExp(_p int) (localctx IListStructExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, SwiftGrammarParserRULE_listStructExp, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(643)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructExpContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(644)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(645)

			var _x = p.expr(0)

			localctx.(*ListStructExpContext)._expr = _x
		}

		var arr []interface{}
		StrExp := environment.NewStructContent((func() string {
			if localctx.(*ListStructExpContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructExpContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructExpContext).Get_expr().GetE())
		arr = append(arr, StrExp)
		localctx.(*ListStructExpContext).l = arr

	case 2:

		localctx.(*ListStructExpContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(660)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructExpContext(p, _parentctx, _parentState)
			localctx.(*ListStructExpContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructExp)
			p.SetState(651)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(652)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(653)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructExpContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(654)
				p.Match(SwiftGrammarParserD_PTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(655)

				var _x = p.expr(0)

				localctx.(*ListStructExpContext)._expr = _x
			}

			var arr []interface{}
			StrExp := environment.NewStructContent((func() string {
				if localctx.(*ListStructExpContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructExpContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructExpContext).Get_expr().GetE())
			arr = append(localctx.(*ListStructExpContext).GetList().GetL(), StrExp)
			localctx.(*ListStructExpContext).l = arr

		}
		p.SetState(662)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 7:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 17:
		var t *CaselistContext = nil
		if localctx != nil {
			t = localctx.(*CaselistContext)
		}
		return p.Caselist_Sempred(t, predIndex)

	case 20:
		var t *ElseiflistContext = nil
		if localctx != nil {
			t = localctx.(*ElseiflistContext)
		}
		return p.Elseiflist_Sempred(t, predIndex)

	case 26:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 27:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 28:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 31:
		var t *ListParamsCallContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsCallContext)
		}
		return p.ListParamsCall_Sempred(t, predIndex)

	case 32:
		var t *ListStructExpContext = nil
		if localctx != nil {
			t = localctx.(*ListStructExpContext)
		}
		return p.ListStructExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsFunc_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Caselist_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Elseiflist_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsCall_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListStructExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
