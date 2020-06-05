// Code generated from TExprParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package ast // TExprParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTExprParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTExprParserVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitMatchExpression(ctx *MatchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitIsTypeExpression(ctx *IsTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitCalcExpression(ctx *CalcExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitIsNotTypeExpression(ctx *IsNotTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitNotInExpression(ctx *NotInExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitComparatorExpression(ctx *ComparatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitContainer(ctx *ContainerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitCalc(ctx *CalcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitBit(ctx *BitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitShift(ctx *ShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitPlus(ctx *PlusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitMultiplying(ctx *MultiplyingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitFuncname(ctx *FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitRegex(ctx *RegexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitKind(ctx *KindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitStrings(ctx *StringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitIntegers(ctx *IntegersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitFloats(ctx *FloatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprParserVisitor) VisitBooleans(ctx *BooleansContext) interface{} {
	return v.VisitChildren(ctx)
}
