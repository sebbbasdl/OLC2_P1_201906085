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
		"'in'", "'var'", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'",
		"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('",
		"')'", "'{'", "'}'", "':'", "'['", "']'", "','", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "NUMBER",
		"STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG",
		"MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "forstmt", "switchtmt", "caselist",
		"case", "ifstmt", "elseiflist", "elseif", "whilestmt", "declarationstmt",
		"asignacion", "types", "expr", "listParams", "listArray",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 317, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 42, 8,
		1, 11, 1, 12, 1, 43, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 69, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 99, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 119, 8, 6, 10, 6, 12, 6, 122, 9,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 168, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 178, 8, 9, 10, 9, 12,
		9, 181, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		222, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 246, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15,
		283, 8, 15, 10, 15, 12, 15, 286, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 297, 8, 16, 10, 16, 12, 16, 300,
		9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 5, 17, 312, 8, 17, 10, 17, 12, 17, 315, 9, 17, 1, 17, 0, 5, 12, 18,
		30, 32, 34, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 0, 6, 1, 0, 18, 19, 1, 0, 30, 31, 1, 0, 32, 33, 2, 0, 26, 26,
		28, 28, 2, 0, 27, 27, 29, 29, 1, 0, 20, 21, 330, 0, 36, 1, 0, 0, 0, 2,
		41, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8, 98, 1, 0, 0, 0,
		10, 100, 1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 167,
		1, 0, 0, 0, 18, 169, 1, 0, 0, 0, 20, 182, 1, 0, 0, 0, 22, 190, 1, 0, 0,
		0, 24, 197, 1, 0, 0, 0, 26, 205, 1, 0, 0, 0, 28, 221, 1, 0, 0, 0, 30, 245,
		1, 0, 0, 0, 32, 287, 1, 0, 0, 0, 34, 301, 1, 0, 0, 0, 36, 37, 3, 2, 1,
		0, 37, 38, 5, 0, 0, 1, 38, 39, 6, 0, -1, 0, 39, 1, 1, 0, 0, 0, 40, 42,
		3, 4, 2, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 6, 1, -1, 0, 46, 3, 1,
		0, 0, 0, 47, 48, 3, 6, 3, 0, 48, 49, 6, 2, -1, 0, 49, 69, 1, 0, 0, 0, 50,
		51, 3, 16, 8, 0, 51, 52, 6, 2, -1, 0, 52, 69, 1, 0, 0, 0, 53, 54, 3, 24,
		12, 0, 54, 55, 6, 2, -1, 0, 55, 69, 1, 0, 0, 0, 56, 57, 3, 26, 13, 0, 57,
		58, 6, 2, -1, 0, 58, 69, 1, 0, 0, 0, 59, 60, 3, 22, 11, 0, 60, 61, 6, 2,
		-1, 0, 61, 69, 1, 0, 0, 0, 62, 63, 3, 8, 4, 0, 63, 64, 6, 2, -1, 0, 64,
		69, 1, 0, 0, 0, 65, 66, 3, 10, 5, 0, 66, 67, 6, 2, -1, 0, 67, 69, 1, 0,
		0, 0, 68, 47, 1, 0, 0, 0, 68, 50, 1, 0, 0, 0, 68, 53, 1, 0, 0, 0, 68, 56,
		1, 0, 0, 0, 68, 59, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0,
		69, 5, 1, 0, 0, 0, 70, 71, 5, 7, 0, 0, 71, 72, 5, 34, 0, 0, 72, 73, 3,
		30, 15, 0, 73, 74, 5, 35, 0, 0, 74, 75, 6, 3, -1, 0, 75, 7, 1, 0, 0, 0,
		76, 77, 5, 11, 0, 0, 77, 78, 5, 19, 0, 0, 78, 79, 5, 15, 0, 0, 79, 80,
		3, 30, 15, 0, 80, 81, 5, 42, 0, 0, 81, 82, 5, 42, 0, 0, 82, 83, 5, 42,
		0, 0, 83, 84, 3, 30, 15, 0, 84, 85, 5, 36, 0, 0, 85, 86, 3, 2, 1, 0, 86,
		87, 5, 37, 0, 0, 87, 88, 6, 4, -1, 0, 88, 99, 1, 0, 0, 0, 89, 90, 5, 11,
		0, 0, 90, 91, 5, 19, 0, 0, 91, 92, 5, 15, 0, 0, 92, 93, 7, 0, 0, 0, 93,
		94, 5, 36, 0, 0, 94, 95, 3, 2, 1, 0, 95, 96, 5, 37, 0, 0, 96, 97, 6, 4,
		-1, 0, 97, 99, 1, 0, 0, 0, 98, 76, 1, 0, 0, 0, 98, 89, 1, 0, 0, 0, 99,
		9, 1, 0, 0, 0, 100, 101, 5, 12, 0, 0, 101, 102, 3, 30, 15, 0, 102, 103,
		5, 36, 0, 0, 103, 104, 3, 12, 6, 0, 104, 105, 5, 14, 0, 0, 105, 106, 5,
		38, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108, 5, 37, 0, 0, 108, 109, 6, 5,
		-1, 0, 109, 11, 1, 0, 0, 0, 110, 111, 6, 6, -1, 0, 111, 112, 3, 14, 7,
		0, 112, 113, 6, 6, -1, 0, 113, 120, 1, 0, 0, 0, 114, 115, 10, 2, 0, 0,
		115, 116, 3, 14, 7, 0, 116, 117, 6, 6, -1, 0, 117, 119, 1, 0, 0, 0, 118,
		114, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 13, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 5, 13,
		0, 0, 124, 125, 3, 30, 15, 0, 125, 126, 5, 38, 0, 0, 126, 127, 3, 2, 1,
		0, 127, 128, 6, 7, -1, 0, 128, 15, 1, 0, 0, 0, 129, 130, 5, 8, 0, 0, 130,
		131, 3, 30, 15, 0, 131, 132, 5, 36, 0, 0, 132, 133, 3, 2, 1, 0, 133, 134,
		5, 37, 0, 0, 134, 135, 6, 8, -1, 0, 135, 168, 1, 0, 0, 0, 136, 137, 5,
		8, 0, 0, 137, 138, 3, 30, 15, 0, 138, 139, 5, 36, 0, 0, 139, 140, 3, 2,
		1, 0, 140, 141, 5, 37, 0, 0, 141, 142, 5, 9, 0, 0, 142, 143, 5, 36, 0,
		0, 143, 144, 3, 2, 1, 0, 144, 145, 5, 37, 0, 0, 145, 146, 6, 8, -1, 0,
		146, 168, 1, 0, 0, 0, 147, 148, 5, 8, 0, 0, 148, 149, 3, 30, 15, 0, 149,
		150, 5, 36, 0, 0, 150, 151, 3, 2, 1, 0, 151, 152, 5, 37, 0, 0, 152, 153,
		3, 18, 9, 0, 153, 154, 6, 8, -1, 0, 154, 168, 1, 0, 0, 0, 155, 156, 5,
		8, 0, 0, 156, 157, 3, 30, 15, 0, 157, 158, 5, 36, 0, 0, 158, 159, 3, 2,
		1, 0, 159, 160, 5, 37, 0, 0, 160, 161, 3, 18, 9, 0, 161, 162, 5, 9, 0,
		0, 162, 163, 5, 36, 0, 0, 163, 164, 3, 2, 1, 0, 164, 165, 5, 37, 0, 0,
		165, 166, 6, 8, -1, 0, 166, 168, 1, 0, 0, 0, 167, 129, 1, 0, 0, 0, 167,
		136, 1, 0, 0, 0, 167, 147, 1, 0, 0, 0, 167, 155, 1, 0, 0, 0, 168, 17, 1,
		0, 0, 0, 169, 170, 6, 9, -1, 0, 170, 171, 3, 20, 10, 0, 171, 172, 6, 9,
		-1, 0, 172, 179, 1, 0, 0, 0, 173, 174, 10, 2, 0, 0, 174, 175, 3, 20, 10,
		0, 175, 176, 6, 9, -1, 0, 176, 178, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 178,
		181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 19, 1,
		0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 183, 5, 9, 0, 0, 183, 184, 5, 8, 0,
		0, 184, 185, 3, 30, 15, 0, 185, 186, 5, 36, 0, 0, 186, 187, 3, 2, 1, 0,
		187, 188, 5, 37, 0, 0, 188, 189, 6, 10, -1, 0, 189, 21, 1, 0, 0, 0, 190,
		191, 5, 10, 0, 0, 191, 192, 3, 30, 15, 0, 192, 193, 5, 36, 0, 0, 193, 194,
		3, 2, 1, 0, 194, 195, 5, 37, 0, 0, 195, 196, 6, 11, -1, 0, 196, 23, 1,
		0, 0, 0, 197, 198, 5, 16, 0, 0, 198, 199, 5, 19, 0, 0, 199, 200, 5, 38,
		0, 0, 200, 201, 3, 28, 14, 0, 201, 202, 5, 25, 0, 0, 202, 203, 3, 30, 15,
		0, 203, 204, 6, 12, -1, 0, 204, 25, 1, 0, 0, 0, 205, 206, 5, 19, 0, 0,
		206, 207, 5, 25, 0, 0, 207, 208, 3, 30, 15, 0, 208, 209, 6, 13, -1, 0,
		209, 27, 1, 0, 0, 0, 210, 211, 5, 1, 0, 0, 211, 222, 6, 14, -1, 0, 212,
		213, 5, 2, 0, 0, 213, 222, 6, 14, -1, 0, 214, 215, 5, 4, 0, 0, 215, 222,
		6, 14, -1, 0, 216, 217, 5, 3, 0, 0, 217, 222, 6, 14, -1, 0, 218, 219, 5,
		39, 0, 0, 219, 220, 5, 40, 0, 0, 220, 222, 6, 14, -1, 0, 221, 210, 1, 0,
		0, 0, 221, 212, 1, 0, 0, 0, 221, 214, 1, 0, 0, 0, 221, 216, 1, 0, 0, 0,
		221, 218, 1, 0, 0, 0, 222, 29, 1, 0, 0, 0, 223, 224, 6, 15, -1, 0, 224,
		225, 5, 34, 0, 0, 225, 226, 3, 30, 15, 0, 226, 227, 5, 35, 0, 0, 227, 228,
		6, 15, -1, 0, 228, 246, 1, 0, 0, 0, 229, 230, 3, 34, 17, 0, 230, 231, 6,
		15, -1, 0, 231, 246, 1, 0, 0, 0, 232, 233, 5, 39, 0, 0, 233, 234, 3, 32,
		16, 0, 234, 235, 5, 40, 0, 0, 235, 236, 6, 15, -1, 0, 236, 246, 1, 0, 0,
		0, 237, 238, 5, 17, 0, 0, 238, 246, 6, 15, -1, 0, 239, 240, 5, 18, 0, 0,
		240, 246, 6, 15, -1, 0, 241, 242, 5, 5, 0, 0, 242, 246, 6, 15, -1, 0, 243,
		244, 5, 6, 0, 0, 244, 246, 6, 15, -1, 0, 245, 223, 1, 0, 0, 0, 245, 229,
		1, 0, 0, 0, 245, 232, 1, 0, 0, 0, 245, 237, 1, 0, 0, 0, 245, 239, 1, 0,
		0, 0, 245, 241, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 284, 1, 0, 0, 0,
		247, 248, 10, 14, 0, 0, 248, 249, 7, 1, 0, 0, 249, 250, 3, 30, 15, 15,
		250, 251, 6, 15, -1, 0, 251, 283, 1, 0, 0, 0, 252, 253, 10, 13, 0, 0, 253,
		254, 7, 2, 0, 0, 254, 255, 3, 30, 15, 14, 255, 256, 6, 15, -1, 0, 256,
		283, 1, 0, 0, 0, 257, 258, 10, 12, 0, 0, 258, 259, 7, 3, 0, 0, 259, 260,
		3, 30, 15, 13, 260, 261, 6, 15, -1, 0, 261, 283, 1, 0, 0, 0, 262, 263,
		10, 11, 0, 0, 263, 264, 7, 4, 0, 0, 264, 265, 3, 30, 15, 12, 265, 266,
		6, 15, -1, 0, 266, 283, 1, 0, 0, 0, 267, 268, 10, 10, 0, 0, 268, 269, 7,
		5, 0, 0, 269, 270, 3, 30, 15, 11, 270, 271, 6, 15, -1, 0, 271, 283, 1,
		0, 0, 0, 272, 273, 10, 9, 0, 0, 273, 274, 5, 24, 0, 0, 274, 275, 3, 30,
		15, 10, 275, 276, 6, 15, -1, 0, 276, 283, 1, 0, 0, 0, 277, 278, 10, 8,
		0, 0, 278, 279, 5, 23, 0, 0, 279, 280, 3, 30, 15, 9, 280, 281, 6, 15, -1,
		0, 281, 283, 1, 0, 0, 0, 282, 247, 1, 0, 0, 0, 282, 252, 1, 0, 0, 0, 282,
		257, 1, 0, 0, 0, 282, 262, 1, 0, 0, 0, 282, 267, 1, 0, 0, 0, 282, 272,
		1, 0, 0, 0, 282, 277, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 284, 285, 1, 0, 0, 0, 285, 31, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0,
		287, 288, 6, 16, -1, 0, 288, 289, 3, 30, 15, 0, 289, 290, 6, 16, -1, 0,
		290, 298, 1, 0, 0, 0, 291, 292, 10, 2, 0, 0, 292, 293, 5, 41, 0, 0, 293,
		294, 3, 30, 15, 0, 294, 295, 6, 16, -1, 0, 295, 297, 1, 0, 0, 0, 296, 291,
		1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0,
		0, 0, 299, 33, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302, 6, 17, -1, 0,
		302, 303, 5, 19, 0, 0, 303, 304, 6, 17, -1, 0, 304, 313, 1, 0, 0, 0, 305,
		306, 10, 2, 0, 0, 306, 307, 5, 39, 0, 0, 307, 308, 3, 30, 15, 0, 308, 309,
		5, 40, 0, 0, 309, 310, 6, 17, -1, 0, 310, 312, 1, 0, 0, 0, 311, 305, 1,
		0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0,
		0, 314, 35, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 12, 43, 68, 98, 120, 167,
		179, 221, 245, 282, 284, 298, 313,
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
	SwiftGrammarParserNUMBER       = 17
	SwiftGrammarParserSTRING       = 18
	SwiftGrammarParserID           = 19
	SwiftGrammarParserDIF          = 20
	SwiftGrammarParserIG_IG        = 21
	SwiftGrammarParserNOT          = 22
	SwiftGrammarParserOR           = 23
	SwiftGrammarParserAND          = 24
	SwiftGrammarParserIG           = 25
	SwiftGrammarParserMAY_IG       = 26
	SwiftGrammarParserMEN_IG       = 27
	SwiftGrammarParserMAYOR        = 28
	SwiftGrammarParserMENOR        = 29
	SwiftGrammarParserMUL          = 30
	SwiftGrammarParserDIV          = 31
	SwiftGrammarParserADD          = 32
	SwiftGrammarParserSUB          = 33
	SwiftGrammarParserPARIZQ       = 34
	SwiftGrammarParserPARDER       = 35
	SwiftGrammarParserLLAVEIZQ     = 36
	SwiftGrammarParserLLAVEDER     = 37
	SwiftGrammarParserD_PTS        = 38
	SwiftGrammarParserCORIZQ       = 39
	SwiftGrammarParserCORDER       = 40
	SwiftGrammarParserCOMA         = 41
	SwiftGrammarParserPUNTO        = 42
	SwiftGrammarParserWHITESPACE   = 43
	SwiftGrammarParserCOMMENT      = 44
	SwiftGrammarParserLINE_COMMENT = 45
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s               = 0
	SwiftGrammarParserRULE_block           = 1
	SwiftGrammarParserRULE_instruction     = 2
	SwiftGrammarParserRULE_printstmt       = 3
	SwiftGrammarParserRULE_forstmt         = 4
	SwiftGrammarParserRULE_switchtmt       = 5
	SwiftGrammarParserRULE_caselist        = 6
	SwiftGrammarParserRULE_case            = 7
	SwiftGrammarParserRULE_ifstmt          = 8
	SwiftGrammarParserRULE_elseiflist      = 9
	SwiftGrammarParserRULE_elseif          = 10
	SwiftGrammarParserRULE_whilestmt       = 11
	SwiftGrammarParserRULE_declarationstmt = 12
	SwiftGrammarParserRULE_asignacion      = 13
	SwiftGrammarParserRULE_types           = 14
	SwiftGrammarParserRULE_expr            = 15
	SwiftGrammarParserRULE_listParams      = 16
	SwiftGrammarParserRULE_listArray       = 17
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
		p.SetState(36)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(37)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(40)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(43)
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

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	inst             interfaces.Instruction
	_printstmt       IPrintstmtContext
	_ifstmt          IIfstmtContext
	_declarationstmt IDeclarationstmtContext
	_asignacion      IAsignacionContext
	_whilestmt       IWhilestmtContext
	_forstmt         IForstmtContext
	_switchtmt       ISwitchtmtContext
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

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_declarationstmt(v IDeclarationstmtContext) { s._declarationstmt = v }

func (s *InstructionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_switchtmt(v ISwitchtmtContext) { s._switchtmt = v }

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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case SwiftGrammarParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case SwiftGrammarParserVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case SwiftGrammarParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)

			var _x = p.Asignacion()

			localctx.(*InstructionContext)._asignacion = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignacion().GetAsign()

	case SwiftGrammarParserWHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhiles()

	case SwiftGrammarParserFOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(62)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFors()

	case SwiftGrammarParserSWITCH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)

			var _x = p.Switchtmt()

			localctx.(*InstructionContext)._switchtmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchtmt().GetSwtch()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prnt   interfaces.Instruction
	_PRINT antlr.Token
	_expr  IExprContext
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

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
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

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
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
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(73)
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
	}()), localctx.(*PrintstmtContext).Get_expr().GetE())

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

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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
	STRING() antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	fors   interfaces.Instruction
	_FOR   antlr.Token
	_ID    antlr.Token
	e1     IExprContext
	e2     IExprContext
	_block IBlockContext
	ope    antlr.Token
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

func (s *ForstmtContext) GetOpe() antlr.Token { return s.ope }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

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
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_forstmt)
	var _la int

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e1 = _x
		}
		{
			p.SetState(80)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e2 = _x
		}
		{
			p.SetState(84)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(86)
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
			p.SetState(89)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)

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
			p.SetState(93)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(95)
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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_switchtmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)

		var _m = p.Match(SwiftGrammarParserSWITCH)

		localctx.(*SwitchtmtContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)

		var _x = p.expr(0)

		localctx.(*SwitchtmtContext)._expr = _x
	}
	{
		p.SetState(102)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)

		var _x = p.caselist(0)

		localctx.(*SwitchtmtContext)._caselist = _x
	}
	{
		p.SetState(104)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _x = p.Block()

		localctx.(*SwitchtmtContext)._block = _x
	}
	{
		p.SetState(107)
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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, SwiftGrammarParserRULE_caselist, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)

		var _x = p.Case_()

		localctx.(*CaselistContext)._case = _x
	}

	localctx.(*CaselistContext).cases = []interface{}{}
	localctx.(*CaselistContext).cases = append(localctx.(*CaselistContext).cases, localctx.(*CaselistContext).Get_case().GetSwtch())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
			p.SetState(114)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(115)

				var _x = p.Case_()

				localctx.(*CaselistContext)._case = _x
			}
			var arr3 []interface{}
			arr3 = append(localctx.(*CaselistContext).GetListcases().GetCases(), localctx.(*CaselistContext).Get_case().GetSwtch())
			localctx.(*CaselistContext).cases = arr3

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CaseContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)

		var _x = p.expr(0)

		localctx.(*CaseContext)._expr = _x
	}
	{
		p.SetState(125)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)

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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_ifstmt)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(131)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(133)
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
			p.SetState(136)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(138)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)

			var _x = p.Block()

			localctx.(*IfstmtContext).b1 = _x
		}
		{
			p.SetState(140)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)

			var _x = p.Block()

			localctx.(*IfstmtContext).b2 = _x
		}
		{
			p.SetState(144)
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
			p.SetState(147)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(149)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(151)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

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
			p.SetState(155)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(157)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)

			var _x = p.Block()

			localctx.(*IfstmtContext).b1 = _x
		}
		{
			p.SetState(159)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _x = p.elseiflist(0)

			localctx.(*IfstmtContext)._elseiflist = _x
		}
		{
			p.SetState(161)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)

			var _x = p.Block()

			localctx.(*IfstmtContext).b2 = _x
		}
		{
			p.SetState(164)
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
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SwiftGrammarParserRULE_elseiflist, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)

		var _x = p.Elseif()

		localctx.(*ElseiflistContext)._elseif = _x
	}

	localctx.(*ElseiflistContext).else_if_list = []interface{}{}
	localctx.(*ElseiflistContext).else_if_list = append(localctx.(*ElseiflistContext).else_if_list, localctx.(*ElseiflistContext).Get_elseif().GetIfinst())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
			p.SetState(173)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(174)

				var _x = p.Elseif()

				localctx.(*ElseiflistContext)._elseif = _x
			}
			var arr2 []interface{}
			arr2 = append(localctx.(*ElseiflistContext).GetListif().GetElse_if_list(), localctx.(*ElseiflistContext).Get_elseif().GetIfinst())
			localctx.(*ElseiflistContext).else_if_list = arr2

		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_elseif)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)

		var _m = p.Match(SwiftGrammarParserIF)

		localctx.(*ElseifContext)._IF = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)

		var _x = p.expr(0)

		localctx.(*ElseifContext)._expr = _x
	}
	{
		p.SetState(185)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)

		var _x = p.Block()

		localctx.(*ElseifContext)._block = _x
	}
	{
		p.SetState(187)
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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(192)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(194)
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

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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

func (s *DeclarationstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

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
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_declarationstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)

		var _m = p.Match(SwiftGrammarParserVAR)

		localctx.(*DeclarationstmtContext)._VAR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*DeclarationstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)

		var _x = p.Types()

		localctx.(*DeclarationstmtContext)._types = _x
	}
	{
		p.SetState(201)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)

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
	}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE()))

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

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	asign  interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
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

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_asignacion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AsignacionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)

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
	}()), localctx.(*AsignacionContext).Get_expr().GetE())

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
	CORDER() antlr.TerminalNode

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

func (s *TypesContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
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
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_types)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(214)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(216)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	PARIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	ListArray() IListArrayContext
	CORIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	CORDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
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
	parser      antlr.Parser
	e           interfaces.Expression
	left        IExprContext
	_expr       IExprContext
	list        IListArrayContext
	_CORIZQ     antlr.Token
	_listParams IListParamsContext
	_NUMBER     antlr.Token
	_STRING     antlr.Token
	_TRU        antlr.Token
	_FAL        antlr.Token
	op          antlr.Token
	right       IExprContext
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

func (s *ExprContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
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

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPARIZQ:
		{
			p.SetState(224)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(226)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case SwiftGrammarParserID:
		{
			p.SetState(229)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case SwiftGrammarParserCORIZQ:
		{
			p.SetState(232)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(234)
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

	case SwiftGrammarParserNUMBER:
		{
			p.SetState(237)

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

	case SwiftGrammarParserSTRING:
		{
			p.SetState(239)

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

	case SwiftGrammarParserTRU:
		{
			p.SetState(241)

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

	case SwiftGrammarParserFAL:
		{
			p.SetState(243)

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

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(284)
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
			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(248)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMUL || _la == SwiftGrammarParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(249)

					var _x = p.expr(15)

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
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(253)

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
					p.SetState(254)

					var _x = p.expr(14)

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
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(258)

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
					p.SetState(259)

					var _x = p.expr(13)

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
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(263)

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
					p.SetState(264)

					var _x = p.expr(12)

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
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(268)

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
					p.SetState(269)

					var _x = p.expr(11)

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
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(273)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(274)

					var _x = p.expr(10)

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
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(278)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(279)

					var _x = p.expr(9)

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
		p.SetState(286)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
			p.SetState(291)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(292)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(293)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)

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
	p.SetState(313)
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
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
			p.SetState(305)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(306)
				p.Match(SwiftGrammarParserCORIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(307)

				var _x = p.expr(0)

				localctx.(*ListArrayContext)._expr = _x
			}
			{
				p.SetState(308)
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

		}
		p.SetState(315)
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

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *CaselistContext = nil
		if localctx != nil {
			t = localctx.(*CaselistContext)
		}
		return p.Caselist_Sempred(t, predIndex)

	case 9:
		var t *ElseiflistContext = nil
		if localctx != nil {
			t = localctx.(*ElseiflistContext)
		}
		return p.Elseiflist_Sempred(t, predIndex)

	case 15:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 16:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 17:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Caselist_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Elseiflist_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
