// Code generated from SwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlexerLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", "'print'",
		"'if'", "'else'", "'while'", "'for'", "'in'", "'var'", "", "", "", "'!='",
		"'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'",
		"'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "':'", "'['",
		"']'", "','", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "IN", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS",
		"CORIZQ", "CORDER", "COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "IN", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS",
		"CORIZQ", "CORDER", "COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 267, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 153,
		8, 13, 11, 13, 12, 13, 154, 1, 13, 1, 13, 4, 13, 159, 8, 13, 11, 13, 12,
		13, 160, 3, 13, 163, 8, 13, 1, 14, 1, 14, 5, 14, 167, 8, 14, 10, 14, 12,
		14, 170, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 176, 8, 15, 10, 15,
		12, 15, 179, 9, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 39, 4, 39, 234, 8, 39, 11, 39, 12, 39, 235, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 244, 8, 40, 10, 40, 12, 40, 247,
		9, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5,
		41, 258, 8, 41, 10, 41, 12, 41, 261, 9, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 245, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 0, 1, 0, 7,
		1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13,
		13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 273,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 87,
		1, 0, 0, 0, 3, 91, 1, 0, 0, 0, 5, 97, 1, 0, 0, 0, 7, 102, 1, 0, 0, 0, 9,
		109, 1, 0, 0, 0, 11, 114, 1, 0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 126, 1,
		0, 0, 0, 17, 129, 1, 0, 0, 0, 19, 134, 1, 0, 0, 0, 21, 140, 1, 0, 0, 0,
		23, 144, 1, 0, 0, 0, 25, 147, 1, 0, 0, 0, 27, 152, 1, 0, 0, 0, 29, 164,
		1, 0, 0, 0, 31, 173, 1, 0, 0, 0, 33, 180, 1, 0, 0, 0, 35, 183, 1, 0, 0,
		0, 37, 186, 1, 0, 0, 0, 39, 188, 1, 0, 0, 0, 41, 191, 1, 0, 0, 0, 43, 194,
		1, 0, 0, 0, 45, 196, 1, 0, 0, 0, 47, 199, 1, 0, 0, 0, 49, 202, 1, 0, 0,
		0, 51, 204, 1, 0, 0, 0, 53, 206, 1, 0, 0, 0, 55, 208, 1, 0, 0, 0, 57, 210,
		1, 0, 0, 0, 59, 212, 1, 0, 0, 0, 61, 214, 1, 0, 0, 0, 63, 216, 1, 0, 0,
		0, 65, 218, 1, 0, 0, 0, 67, 220, 1, 0, 0, 0, 69, 222, 1, 0, 0, 0, 71, 224,
		1, 0, 0, 0, 73, 226, 1, 0, 0, 0, 75, 228, 1, 0, 0, 0, 77, 230, 1, 0, 0,
		0, 79, 233, 1, 0, 0, 0, 81, 239, 1, 0, 0, 0, 83, 253, 1, 0, 0, 0, 85, 264,
		1, 0, 0, 0, 87, 88, 5, 73, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 116,
		0, 0, 90, 2, 1, 0, 0, 0, 91, 92, 5, 70, 0, 0, 92, 93, 5, 108, 0, 0, 93,
		94, 5, 111, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 116, 0, 0, 96, 4, 1,
		0, 0, 0, 97, 98, 5, 66, 0, 0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 111, 0,
		0, 100, 101, 5, 108, 0, 0, 101, 6, 1, 0, 0, 0, 102, 103, 5, 83, 0, 0, 103,
		104, 5, 116, 0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 105, 0, 0, 106,
		107, 5, 110, 0, 0, 107, 108, 5, 103, 0, 0, 108, 8, 1, 0, 0, 0, 109, 110,
		5, 116, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 117, 0, 0, 112, 113,
		5, 101, 0, 0, 113, 10, 1, 0, 0, 0, 114, 115, 5, 102, 0, 0, 115, 116, 5,
		97, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5, 115, 0, 0, 118, 119, 5,
		101, 0, 0, 119, 12, 1, 0, 0, 0, 120, 121, 5, 112, 0, 0, 121, 122, 5, 114,
		0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125, 5, 116,
		0, 0, 125, 14, 1, 0, 0, 0, 126, 127, 5, 105, 0, 0, 127, 128, 5, 102, 0,
		0, 128, 16, 1, 0, 0, 0, 129, 130, 5, 101, 0, 0, 130, 131, 5, 108, 0, 0,
		131, 132, 5, 115, 0, 0, 132, 133, 5, 101, 0, 0, 133, 18, 1, 0, 0, 0, 134,
		135, 5, 119, 0, 0, 135, 136, 5, 104, 0, 0, 136, 137, 5, 105, 0, 0, 137,
		138, 5, 108, 0, 0, 138, 139, 5, 101, 0, 0, 139, 20, 1, 0, 0, 0, 140, 141,
		5, 102, 0, 0, 141, 142, 5, 111, 0, 0, 142, 143, 5, 114, 0, 0, 143, 22,
		1, 0, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 110, 0, 0, 146, 24, 1,
		0, 0, 0, 147, 148, 5, 118, 0, 0, 148, 149, 5, 97, 0, 0, 149, 150, 5, 114,
		0, 0, 150, 26, 1, 0, 0, 0, 151, 153, 7, 0, 0, 0, 152, 151, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		162, 1, 0, 0, 0, 156, 158, 5, 46, 0, 0, 157, 159, 7, 0, 0, 0, 158, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0,
		0, 0, 161, 163, 1, 0, 0, 0, 162, 156, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 28, 1, 0, 0, 0, 164, 168, 5, 34, 0, 0, 165, 167, 8, 1, 0, 0, 166,
		165, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169,
		1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 172, 5, 34,
		0, 0, 172, 30, 1, 0, 0, 0, 173, 177, 7, 2, 0, 0, 174, 176, 7, 3, 0, 0,
		175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 32, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 181, 5,
		33, 0, 0, 181, 182, 5, 61, 0, 0, 182, 34, 1, 0, 0, 0, 183, 184, 5, 61,
		0, 0, 184, 185, 5, 61, 0, 0, 185, 36, 1, 0, 0, 0, 186, 187, 5, 33, 0, 0,
		187, 38, 1, 0, 0, 0, 188, 189, 5, 124, 0, 0, 189, 190, 5, 124, 0, 0, 190,
		40, 1, 0, 0, 0, 191, 192, 5, 38, 0, 0, 192, 193, 5, 38, 0, 0, 193, 42,
		1, 0, 0, 0, 194, 195, 5, 61, 0, 0, 195, 44, 1, 0, 0, 0, 196, 197, 5, 62,
		0, 0, 197, 198, 5, 61, 0, 0, 198, 46, 1, 0, 0, 0, 199, 200, 5, 60, 0, 0,
		200, 201, 5, 61, 0, 0, 201, 48, 1, 0, 0, 0, 202, 203, 5, 62, 0, 0, 203,
		50, 1, 0, 0, 0, 204, 205, 5, 60, 0, 0, 205, 52, 1, 0, 0, 0, 206, 207, 5,
		42, 0, 0, 207, 54, 1, 0, 0, 0, 208, 209, 5, 47, 0, 0, 209, 56, 1, 0, 0,
		0, 210, 211, 5, 43, 0, 0, 211, 58, 1, 0, 0, 0, 212, 213, 5, 45, 0, 0, 213,
		60, 1, 0, 0, 0, 214, 215, 5, 40, 0, 0, 215, 62, 1, 0, 0, 0, 216, 217, 5,
		41, 0, 0, 217, 64, 1, 0, 0, 0, 218, 219, 5, 123, 0, 0, 219, 66, 1, 0, 0,
		0, 220, 221, 5, 125, 0, 0, 221, 68, 1, 0, 0, 0, 222, 223, 5, 58, 0, 0,
		223, 70, 1, 0, 0, 0, 224, 225, 5, 91, 0, 0, 225, 72, 1, 0, 0, 0, 226, 227,
		5, 93, 0, 0, 227, 74, 1, 0, 0, 0, 228, 229, 5, 44, 0, 0, 229, 76, 1, 0,
		0, 0, 230, 231, 5, 46, 0, 0, 231, 78, 1, 0, 0, 0, 232, 234, 7, 4, 0, 0,
		233, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235,
		236, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 6, 39, 0, 0, 238, 80,
		1, 0, 0, 0, 239, 240, 5, 47, 0, 0, 240, 241, 5, 42, 0, 0, 241, 245, 1,
		0, 0, 0, 242, 244, 9, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 247, 1, 0, 0,
		0, 245, 246, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 248, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 248, 249, 5, 42, 0, 0, 249, 250, 5, 47, 0, 0, 250, 251,
		1, 0, 0, 0, 251, 252, 6, 40, 0, 0, 252, 82, 1, 0, 0, 0, 253, 254, 5, 47,
		0, 0, 254, 255, 5, 47, 0, 0, 255, 259, 1, 0, 0, 0, 256, 258, 8, 5, 0, 0,
		257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263,
		6, 41, 0, 0, 263, 84, 1, 0, 0, 0, 264, 265, 5, 92, 0, 0, 265, 266, 7, 6,
		0, 0, 266, 86, 1, 0, 0, 0, 9, 0, 154, 160, 162, 168, 177, 235, 245, 259,
		1, 6, 0, 0,
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

// SwiftLexerInit initializes any static state used to implement SwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.once.Do(swiftlexerLexerInit)
}

// NewSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLexer(input antlr.CharStream) *SwiftLexer {
	SwiftLexerInit()
	l := new(SwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLexer tokens.
const (
	SwiftLexerINT          = 1
	SwiftLexerFLOAT        = 2
	SwiftLexerBOOL         = 3
	SwiftLexerSTR          = 4
	SwiftLexerTRU          = 5
	SwiftLexerFAL          = 6
	SwiftLexerPRINT        = 7
	SwiftLexerIF           = 8
	SwiftLexerELSE         = 9
	SwiftLexerWHILE        = 10
	SwiftLexerFOR          = 11
	SwiftLexerIN           = 12
	SwiftLexerVAR          = 13
	SwiftLexerNUMBER       = 14
	SwiftLexerSTRING       = 15
	SwiftLexerID           = 16
	SwiftLexerDIF          = 17
	SwiftLexerIG_IG        = 18
	SwiftLexerNOT          = 19
	SwiftLexerOR           = 20
	SwiftLexerAND          = 21
	SwiftLexerIG           = 22
	SwiftLexerMAY_IG       = 23
	SwiftLexerMEN_IG       = 24
	SwiftLexerMAYOR        = 25
	SwiftLexerMENOR        = 26
	SwiftLexerMUL          = 27
	SwiftLexerDIV          = 28
	SwiftLexerADD          = 29
	SwiftLexerSUB          = 30
	SwiftLexerPARIZQ       = 31
	SwiftLexerPARDER       = 32
	SwiftLexerLLAVEIZQ     = 33
	SwiftLexerLLAVEDER     = 34
	SwiftLexerD_PTS        = 35
	SwiftLexerCORIZQ       = 36
	SwiftLexerCORDER       = 37
	SwiftLexerCOMA         = 38
	SwiftLexerPUNTO        = 39
	SwiftLexerWHITESPACE   = 40
	SwiftLexerCOMMENT      = 41
	SwiftLexerLINE_COMMENT = 42
)