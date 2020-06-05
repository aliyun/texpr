// Code generated from TExprParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package ast // TExprParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 301,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 67, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 94, 10, 3, 12, 3, 14, 3, 97, 11, 3, 3, 4, 3, 4, 3, 4, 5,
	4, 102, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 5, 9, 115, 10, 9, 3, 10, 3, 10, 5, 10, 119, 10, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 124, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 129, 10, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 134, 10, 10, 3, 10, 5, 10, 137, 10, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 153, 10, 12, 12, 12, 14, 12, 156, 11, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 170, 10, 13, 12, 13, 14, 13, 173, 11, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 184, 10, 14, 12, 14,
	14, 14, 187, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 201, 10, 15, 12, 15, 14, 15, 204,
	11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 213, 10,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 220, 10, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 232, 10,
	20, 12, 20, 14, 20, 235, 11, 20, 3, 21, 5, 21, 238, 10, 21, 3, 21, 6, 21,
	241, 10, 21, 13, 21, 14, 21, 242, 3, 21, 3, 21, 6, 21, 247, 10, 21, 13,
	21, 14, 21, 248, 5, 21, 251, 10, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 7, 24, 260, 10, 24, 12, 24, 14, 24, 263, 11, 24, 3, 24, 5,
	24, 266, 10, 24, 3, 25, 3, 25, 3, 25, 7, 25, 271, 10, 25, 12, 25, 14, 25,
	274, 11, 25, 3, 25, 5, 25, 277, 10, 25, 3, 26, 3, 26, 3, 26, 7, 26, 282,
	10, 26, 12, 26, 14, 26, 285, 11, 26, 3, 26, 5, 26, 288, 10, 26, 3, 27,
	3, 27, 3, 27, 7, 27, 293, 10, 27, 12, 27, 14, 27, 296, 11, 27, 3, 27, 5,
	27, 299, 10, 27, 3, 27, 2, 7, 4, 22, 24, 26, 28, 28, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 2, 6, 3, 2, 14, 19, 3, 2, 7, 8, 3, 2, 12, 13, 3, 2, 3, 6, 2, 323,
	2, 54, 3, 2, 2, 2, 4, 66, 3, 2, 2, 2, 6, 101, 3, 2, 2, 2, 8, 103, 3, 2,
	2, 2, 10, 105, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16,
	114, 3, 2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 138, 3, 2, 2, 2, 22, 140, 3,
	2, 2, 2, 24, 157, 3, 2, 2, 2, 26, 174, 3, 2, 2, 2, 28, 188, 3, 2, 2, 2,
	30, 212, 3, 2, 2, 2, 32, 219, 3, 2, 2, 2, 34, 221, 3, 2, 2, 2, 36, 226,
	3, 2, 2, 2, 38, 228, 3, 2, 2, 2, 40, 237, 3, 2, 2, 2, 42, 252, 3, 2, 2,
	2, 44, 254, 3, 2, 2, 2, 46, 256, 3, 2, 2, 2, 48, 267, 3, 2, 2, 2, 50, 278,
	3, 2, 2, 2, 52, 289, 3, 2, 2, 2, 54, 55, 5, 4, 3, 2, 55, 56, 7, 2, 2, 3,
	56, 3, 3, 2, 2, 2, 57, 58, 8, 3, 1, 2, 58, 59, 7, 21, 2, 2, 59, 60, 5,
	4, 3, 2, 60, 61, 7, 22, 2, 2, 61, 67, 3, 2, 2, 2, 62, 63, 7, 9, 2, 2, 63,
	67, 5, 4, 3, 12, 64, 67, 5, 6, 4, 2, 65, 67, 5, 20, 11, 2, 66, 57, 3, 2,
	2, 2, 66, 62, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 95,
	3, 2, 2, 2, 68, 69, 12, 6, 2, 2, 69, 70, 5, 8, 5, 2, 70, 71, 5, 4, 3, 7,
	71, 94, 3, 2, 2, 2, 72, 73, 12, 5, 2, 2, 73, 74, 5, 10, 6, 2, 74, 75, 5,
	4, 3, 6, 75, 94, 3, 2, 2, 2, 76, 77, 12, 11, 2, 2, 77, 78, 7, 11, 2, 2,
	78, 94, 5, 16, 9, 2, 79, 80, 12, 10, 2, 2, 80, 81, 7, 9, 2, 2, 81, 82,
	7, 11, 2, 2, 82, 94, 5, 16, 9, 2, 83, 84, 12, 9, 2, 2, 84, 85, 7, 20, 2,
	2, 85, 94, 5, 42, 22, 2, 86, 87, 12, 8, 2, 2, 87, 88, 7, 10, 2, 2, 88,
	94, 5, 44, 23, 2, 89, 90, 12, 7, 2, 2, 90, 91, 7, 10, 2, 2, 91, 92, 7,
	9, 2, 2, 92, 94, 5, 44, 23, 2, 93, 68, 3, 2, 2, 2, 93, 72, 3, 2, 2, 2,
	93, 76, 3, 2, 2, 2, 93, 79, 3, 2, 2, 2, 93, 83, 3, 2, 2, 2, 93, 86, 3,
	2, 2, 2, 93, 89, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95,
	96, 3, 2, 2, 2, 96, 5, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 102, 7, 26,
	2, 2, 99, 102, 5, 14, 8, 2, 100, 102, 5, 18, 10, 2, 101, 98, 3, 2, 2, 2,
	101, 99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 7, 3, 2, 2, 2, 103, 104,
	9, 2, 2, 2, 104, 9, 3, 2, 2, 2, 105, 106, 9, 3, 2, 2, 106, 11, 3, 2, 2,
	2, 107, 108, 9, 4, 2, 2, 108, 13, 3, 2, 2, 2, 109, 110, 9, 5, 2, 2, 110,
	15, 3, 2, 2, 2, 111, 115, 5, 18, 10, 2, 112, 115, 5, 6, 4, 2, 113, 115,
	7, 6, 2, 2, 114, 111, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 113, 3, 2,
	2, 2, 115, 17, 3, 2, 2, 2, 116, 118, 7, 43, 2, 2, 117, 119, 5, 48, 25,
	2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120,
	137, 7, 44, 2, 2, 121, 123, 7, 43, 2, 2, 122, 124, 5, 46, 24, 2, 123, 122,
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 137, 7, 44,
	2, 2, 126, 128, 7, 43, 2, 2, 127, 129, 5, 50, 26, 2, 128, 127, 3, 2, 2,
	2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 137, 7, 44, 2, 2, 131,
	133, 7, 43, 2, 2, 132, 134, 5, 52, 27, 2, 133, 132, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 7, 44, 2, 2, 136, 116, 3, 2,
	2, 2, 136, 121, 3, 2, 2, 2, 136, 126, 3, 2, 2, 2, 136, 131, 3, 2, 2, 2,
	137, 19, 3, 2, 2, 2, 138, 139, 5, 22, 12, 2, 139, 21, 3, 2, 2, 2, 140,
	141, 8, 12, 1, 2, 141, 142, 5, 24, 13, 2, 142, 154, 3, 2, 2, 2, 143, 144,
	12, 6, 2, 2, 144, 145, 7, 37, 2, 2, 145, 153, 5, 24, 13, 2, 146, 147, 12,
	5, 2, 2, 147, 148, 7, 38, 2, 2, 148, 153, 5, 24, 13, 2, 149, 150, 12, 4,
	2, 2, 150, 151, 7, 39, 2, 2, 151, 153, 5, 24, 13, 2, 152, 143, 3, 2, 2,
	2, 152, 146, 3, 2, 2, 2, 152, 149, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154,
	152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 23, 3, 2, 2, 2, 156, 154, 3,
	2, 2, 2, 157, 158, 8, 13, 1, 2, 158, 159, 5, 26, 14, 2, 159, 171, 3, 2,
	2, 2, 160, 161, 12, 6, 2, 2, 161, 162, 7, 34, 2, 2, 162, 170, 5, 26, 14,
	2, 163, 164, 12, 5, 2, 2, 164, 165, 7, 35, 2, 2, 165, 170, 5, 26, 14, 2,
	166, 167, 12, 4, 2, 2, 167, 168, 7, 36, 2, 2, 168, 170, 5, 26, 14, 2, 169,
	160, 3, 2, 2, 2, 169, 163, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2, 170, 173,
	3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 25, 3, 2,
	2, 2, 173, 171, 3, 2, 2, 2, 174, 175, 8, 14, 1, 2, 175, 176, 5, 28, 15,
	2, 176, 185, 3, 2, 2, 2, 177, 178, 12, 5, 2, 2, 178, 179, 7, 27, 2, 2,
	179, 184, 5, 28, 15, 2, 180, 181, 12, 4, 2, 2, 181, 182, 7, 28, 2, 2, 182,
	184, 5, 28, 15, 2, 183, 177, 3, 2, 2, 2, 183, 180, 3, 2, 2, 2, 184, 187,
	3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 27, 3, 2,
	2, 2, 187, 185, 3, 2, 2, 2, 188, 189, 8, 15, 1, 2, 189, 190, 5, 30, 16,
	2, 190, 202, 3, 2, 2, 2, 191, 192, 12, 6, 2, 2, 192, 193, 7, 29, 2, 2,
	193, 201, 5, 30, 16, 2, 194, 195, 12, 5, 2, 2, 195, 196, 7, 30, 2, 2, 196,
	201, 5, 30, 16, 2, 197, 198, 12, 4, 2, 2, 198, 199, 7, 31, 2, 2, 199, 201,
	5, 30, 16, 2, 200, 191, 3, 2, 2, 2, 200, 194, 3, 2, 2, 2, 200, 197, 3,
	2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2,
	2, 203, 29, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 213, 5, 6, 4, 2, 206,
	213, 5, 32, 17, 2, 207, 208, 7, 21, 2, 2, 208, 209, 5, 22, 12, 2, 209,
	210, 7, 22, 2, 2, 210, 213, 3, 2, 2, 2, 211, 213, 5, 34, 18, 2, 212, 205,
	3, 2, 2, 2, 212, 206, 3, 2, 2, 2, 212, 207, 3, 2, 2, 2, 212, 211, 3, 2,
	2, 2, 213, 31, 3, 2, 2, 2, 214, 215, 5, 40, 21, 2, 215, 216, 7, 33, 2,
	2, 216, 217, 5, 40, 21, 2, 217, 220, 3, 2, 2, 2, 218, 220, 5, 40, 21, 2,
	219, 214, 3, 2, 2, 2, 219, 218, 3, 2, 2, 2, 220, 33, 3, 2, 2, 2, 221, 222,
	5, 36, 19, 2, 222, 223, 7, 21, 2, 2, 223, 224, 5, 38, 20, 2, 224, 225,
	7, 22, 2, 2, 225, 35, 3, 2, 2, 2, 226, 227, 7, 25, 2, 2, 227, 37, 3, 2,
	2, 2, 228, 233, 5, 4, 3, 2, 229, 230, 7, 42, 2, 2, 230, 232, 5, 4, 3, 2,
	231, 229, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233,
	234, 3, 2, 2, 2, 234, 39, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236, 238, 7,
	28, 2, 2, 237, 236, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 240, 3, 2, 2,
	2, 239, 241, 7, 41, 2, 2, 240, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242,
	240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 250, 3, 2, 2, 2, 244, 246,
	7, 32, 2, 2, 245, 247, 7, 41, 2, 2, 246, 245, 3, 2, 2, 2, 247, 248, 3,
	2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 251, 3, 2, 2,
	2, 250, 244, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 41, 3, 2, 2, 2, 252,
	253, 7, 47, 2, 2, 253, 43, 3, 2, 2, 2, 254, 255, 7, 25, 2, 2, 255, 45,
	3, 2, 2, 2, 256, 261, 7, 6, 2, 2, 257, 258, 7, 42, 2, 2, 258, 260, 7, 6,
	2, 2, 259, 257, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2,
	261, 262, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264,
	266, 7, 42, 2, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 47,
	3, 2, 2, 2, 267, 272, 7, 3, 2, 2, 268, 269, 7, 42, 2, 2, 269, 271, 7, 3,
	2, 2, 270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2,
	272, 273, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275,
	277, 7, 42, 2, 2, 276, 275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 49,
	3, 2, 2, 2, 278, 283, 7, 4, 2, 2, 279, 280, 7, 42, 2, 2, 280, 282, 7, 4,
	2, 2, 281, 279, 3, 2, 2, 2, 282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2,
	283, 284, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286,
	288, 7, 42, 2, 2, 287, 286, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 51,
	3, 2, 2, 2, 289, 294, 7, 5, 2, 2, 290, 291, 7, 42, 2, 2, 291, 293, 7, 5,
	2, 2, 292, 290, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2,
	294, 295, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297,
	299, 7, 42, 2, 2, 298, 297, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 53,
	3, 2, 2, 2, 35, 66, 93, 95, 101, 114, 118, 123, 128, 133, 136, 152, 154,
	169, 171, 183, 185, 200, 202, 212, 219, 233, 237, 242, 248, 250, 261, 265,
	272, 276, 283, 287, 294, 298,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'&&'", "'||'", "'not'", "'is'", "'in'", "'true'",
	"'false'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='", "'=~'", "'('",
	"')'", "", "", "", "", "'+'", "'-'", "'*'", "", "'%'", "'.'", "", "'<<'",
	"'>>'", "'>>>'", "'&'", "'^'", "'|'", "'\n'", "", "','", "'['", "']'",
}
var symbolicNames = []string{
	"", "Integer", "Float", "Boolean", "TString", "AND", "OR", "NOT", "IS",
	"IN", "TRUE", "FALSE", "GT", "GE", "LT", "LE", "EQ", "NE", "MATCH", "LPAREN",
	"RPAREN", "INT", "FLOAT", "IDENTIFIER", "VARIABLE", "PLUS", "MINUS", "MUL",
	"DIV", "MOD", "POINT", "E", "LSHIFT", "RSHIFT", "RSHIFT3", "BAND", "BEOR",
	"BIOR", "NL", "DIGIT", "COMMA", "LBRACKET", "RBRACKET", "WS", "SLASH",
	"REGEX",
}

var ruleNames = []string{
	"parse", "expression", "variable", "comparator", "binary", "boolean", "literal",
	"container", "array", "calc", "bit", "shift", "plus", "multiplying", "atom",
	"scientific", "function", "funcname", "parameters", "number", "regex",
	"kind", "strings", "integers", "floats", "booleans",
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
	this.GrammarFileName = "TExprParser.g4"

	return this
}

// TExprParser tokens.
const (
	TExprParserEOF        = antlr.TokenEOF
	TExprParserInteger    = 1
	TExprParserFloat      = 2
	TExprParserBoolean    = 3
	TExprParserTString    = 4
	TExprParserAND        = 5
	TExprParserOR         = 6
	TExprParserNOT        = 7
	TExprParserIS         = 8
	TExprParserIN         = 9
	TExprParserTRUE       = 10
	TExprParserFALSE      = 11
	TExprParserGT         = 12
	TExprParserGE         = 13
	TExprParserLT         = 14
	TExprParserLE         = 15
	TExprParserEQ         = 16
	TExprParserNE         = 17
	TExprParserMATCH      = 18
	TExprParserLPAREN     = 19
	TExprParserRPAREN     = 20
	TExprParserINT        = 21
	TExprParserFLOAT      = 22
	TExprParserIDENTIFIER = 23
	TExprParserVARIABLE   = 24
	TExprParserPLUS       = 25
	TExprParserMINUS      = 26
	TExprParserMUL        = 27
	TExprParserDIV        = 28
	TExprParserMOD        = 29
	TExprParserPOINT      = 30
	TExprParserE          = 31
	TExprParserLSHIFT     = 32
	TExprParserRSHIFT     = 33
	TExprParserRSHIFT3    = 34
	TExprParserBAND       = 35
	TExprParserBEOR       = 36
	TExprParserBIOR       = 37
	TExprParserNL         = 38
	TExprParserDIGIT      = 39
	TExprParserCOMMA      = 40
	TExprParserLBRACKET   = 41
	TExprParserRBRACKET   = 42
	TExprParserWS         = 43
	TExprParserSLASH      = 44
	TExprParserREGEX      = 45
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
	TExprParserRULE_container   = 7
	TExprParserRULE_array       = 8
	TExprParserRULE_calc        = 9
	TExprParserRULE_bit         = 10
	TExprParserRULE_shift       = 11
	TExprParserRULE_plus        = 12
	TExprParserRULE_multiplying = 13
	TExprParserRULE_atom        = 14
	TExprParserRULE_scientific  = 15
	TExprParserRULE_function    = 16
	TExprParserRULE_funcname    = 17
	TExprParserRULE_parameters  = 18
	TExprParserRULE_number      = 19
	TExprParserRULE_regex       = 20
	TExprParserRULE_kind        = 21
	TExprParserRULE_strings     = 22
	TExprParserRULE_integers    = 23
	TExprParserRULE_floats      = 24
	TExprParserRULE_booleans    = 25
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
	case TExprParserVisitor:
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
		p.SetState(52)
		p.expression(0)
	}
	{
		p.SetState(53)
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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

func (s *InExpressionContext) Container() IContainerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainerContext)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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

func (s *NotInExpressionContext) Container() IContainerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainerContext)
}

func (s *NotInExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	case TExprParserVisitor:
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(56)
			p.Match(TExprParserLPAREN)
		}
		{
			p.SetState(57)
			p.expression(0)
		}
		{
			p.SetState(58)
			p.Match(TExprParserRPAREN)
		}

	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(TExprParserNOT)
		}
		{
			p.SetState(61)
			p.expression(10)
		}

	case 3:
		localctx = NewVariableExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Variable()
		}

	case 4:
		localctx = NewCalcExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Calc()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(67)
					p.Comparator()
				}
				{
					p.SetState(68)
					p.expression(5)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(71)
					p.Binary()
				}
				{
					p.SetState(72)
					p.expression(4)
				}

			case 3:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(75)
					p.Match(TExprParserIN)
				}
				{
					p.SetState(76)
					p.Container()
				}

			case 4:
				localctx = NewNotInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(78)
					p.Match(TExprParserNOT)
				}
				{
					p.SetState(79)
					p.Match(TExprParserIN)
				}
				{
					p.SetState(80)
					p.Container()
				}

			case 5:
				localctx = NewMatchExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(82)
					p.Match(TExprParserMATCH)
				}
				{
					p.SetState(83)
					p.Regex()
				}

			case 6:
				localctx = NewIsTypeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(85)
					p.Match(TExprParserIS)
				}
				{
					p.SetState(86)
					p.Kind()
				}

			case 7:
				localctx = NewIsNotTypeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(88)
					p.Match(TExprParserIS)
				}
				{
					p.SetState(89)
					p.Match(TExprParserNOT)
				}
				{
					p.SetState(90)
					p.Kind()
				}

			}

		}
		p.SetState(95)
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
	case TExprParserVisitor:
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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TExprParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(TExprParserVARIABLE)
		}

	case TExprParserInteger, TExprParserFloat, TExprParserBoolean, TExprParserTString:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Literal()
		}

	case TExprParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
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
	case TExprParserVisitor:
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
		p.SetState(101)
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
	case TExprParserVisitor:
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
		p.SetState(103)
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
	case TExprParserVisitor:
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
		p.SetState(105)
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

func (s *LiteralContext) TString() antlr.TerminalNode {
	return s.GetToken(TExprParserTString, 0)
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
	case TExprParserVisitor:
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
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TExprParserInteger)|(1<<TExprParserFloat)|(1<<TExprParserBoolean)|(1<<TExprParserTString))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IContainerContext is an interface to support dynamic dispatch.
type IContainerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainerContext differentiates from other interfaces.
	IsContainerContext()
}

type ContainerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainerContext() *ContainerContext {
	var p = new(ContainerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_container
	return p
}

func (*ContainerContext) IsContainerContext() {}

func NewContainerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainerContext {
	var p = new(ContainerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_container

	return p
}

func (s *ContainerContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainerContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ContainerContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ContainerContext) TString() antlr.TerminalNode {
	return s.GetToken(TExprParserTString, 0)
}

func (s *ContainerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitContainer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Container() (localctx IContainerContext) {
	localctx = NewContainerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TExprParserRULE_container)

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Array()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Match(TExprParserTString)
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

func (s *ArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(TExprParserLBRACKET, 0)
}

func (s *ArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(TExprParserRBRACKET, 0)
}

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
	case TExprParserVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TExprParserRULE_array)
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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(TExprParserLBRACKET)
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserInteger {
			{
				p.SetState(115)
				p.Integers()
			}

		}
		{
			p.SetState(118)
			p.Match(TExprParserRBRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Match(TExprParserLBRACKET)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserTString {
			{
				p.SetState(120)
				p.Strings()
			}

		}
		{
			p.SetState(123)
			p.Match(TExprParserRBRACKET)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(TExprParserLBRACKET)
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserFloat {
			{
				p.SetState(125)
				p.Floats()
			}

		}
		{
			p.SetState(128)
			p.Match(TExprParserRBRACKET)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.Match(TExprParserLBRACKET)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TExprParserBoolean {
			{
				p.SetState(130)
				p.Booleans()
			}

		}
		{
			p.SetState(133)
			p.Match(TExprParserRBRACKET)
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

func (s *CalcContext) Bit() IBitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitContext)
}

func (s *CalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitCalc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Calc() (localctx ICalcContext) {
	localctx = NewCalcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TExprParserRULE_calc)

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
		p.SetState(136)
		p.bit(0)
	}

	return localctx
}

// IBitContext is an interface to support dynamic dispatch.
type IBitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitContext differentiates from other interfaces.
	IsBitContext()
}

type BitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitContext() *BitContext {
	var p = new(BitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_bit
	return p
}

func (*BitContext) IsBitContext() {}

func NewBitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitContext {
	var p = new(BitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_bit

	return p
}

func (s *BitContext) GetParser() antlr.Parser { return s.parser }

func (s *BitContext) Shift() IShiftContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftContext)
}

func (s *BitContext) Bit() IBitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitContext)
}

func (s *BitContext) BAND() antlr.TerminalNode {
	return s.GetToken(TExprParserBAND, 0)
}

func (s *BitContext) BEOR() antlr.TerminalNode {
	return s.GetToken(TExprParserBEOR, 0)
}

func (s *BitContext) BIOR() antlr.TerminalNode {
	return s.GetToken(TExprParserBIOR, 0)
}

func (s *BitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitBit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Bit() (localctx IBitContext) {
	return p.bit(0)
}

func (p *TExprParser) bit(_p int) (localctx IBitContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBitContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBitContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, TExprParserRULE_bit, _p)

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
		p.SetState(139)
		p.shift(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_bit)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(142)
					p.Match(TExprParserBAND)
				}
				{
					p.SetState(143)
					p.shift(0)
				}

			case 2:
				localctx = NewBitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_bit)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(145)
					p.Match(TExprParserBEOR)
				}
				{
					p.SetState(146)
					p.shift(0)
				}

			case 3:
				localctx = NewBitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_bit)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(148)
					p.Match(TExprParserBIOR)
				}
				{
					p.SetState(149)
					p.shift(0)
				}

			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IShiftContext is an interface to support dynamic dispatch.
type IShiftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShiftContext differentiates from other interfaces.
	IsShiftContext()
}

type ShiftContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftContext() *ShiftContext {
	var p = new(ShiftContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TExprParserRULE_shift
	return p
}

func (*ShiftContext) IsShiftContext() {}

func NewShiftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftContext {
	var p = new(ShiftContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TExprParserRULE_shift

	return p
}

func (s *ShiftContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftContext) Plus() IPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusContext)
}

func (s *ShiftContext) Shift() IShiftContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftContext)
}

func (s *ShiftContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(TExprParserLSHIFT, 0)
}

func (s *ShiftContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(TExprParserRSHIFT, 0)
}

func (s *ShiftContext) RSHIFT3() antlr.TerminalNode {
	return s.GetToken(TExprParserRSHIFT3, 0)
}

func (s *ShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitShift(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Shift() (localctx IShiftContext) {
	return p.shift(0)
}

func (p *TExprParser) shift(_p int) (localctx IShiftContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewShiftContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, TExprParserRULE_shift, _p)

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
		p.SetState(156)
		p.plus(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewShiftContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_shift)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(159)
					p.Match(TExprParserLSHIFT)
				}
				{
					p.SetState(160)
					p.plus(0)
				}

			case 2:
				localctx = NewShiftContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_shift)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(162)
					p.Match(TExprParserRSHIFT)
				}
				{
					p.SetState(163)
					p.plus(0)
				}

			case 3:
				localctx = NewShiftContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_shift)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(165)
					p.Match(TExprParserRSHIFT3)
				}
				{
					p.SetState(166)
					p.plus(0)
				}

			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	case TExprParserVisitor:
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, TExprParserRULE_plus, _p)

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
		p.SetState(173)
		p.multiplying(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPlusContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_plus)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(176)
					p.Match(TExprParserPLUS)
				}
				{
					p.SetState(177)
					p.multiplying(0)
				}

			case 2:
				localctx = NewPlusContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_plus)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(179)
					p.Match(TExprParserMINUS)
				}
				{
					p.SetState(180)
					p.multiplying(0)
				}

			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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

func (s *MultiplyingContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
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
	case TExprParserVisitor:
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, TExprParserRULE_multiplying, _p)

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
		p.SetState(187)
		p.Atom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(190)
					p.Match(TExprParserMUL)
				}
				{
					p.SetState(191)
					p.Atom()
				}

			case 2:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(193)
					p.Match(TExprParserDIV)
				}
				{
					p.SetState(194)
					p.Atom()
				}

			case 3:
				localctx = NewMultiplyingContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TExprParserRULE_multiplying)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(196)
					p.Match(TExprParserMOD)
				}
				{
					p.SetState(197)
					p.Atom()
				}

			}

		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
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

func (s *AtomContext) Bit() IBitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitContext)
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
	case TExprParserVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TExprParserRULE_atom)

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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TExprParserInteger, TExprParserFloat, TExprParserBoolean, TExprParserTString, TExprParserVARIABLE, TExprParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Variable()
		}

	case TExprParserMINUS, TExprParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Scientific()
		}

	case TExprParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Match(TExprParserLPAREN)
		}
		{
			p.SetState(206)
			p.bit(0)
		}
		{
			p.SetState(207)
			p.Match(TExprParserRPAREN)
		}

	case TExprParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)
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
	case TExprParserVisitor:
		return t.VisitScientific(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Scientific() (localctx IScientificContext) {
	localctx = NewScientificContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TExprParserRULE_scientific)

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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Number()
		}
		{
			p.SetState(213)
			p.Match(TExprParserE)
		}
		{
			p.SetState(214)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
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
	case TExprParserVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TExprParserRULE_function)

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
		p.SetState(219)
		p.Funcname()
	}
	{
		p.SetState(220)
		p.Match(TExprParserLPAREN)
	}
	{
		p.SetState(221)
		p.Parameters()
	}
	{
		p.SetState(222)
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
	case TExprParserVisitor:
		return t.VisitFuncname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TExprParserRULE_funcname)

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
		p.SetState(224)
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

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TExprParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TExprParserRULE_parameters)
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
		p.SetState(226)
		p.expression(0)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TExprParserCOMMA {
		{
			p.SetState(227)
			p.Match(TExprParserCOMMA)
		}
		{
			p.SetState(228)
			p.expression(0)
		}

		p.SetState(233)
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
	case TExprParserVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TExprParserRULE_number)
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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserMINUS {
		{
			p.SetState(234)
			p.Match(TExprParserMINUS)
		}

	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(237)
				p.Match(TExprParserDIGIT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(242)
			p.Match(TExprParserPOINT)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(243)
					p.Match(TExprParserDIGIT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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

func (s *RegexContext) REGEX() antlr.TerminalNode {
	return s.GetToken(TExprParserREGEX, 0)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitRegex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Regex() (localctx IRegexContext) {
	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TExprParserRULE_regex)

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
		p.SetState(250)
		p.Match(TExprParserREGEX)
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
	case TExprParserVisitor:
		return t.VisitKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Kind() (localctx IKindContext) {
	localctx = NewKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TExprParserRULE_kind)

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
		p.SetState(252)
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

func (s *StringsContext) AllTString() []antlr.TerminalNode {
	return s.GetTokens(TExprParserTString)
}

func (s *StringsContext) TString(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserTString, i)
}

func (s *StringsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TExprParserCOMMA)
}

func (s *StringsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserCOMMA, i)
}

func (s *StringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Strings() (localctx IStringsContext) {
	localctx = NewStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TExprParserRULE_strings)
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
		p.SetState(254)
		p.Match(TExprParserTString)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(255)
				p.Match(TExprParserCOMMA)
			}
			{
				p.SetState(256)
				p.Match(TExprParserTString)
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserCOMMA {
		{
			p.SetState(262)
			p.Match(TExprParserCOMMA)
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

func (s *IntegersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TExprParserCOMMA)
}

func (s *IntegersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserCOMMA, i)
}

func (s *IntegersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitIntegers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Integers() (localctx IIntegersContext) {
	localctx = NewIntegersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TExprParserRULE_integers)
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
		p.SetState(265)
		p.Match(TExprParserInteger)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(266)
				p.Match(TExprParserCOMMA)
			}
			{
				p.SetState(267)
				p.Match(TExprParserInteger)
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserCOMMA {
		{
			p.SetState(273)
			p.Match(TExprParserCOMMA)
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

func (s *FloatsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TExprParserCOMMA)
}

func (s *FloatsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserCOMMA, i)
}

func (s *FloatsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitFloats(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Floats() (localctx IFloatsContext) {
	localctx = NewFloatsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TExprParserRULE_floats)
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
		p.SetState(276)
		p.Match(TExprParserFloat)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(277)
				p.Match(TExprParserCOMMA)
			}
			{
				p.SetState(278)
				p.Match(TExprParserFloat)
			}

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserCOMMA {
		{
			p.SetState(284)
			p.Match(TExprParserCOMMA)
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

func (s *BooleansContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TExprParserCOMMA)
}

func (s *BooleansContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TExprParserCOMMA, i)
}

func (s *BooleansContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleansContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleansContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TExprParserVisitor:
		return t.VisitBooleans(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TExprParser) Booleans() (localctx IBooleansContext) {
	localctx = NewBooleansContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TExprParserRULE_booleans)
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
		p.SetState(287)
		p.Match(TExprParserBoolean)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(288)
				p.Match(TExprParserCOMMA)
			}
			{
				p.SetState(289)
				p.Match(TExprParserBoolean)
			}

		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TExprParserCOMMA {
		{
			p.SetState(295)
			p.Match(TExprParserCOMMA)
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

	case 10:
		var t *BitContext = nil
		if localctx != nil {
			t = localctx.(*BitContext)
		}
		return p.Bit_Sempred(t, predIndex)

	case 11:
		var t *ShiftContext = nil
		if localctx != nil {
			t = localctx.(*ShiftContext)
		}
		return p.Shift_Sempred(t, predIndex)

	case 12:
		var t *PlusContext = nil
		if localctx != nil {
			t = localctx.(*PlusContext)
		}
		return p.Plus_Sempred(t, predIndex)

	case 13:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Bit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Shift_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Plus_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TExprParser) Multiplying_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
