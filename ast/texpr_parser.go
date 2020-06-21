// Code generated from TExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package ast // TExpr
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 265,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 90, 10, 3, 12,
	3, 14, 3, 93, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 98, 10, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 110, 10, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 115, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 120, 10, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 125, 10, 9, 3, 9, 5, 9, 128, 10, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 141, 10,
	11, 12, 11, 14, 11, 144, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 158, 10, 12, 12, 12,
	14, 12, 161, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 168, 10,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 177, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 184, 10, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 196, 10, 18,
	12, 18, 14, 18, 199, 11, 18, 3, 19, 5, 19, 202, 10, 19, 3, 19, 6, 19, 205,
	10, 19, 13, 19, 14, 19, 206, 3, 19, 3, 19, 6, 19, 211, 10, 19, 13, 19,
	14, 19, 212, 5, 19, 215, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 7, 22, 224, 10, 22, 12, 22, 14, 22, 227, 11, 22, 3, 22, 5, 22,
	230, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 235, 10, 23, 12, 23, 14, 23, 238,
	11, 23, 3, 23, 5, 23, 241, 10, 23, 3, 24, 3, 24, 3, 24, 7, 24, 246, 10,
	24, 12, 24, 14, 24, 249, 11, 24, 3, 24, 5, 24, 252, 10, 24, 3, 25, 3, 25,
	3, 25, 7, 25, 257, 10, 25, 12, 25, 14, 25, 260, 11, 25, 3, 25, 5, 25, 263,
	10, 25, 3, 25, 2, 5, 4, 20, 22, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 6, 3, 2, 18,
	23, 3, 2, 11, 12, 3, 2, 16, 17, 3, 2, 6, 9, 2, 282, 2, 50, 3, 2, 2, 2,
	4, 62, 3, 2, 2, 2, 6, 97, 3, 2, 2, 2, 8, 99, 3, 2, 2, 2, 10, 101, 3, 2,
	2, 2, 12, 103, 3, 2, 2, 2, 14, 105, 3, 2, 2, 2, 16, 127, 3, 2, 2, 2, 18,
	129, 3, 2, 2, 2, 20, 131, 3, 2, 2, 2, 22, 145, 3, 2, 2, 2, 24, 167, 3,
	2, 2, 2, 26, 176, 3, 2, 2, 2, 28, 183, 3, 2, 2, 2, 30, 185, 3, 2, 2, 2,
	32, 190, 3, 2, 2, 2, 34, 192, 3, 2, 2, 2, 36, 201, 3, 2, 2, 2, 38, 216,
	3, 2, 2, 2, 40, 218, 3, 2, 2, 2, 42, 220, 3, 2, 2, 2, 44, 231, 3, 2, 2,
	2, 46, 242, 3, 2, 2, 2, 48, 253, 3, 2, 2, 2, 50, 51, 5, 4, 3, 2, 51, 52,
	7, 2, 2, 3, 52, 3, 3, 2, 2, 2, 53, 54, 8, 3, 1, 2, 54, 55, 7, 25, 2, 2,
	55, 56, 5, 4, 3, 2, 56, 57, 7, 26, 2, 2, 57, 63, 3, 2, 2, 2, 58, 59, 7,
	13, 2, 2, 59, 63, 5, 4, 3, 12, 60, 63, 5, 6, 4, 2, 61, 63, 5, 18, 10, 2,
	62, 53, 3, 2, 2, 2, 62, 58, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3,
	2, 2, 2, 63, 91, 3, 2, 2, 2, 64, 65, 12, 7, 2, 2, 65, 66, 5, 8, 5, 2, 66,
	67, 5, 4, 3, 8, 67, 90, 3, 2, 2, 2, 68, 69, 12, 6, 2, 2, 69, 70, 5, 10,
	6, 2, 70, 71, 5, 4, 3, 7, 71, 90, 3, 2, 2, 2, 72, 73, 12, 11, 2, 2, 73,
	74, 7, 15, 2, 2, 74, 90, 5, 16, 9, 2, 75, 76, 12, 10, 2, 2, 76, 77, 7,
	13, 2, 2, 77, 78, 7, 15, 2, 2, 78, 90, 5, 16, 9, 2, 79, 80, 12, 9, 2, 2,
	80, 81, 7, 14, 2, 2, 81, 90, 5, 40, 21, 2, 82, 83, 12, 8, 2, 2, 83, 84,
	7, 14, 2, 2, 84, 85, 7, 13, 2, 2, 85, 90, 5, 40, 21, 2, 86, 87, 12, 5,
	2, 2, 87, 88, 7, 24, 2, 2, 88, 90, 5, 38, 20, 2, 89, 64, 3, 2, 2, 2, 89,
	68, 3, 2, 2, 2, 89, 72, 3, 2, 2, 2, 89, 75, 3, 2, 2, 2, 89, 79, 3, 2, 2,
	2, 89, 82, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 5, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2,
	94, 98, 7, 30, 2, 2, 95, 98, 5, 14, 8, 2, 96, 98, 5, 16, 9, 2, 97, 94,
	3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 7, 3, 2, 2, 2,
	99, 100, 9, 2, 2, 2, 100, 9, 3, 2, 2, 2, 101, 102, 9, 3, 2, 2, 102, 11,
	3, 2, 2, 2, 103, 104, 9, 4, 2, 2, 104, 13, 3, 2, 2, 2, 105, 106, 9, 5,
	2, 2, 106, 15, 3, 2, 2, 2, 107, 109, 7, 3, 2, 2, 108, 110, 5, 44, 23, 2,
	109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	128, 7, 4, 2, 2, 112, 114, 7, 3, 2, 2, 113, 115, 5, 42, 22, 2, 114, 113,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 128, 7, 4,
	2, 2, 117, 119, 7, 3, 2, 2, 118, 120, 5, 46, 24, 2, 119, 118, 3, 2, 2,
	2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 128, 7, 4, 2, 2, 122,
	124, 7, 3, 2, 2, 123, 125, 5, 48, 25, 2, 124, 123, 3, 2, 2, 2, 124, 125,
	3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128, 7, 4, 2, 2, 127, 107, 3, 2,
	2, 2, 127, 112, 3, 2, 2, 2, 127, 117, 3, 2, 2, 2, 127, 122, 3, 2, 2, 2,
	128, 17, 3, 2, 2, 2, 129, 130, 5, 20, 11, 2, 130, 19, 3, 2, 2, 2, 131,
	132, 8, 11, 1, 2, 132, 133, 5, 22, 12, 2, 133, 142, 3, 2, 2, 2, 134, 135,
	12, 5, 2, 2, 135, 136, 7, 39, 2, 2, 136, 141, 5, 22, 12, 2, 137, 138, 12,
	4, 2, 2, 138, 139, 7, 40, 2, 2, 139, 141, 5, 22, 12, 2, 140, 134, 3, 2,
	2, 2, 140, 137, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 21, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146,
	8, 12, 1, 2, 146, 147, 5, 24, 13, 2, 147, 159, 3, 2, 2, 2, 148, 149, 12,
	6, 2, 2, 149, 150, 7, 41, 2, 2, 150, 158, 5, 24, 13, 2, 151, 152, 12, 5,
	2, 2, 152, 153, 7, 42, 2, 2, 153, 158, 5, 24, 13, 2, 154, 155, 12, 4, 2,
	2, 155, 156, 7, 43, 2, 2, 156, 158, 5, 24, 13, 2, 157, 148, 3, 2, 2, 2,
	157, 151, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159,
	157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 23, 3, 2, 2, 2, 161, 159, 3,
	2, 2, 2, 162, 163, 5, 26, 14, 2, 163, 164, 7, 46, 2, 2, 164, 165, 5, 26,
	14, 2, 165, 168, 3, 2, 2, 2, 166, 168, 5, 26, 14, 2, 167, 162, 3, 2, 2,
	2, 167, 166, 3, 2, 2, 2, 168, 25, 3, 2, 2, 2, 169, 177, 5, 6, 4, 2, 170,
	177, 5, 28, 15, 2, 171, 172, 7, 25, 2, 2, 172, 173, 5, 20, 11, 2, 173,
	174, 7, 26, 2, 2, 174, 177, 3, 2, 2, 2, 175, 177, 5, 30, 16, 2, 176, 169,
	3, 2, 2, 2, 176, 170, 3, 2, 2, 2, 176, 171, 3, 2, 2, 2, 176, 175, 3, 2,
	2, 2, 177, 27, 3, 2, 2, 2, 178, 179, 5, 36, 19, 2, 179, 180, 7, 45, 2,
	2, 180, 181, 5, 36, 19, 2, 181, 184, 3, 2, 2, 2, 182, 184, 5, 36, 19, 2,
	183, 178, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 29, 3, 2, 2, 2, 185, 186,
	5, 32, 17, 2, 186, 187, 7, 25, 2, 2, 187, 188, 5, 34, 18, 2, 188, 189,
	7, 26, 2, 2, 189, 31, 3, 2, 2, 2, 190, 191, 7, 29, 2, 2, 191, 33, 3, 2,
	2, 2, 192, 197, 5, 4, 3, 2, 193, 194, 7, 5, 2, 2, 194, 196, 5, 4, 3, 2,
	195, 193, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197,
	198, 3, 2, 2, 2, 198, 35, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 202, 7,
	40, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2,
	2, 203, 205, 7, 48, 2, 2, 204, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206,
	204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 214, 3, 2, 2, 2, 208, 210,
	7, 44, 2, 2, 209, 211, 7, 48, 2, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3,
	2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2,
	2, 214, 208, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 37, 3, 2, 2, 2, 216,
	217, 7, 10, 2, 2, 217, 39, 3, 2, 2, 2, 218, 219, 7, 29, 2, 2, 219, 41,
	3, 2, 2, 2, 220, 225, 7, 9, 2, 2, 221, 222, 7, 5, 2, 2, 222, 224, 7, 9,
	2, 2, 223, 221, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2,
	225, 226, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228,
	230, 7, 5, 2, 2, 229, 228, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 43, 3,
	2, 2, 2, 231, 236, 7, 6, 2, 2, 232, 233, 7, 5, 2, 2, 233, 235, 7, 6, 2,
	2, 234, 232, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236,
	237, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 241,
	7, 5, 2, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 45, 3, 2,
	2, 2, 242, 247, 7, 7, 2, 2, 243, 244, 7, 5, 2, 2, 244, 246, 7, 7, 2, 2,
	245, 243, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 252,
	7, 5, 2, 2, 251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 47, 3, 2,
	2, 2, 253, 258, 7, 8, 2, 2, 254, 255, 7, 5, 2, 2, 255, 257, 7, 8, 2, 2,
	256, 254, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258,
	259, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261, 263,
	7, 5, 2, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 49, 3, 2,
	2, 2, 31, 62, 89, 91, 97, 109, 114, 119, 124, 127, 140, 142, 157, 159,
	167, 176, 183, 197, 201, 206, 212, 214, 225, 229, 236, 240, 247, 251, 258,
	262,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "','", "", "", "", "", "", "'&&'", "'||'", "'not'", "'is'",
	"'in'", "'true'", "'false'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='",
	"'=~'", "'('", "')'", "", "", "", "", "'cos'", "'sin'", "'tan'", "'acos'",
	"'asin'", "'atan'", "'ln'", "'log'", "'+'", "'-'", "'*'", "'/'", "'%'",
	"'.'", "", "'^'", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "Integer", "Float", "Boolean", "Varchar", "Regex", "AND",
	"OR", "NOT", "IS", "IN", "TRUE", "FALSE", "GT", "GE", "LT", "LE", "EQ",
	"NE", "MATCH", "LPAREN", "RPAREN", "INT", "FLOAT", "IDENTIFIER", "VARIABLE",
	"COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "PLUS", "MINUS",
	"MUL", "DIV", "MOD", "POINT", "E", "POW", "NL", "DIGIT", "WS",
}

var ruleNames = []string{
	"parse", "expression", "variable", "comparator", "binary", "boolean", "literal",
	"array", "calc", "plus", "multiplying", "pow", "atom", "scientific", "function",
	"funcname", "parameters", "number", "regex", "kind", "strings", "integers",
	"floats", "booleans",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TExprParser struct {
	*antlr.BaseParser
}

func NewTExprParser(input antlr.TokenStream) *TExprParser {
	this := new(TExprParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TExpr.g4"

	return this
}

// TExprParser tokens.
const (
	TExprParserEOF        = antlr.TokenEOF
	TExprParserT__0       = 1
	TExprParserT__1       = 2
	TExprParserT__2       = 3
	TExprParserInteger    = 4
	TExprParserFloat      = 5
	TExprParserBoolean    = 6
	TExprParserVarchar    = 7
	TExprParserRegex      = 8
	TExprParserAND        = 9
	TExprParserOR         = 10
	TExprParserNOT        = 11
	TExprParserIS         = 12
	TExprParserIN         = 13
	TExprParserTRUE       = 14
	TExprParserFALSE      = 15
	TExprParserGT         = 16
	TExprParserGE         = 17
	TExprParserLT         = 18
	TExprParserLE         = 19
	TExprParserEQ         = 20
	TExprParserNE         = 21
	TExprParserMATCH      = 22
	TExprParserLPAREN     = 23
	TExprParserRPAREN     = 24
	TExprParserINT        = 25
	TExprParserFLOAT      = 26
	TExprParserIDENTIFIER = 27
	TExprParserVARIABLE   = 28
	TExprParserCOS        = 29
	TExprParserSIN        = 30
	TExprParserTAN        = 31
	TExprParserACOS       = 32
	TExprParserASIN       = 33
	TExprParserATAN       = 34
	TExprParserLN         = 35
	TExprParserLOG        = 36
	TExprParserPLUS       = 37
	TExprParserMINUS      = 38
	TExprParserMUL        = 39
	TExprParserDIV        = 40
	TExprParserMOD        = 41
	TExprParserPOINT      = 42
	TExprParserE          = 43
	TExprParserPOW        = 44
	TExprParserNL         = 45
	TExprParserDIGIT      = 46
	TExprParserWS         = 47
)

// TExprParser rules.
const (
	TExprParserRULE_parse       = 0
	TExprParserRULE_expression  = 1
	TExprParserRULE_variable    = 2
	TExprParserRULE_comparator  = 3
	TExprParserRULE_binary      = 4
	TExprParserRULE_boolean     = 5
	TExprParserRULE_literal     = 6
	TExprParserRULE_array       = 7
	TExprParserRULE_calc        = 8
	TExprParserRULE_plus        = 9
	TExprParserRULE_multiplying = 10
	TExprParserRULE_pow         = 11
	TExprParserRULE_atom        = 12
	TExprParserRULE_scientific  = 13
	TExprParserRULE_function    = 14
	TExprParserRULE_funcname    = 15
	TExprParserRULE_parameters  = 16
	TExprParserRULE_number      = 17
	TExprParserRULE_regex       = 18
	TExprParserRULE_kind        = 19
	TExprParserRULE_strings     = 20
	TExprParserRULE_integers    = 21
	TExprParserRULE_floats      = 22
	TExprParserRULE_booleans    = 23
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(TExprParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TExprParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.expression(0)
	}
	{
		p.SetState(49)
		p.Match(TExprParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExpressionContext struct {
	*ExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) Binary() IBinaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExpressionContext struct {
	*ExpressionContext
}

func NewMatchExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExpressionContext {
	var p = new(MatchExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MatchExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MatchExpressionContext) MATCH() antlr.TerminalNode {
	return s.GetToken(TExprParserMATCH, 0)
}

func (s *MatchExpressionContext) Regex() IRegexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegexContext)
}

func (s *MatchExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitMatchExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(TExprParserIN, 0)
}

func (s *InExpressionContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsTypeExpressionContext struct {
	*ExpressionContext
}

func NewIsTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsTypeExpressionContext {
	var p = new(IsTypeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IsTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsTypeExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsTypeExpressionContext) IS() antlr.TerminalNode {
	return s.GetToken(TExprParserIS, 0)
}

func (s *IsTypeExpressionContext) Kind() IKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKindContext)
}

func (s *IsTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitIsTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CalcExpressionContext struct {
	*ExpressionContext
}

func NewCalcExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalcExpressionContext {
	var p = new(CalcExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CalcExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcExpressionContext) Calc() ICalcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalcContext)
}

func (s *CalcExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitCalcExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNotTypeExpressionContext struct {
	*ExpressionContext
}

func NewIsNotTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNotTypeExpressionContext {
	var p = new(IsNotTypeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IsNotTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNotTypeExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsNotTypeExpressionContext) IS() antlr.TerminalNode {
	return s.GetToken(TExprParserIS, 0)
}

func (s *IsNotTypeExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TExprParserNOT, 0)
}

func (s *IsNotTypeExpressionContext) Kind() IKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKindContext)
}

func (s *IsNotTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitIsNotTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TExprParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	*ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserLPAREN, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserRPAREN, 0)
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotInExpressionContext struct {
	*ExpressionContext
}

func NewNotInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInExpressionContext {
	var p = new(NotInExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotInExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotInExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TExprParserNOT, 0)
}

func (s *NotInExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(TExprParserIN, 0)
}

func (s *NotInExpressionContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *NotInExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitNotInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparatorExpressionContext struct {
	*ExpressionContext
}

func NewComparatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorExpressionContext {
	var p = new(ComparatorExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ComparatorExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparatorExpressionContext) Comparator() IComparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ComparatorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitComparatorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableExpressionContext struct {
	*ExpressionContext
}

func NewVariableExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableExpressionContext {
	var p = new(VariableExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VariableExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExpressionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitVariableExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TExprParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TExprParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(52)
			p.Match(TExprParserLPAREN)
		}
		{
			p.SetState(53)
			p.expression(0)
		}
		{
			p.SetState(54)
			p.Match(TExprParserRPAREN)
		}

	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.Match(TExprParserNOT)
		}
		{
			p.SetState(57)
			p.expression(10)
		}

	case 3:
		localctx = NewVariableExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Variable()
		}

	case 4:
		localctx = NewCalcExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Calc()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(63)
					p.Comparator()
				}
				{
					p.SetState(64)
					p.expression(6)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(67)
					p.Binary()
				}
				{
					p.SetState(68)
					p.expression(5)
				}

			case 3:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(71)
					p.Match(TExprParserIN)
				}
				{
					p.SetState(72)
					p.Array()
				}

			case 4:
				localctx = NewNotInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(74)
					p.Match(TExprParserNOT)
				}
				{
					p.SetState(75)
					p.Match(TExprParserIN)
				}
				{
					p.SetState(76)
					p.Array()
				}

			case 5:
				localctx = NewIsTypeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(78)
					p.Match(TExprParserIS)
				}
				{
					p.SetState(79)
					p.Kind()
				}

			case 6:
				localctx = NewIsNotTypeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(81)
					p.Match(TExprParserIS)
				}
				{
					p.SetState(82)
					p.Match(TExprParserNOT)
				}
				{
					p.SetState(83)
					p.Kind()
				}

			case 7:
				localctx = NewMatchExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(85)
					p.Match(TExprParserMATCH)
				}
				{
					p.SetState(86)
					p.Regex()
				}

			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(TExprParserVARIABLE, 0)
}

func (s *VariableContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *VariableContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TExprParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TExprParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(TExprParserVARIABLE)
		}

	case TExprParserInteger, TExprParserFloat, TExprParserBoolean, TExprParserVarchar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Literal()
		}

	case TExprParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(TExprParserGT, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(TExprParserGE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(TExprParserLT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(TExprParserLE, 0)
}

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(TExprParserEQ, 0)
}

func (s *ComparatorContext) NE() antlr.TerminalNode {
	return s.GetToken(TExprParserNE, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TExprParserRULE_comparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TExprParserGT)|(1<<TExprParserGE)|(1<<TExprParserLT)|(1<<TExprParserLE)|(1<<TExprParserEQ)|(1<<TExprParserNE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(TExprParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(TExprParserOR, 0)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TExprParserRULE_binary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TExprParserAND || _la == TExprParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_boolean
	return p
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TExprParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TExprParserFALSE, 0)
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Boolean() (localctx IBooleanContext) {
	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TExprParserRULE_boolean)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TExprParserTRUE || _la == TExprParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Varchar() antlr.TerminalNode {
	return s.GetToken(TExprParserVarchar, 0)
}

func (s *LiteralContext) Integer() antlr.TerminalNode {
	return s.GetToken(TExprParserInteger, 0)
}

func (s *LiteralContext) Float() antlr.TerminalNode {
	return s.GetToken(TExprParserFloat, 0)
}

func (s *LiteralContext) Boolean() antlr.TerminalNode {
	return s.GetToken(TExprParserBoolean, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TExprParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TExprParserInteger)|(1<<TExprParserFloat)|(1<<TExprParserBoolean)|(1<<TExprParserVarchar))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) Integers() IIntegersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegersContext)
}

func (s *ArrayContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *ArrayContext) Floats() IFloatsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatsContext)
}

func (s *ArrayContext) Booleans() IBooleansContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleansContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleansContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TExprParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(TExprParserT__0)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserInteger {
			{
				p.SetState(106)
				p.Integers()
			}

		}
		{
			p.SetState(109)
			p.Match(TExprParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(TExprParserT__0)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserVarchar {
			{
				p.SetState(111)
				p.Strings()
			}

		}
		{
			p.SetState(114)
			p.Match(TExprParserT__1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(TExprParserT__0)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserFloat {
			{
				p.SetState(116)
				p.Floats()
			}

		}
		{
			p.SetState(119)
			p.Match(TExprParserT__1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Match(TExprParserT__0)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserBoolean {
			{
				p.SetState(121)
				p.Booleans()
			}

		}
		{
			p.SetState(124)
			p.Match(TExprParserT__1)
		}

	}

	return localctx
}

// ICalcContext is an interface to support dynamic dispatch.
type ICalcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalcContext differentiates from other interfaces.
	IsCalcContext()
}

type CalcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalcContext() *CalcContext {
	var p = new(CalcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_calc
	return p
}

func (*CalcContext) IsCalcContext() {}

func NewCalcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcContext {
	var p = new(CalcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_calc

	return p
}

func (s *CalcContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcContext) Plus() IPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusContext)
}

func (s *CalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitCalc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Calc() (localctx ICalcContext) {
	localctx = NewCalcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TExprParserRULE_calc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.plus(0)
	}

	return localctx
}

// IPlusContext is an interface to support dynamic dispatch.
type IPlusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlusContext differentiates from other interfaces.
	IsPlusContext()
}

type PlusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlusContext() *PlusContext {
	var p = new(PlusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_plus
	return p
}

func (*PlusContext) IsPlusContext() {}

func NewPlusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlusContext {
	var p = new(PlusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_plus

	return p
}

func (s *PlusContext) GetParser() antlr.Parser { return s.parser }

func (s *PlusContext) Multiplying() IMultiplyingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingContext)
}

func (s *PlusContext) Plus() IPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusContext)
}

func (s *PlusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TExprParserPLUS, 0)
}

func (s *PlusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TExprParserMINUS, 0)
}

func (s *PlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitPlus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Plus() (localctx IPlusContext) {
	return p.plus(0)
}

func (p *TExprParser) plus(_p int) (localctx IPlusContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPlusContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPlusContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, TExprParserRULE_plus, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.multiplying(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPlusContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_plus)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(133)
					p.Match(TExprParserPLUS)
				}
				{
					p.SetState(134)
					p.multiplying(0)
				}

			case 2:
				localctx = NewPlusContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_plus)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(136)
					p.Match(TExprParserMINUS)
				}
				{
					p.SetState(137)
					p.multiplying(0)
				}

			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplyingContext is an interface to support dynamic dispatch.
type IMultiplyingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingContext differentiates from other interfaces.
	IsMultiplyingContext()
}

type MultiplyingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingContext() *MultiplyingContext {
	var p = new(MultiplyingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_multiplying
	return p
}

func (*MultiplyingContext) IsMultiplyingContext() {}

func NewMultiplyingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingContext {
	var p = new(MultiplyingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_multiplying

	return p
}

func (s *MultiplyingContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *MultiplyingContext) Multiplying() IMultiplyingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingContext)
}

func (s *MultiplyingContext) MUL() antlr.TerminalNode {
	return s.GetToken(TExprParserMUL, 0)
}

func (s *MultiplyingContext) DIV() antlr.TerminalNode {
	return s.GetToken(TExprParserDIV, 0)
}

func (s *MultiplyingContext) MOD() antlr.TerminalNode {
	return s.GetToken(TExprParserMOD, 0)
}

func (s *MultiplyingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitMultiplying(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Multiplying() (localctx IMultiplyingContext) {
	return p.multiplying(0)
}

func (p *TExprParser) multiplying(_p int) (localctx IMultiplyingContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplyingContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplyingContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, TExprParserRULE_multiplying, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Pow()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(147)
					p.Match(TExprParserMUL)
				}
				{
					p.SetState(148)
					p.Pow()
				}

			case 2:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(150)
					p.Match(TExprParserDIV)
				}
				{
					p.SetState(151)
					p.Pow()
				}

			case 3:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(153)
					p.Match(TExprParserMOD)
				}
				{
					p.SetState(154)
					p.Pow()
				}

			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IPowContext is an interface to support dynamic dispatch.
type IPowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowContext differentiates from other interfaces.
	IsPowContext()
}

type PowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowContext() *PowContext {
	var p = new(PowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_pow
	return p
}

func (*PowContext) IsPowContext() {}

func NewPowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowContext {
	var p = new(PowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_pow

	return p
}

func (s *PowContext) GetParser() antlr.Parser { return s.parser }

func (s *PowContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *PowContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *PowContext) POW() antlr.TerminalNode {
	return s.GetToken(TExprParserPOW, 0)
}

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitPow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Pow() (localctx IPowContext) {
	localctx = NewPowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TExprParserRULE_pow)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Atom()
		}
		{
			p.SetState(161)
			p.Match(TExprParserPOW)
		}
		{
			p.SetState(162)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Atom()
		}

	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) Scientific() IScientificContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScientificContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScientificContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserLPAREN, 0)
}

func (s *AtomContext) Plus() IPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserRPAREN, 0)
}

func (s *AtomContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TExprParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TExprParserT__0, TExprParserInteger, TExprParserFloat, TExprParserBoolean, TExprParserVarchar, TExprParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Variable()
		}

	case TExprParserMINUS, TExprParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Scientific()
		}

	case TExprParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Match(TExprParserLPAREN)
		}
		{
			p.SetState(170)
			p.plus(0)
		}
		{
			p.SetState(171)
			p.Match(TExprParserRPAREN)
		}

	case TExprParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.Function()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IScientificContext is an interface to support dynamic dispatch.
type IScientificContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScientificContext differentiates from other interfaces.
	IsScientificContext()
}

type ScientificContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScientificContext() *ScientificContext {
	var p = new(ScientificContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_scientific
	return p
}

func (*ScientificContext) IsScientificContext() {}

func NewScientificContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScientificContext {
	var p = new(ScientificContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_scientific

	return p
}

func (s *ScientificContext) GetParser() antlr.Parser { return s.parser }

func (s *ScientificContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *ScientificContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ScientificContext) E() antlr.TerminalNode {
	return s.GetToken(TExprParserE, 0)
}

func (s *ScientificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScientificContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitScientific(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Scientific() (localctx IScientificContext) {
	localctx = NewScientificContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TExprParserRULE_scientific)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Number()
		}
		{
			p.SetState(177)
			p.Match(TExprParserE)
		}
		{
			p.SetState(178)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Number()
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserLPAREN, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TExprParserRPAREN, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TExprParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Funcname()
	}
	{
		p.SetState(184)
		p.Match(TExprParserLPAREN)
	}
	{
		p.SetState(185)
		p.Parameters()
	}
	{
		p.SetState(186)
		p.Match(TExprParserRPAREN)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TExprParserIDENTIFIER, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitFuncname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TExprParserRULE_funcname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(TExprParserIDENTIFIER)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ParametersContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TExprParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.expression(0)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TExprParserT__2 {
		{
			p.SetState(191)
			p.Match(TExprParserT__2)
		}
		{
			p.SetState(192)
			p.expression(0)
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TExprParserMINUS, 0)
}

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(TExprParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserDIGIT, i)
}

func (s *NumberContext) POINT() antlr.TerminalNode {
	return s.GetToken(TExprParserPOINT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TExprParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserMINUS {
		{
			p.SetState(198)
			p.Match(TExprParserMINUS)
		}

	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(201)
				p.Match(TExprParserDIGIT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(206)
			p.Match(TExprParserPOINT)
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(207)
					p.Match(TExprParserDIGIT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IRegexContext is an interface to support dynamic dispatch.
type IRegexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexContext differentiates from other interfaces.
	IsRegexContext()
}

type RegexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexContext() *RegexContext {
	var p = new(RegexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_regex
	return p
}

func (*RegexContext) IsRegexContext() {}

func NewRegexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexContext {
	var p = new(RegexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_regex

	return p
}

func (s *RegexContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexContext) Regex() antlr.TerminalNode {
	return s.GetToken(TExprParserRegex, 0)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitRegex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Regex() (localctx IRegexContext) {
	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TExprParserRULE_regex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(TExprParserRegex)
	}

	return localctx
}

// IKindContext is an interface to support dynamic dispatch.
type IKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKindContext differentiates from other interfaces.
	IsKindContext()
}

type KindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKindContext() *KindContext {
	var p = new(KindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_kind
	return p
}

func (*KindContext) IsKindContext() {}

func NewKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KindContext {
	var p = new(KindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_kind

	return p
}

func (s *KindContext) GetParser() antlr.Parser { return s.parser }

func (s *KindContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TExprParserIDENTIFIER, 0)
}

func (s *KindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Kind() (localctx IKindContext) {
	localctx = NewKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TExprParserRULE_kind)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(TExprParserIDENTIFIER)
	}

	return localctx
}

// IStringsContext is an interface to support dynamic dispatch.
type IStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringsContext differentiates from other interfaces.
	IsStringsContext()
}

type StringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringsContext() *StringsContext {
	var p = new(StringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_strings
	return p
}

func (*StringsContext) IsStringsContext() {}

func NewStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringsContext {
	var p = new(StringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_strings

	return p
}

func (s *StringsContext) GetParser() antlr.Parser { return s.parser }

func (s *StringsContext) AllVarchar() []antlr.TerminalNode {
	return s.GetTokens(TExprParserVarchar)
}

func (s *StringsContext) Varchar(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserVarchar, i)
}

func (s *StringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Strings() (localctx IStringsContext) {
	localctx = NewStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TExprParserRULE_strings)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(TExprParserVarchar)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(219)
				p.Match(TExprParserT__2)
			}
			{
				p.SetState(220)
				p.Match(TExprParserVarchar)
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserT__2 {
		{
			p.SetState(226)
			p.Match(TExprParserT__2)
		}

	}

	return localctx
}

// IIntegersContext is an interface to support dynamic dispatch.
type IIntegersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegersContext differentiates from other interfaces.
	IsIntegersContext()
}

type IntegersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegersContext() *IntegersContext {
	var p = new(IntegersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_integers
	return p
}

func (*IntegersContext) IsIntegersContext() {}

func NewIntegersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegersContext {
	var p = new(IntegersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_integers

	return p
}

func (s *IntegersContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegersContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(TExprParserInteger)
}

func (s *IntegersContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserInteger, i)
}

func (s *IntegersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitIntegers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Integers() (localctx IIntegersContext) {
	localctx = NewIntegersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TExprParserRULE_integers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(TExprParserInteger)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				p.Match(TExprParserT__2)
			}
			{
				p.SetState(231)
				p.Match(TExprParserInteger)
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserT__2 {
		{
			p.SetState(237)
			p.Match(TExprParserT__2)
		}

	}

	return localctx
}

// IFloatsContext is an interface to support dynamic dispatch.
type IFloatsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatsContext differentiates from other interfaces.
	IsFloatsContext()
}

type FloatsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatsContext() *FloatsContext {
	var p = new(FloatsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_floats
	return p
}

func (*FloatsContext) IsFloatsContext() {}

func NewFloatsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatsContext {
	var p = new(FloatsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_floats

	return p
}

func (s *FloatsContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatsContext) AllFloat() []antlr.TerminalNode {
	return s.GetTokens(TExprParserFloat)
}

func (s *FloatsContext) Float(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserFloat, i)
}

func (s *FloatsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitFloats(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Floats() (localctx IFloatsContext) {
	localctx = NewFloatsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TExprParserRULE_floats)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(TExprParserFloat)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)
				p.Match(TExprParserT__2)
			}
			{
				p.SetState(242)
				p.Match(TExprParserFloat)
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserT__2 {
		{
			p.SetState(248)
			p.Match(TExprParserT__2)
		}

	}

	return localctx
}

// IBooleansContext is an interface to support dynamic dispatch.
type IBooleansContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleansContext differentiates from other interfaces.
	IsBooleansContext()
}

type BooleansContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleansContext() *BooleansContext {
	var p = new(BooleansContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_booleans
	return p
}

func (*BooleansContext) IsBooleansContext() {}

func NewBooleansContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleansContext {
	var p = new(BooleansContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_booleans

	return p
}

func (s *BooleansContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleansContext) AllBoolean() []antlr.TerminalNode {
	return s.GetTokens(TExprParserBoolean)
}

func (s *BooleansContext) Boolean(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserBoolean, i)
}

func (s *BooleansContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleansContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleansContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprVisitor:
		return t.VisitBooleans(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Booleans() (localctx IBooleansContext) {
	localctx = NewBooleansContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TExprParserRULE_booleans)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(TExprParserBoolean)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(252)
				p.Match(TExprParserT__2)
			}
			{
				p.SetState(253)
				p.Match(TExprParserBoolean)
			}

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserT__2 {
		{
			p.SetState(259)
			p.Match(TExprParserT__2)
		}

	}

	return localctx
}

func (p *TExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 9:
		var t *PlusContext = nil
		if localctx != nil {
			t = localctx.(*PlusContext)
		}
		return p.Plus_Sempred(t, predIndex)

	case 10:
		var t *MultiplyingContext = nil
		if localctx != nil {
			t = localctx.(*MultiplyingContext)
		}
		return p.Multiplying_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TExprParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Plus_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Multiplying_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
