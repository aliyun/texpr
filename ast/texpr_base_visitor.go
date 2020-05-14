// Code generated from TExpr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package ast // TExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTExprVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitMatchExpression(ctx *MatchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitIsTypeExpression(ctx *IsTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitCalcExpression(ctx *CalcExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitIsNotTypeExpression(ctx *IsNotTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitNotInExpression(ctx *NotInExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitComparatorExpression(ctx *ComparatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitCalc(ctx *CalcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitPlus(ctx *PlusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitMultiplying(ctx *MultiplyingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitPow(ctx *PowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitFuncname(ctx *FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitRegex(ctx *RegexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitKind(ctx *KindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitStrings(ctx *StringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitIntegers(ctx *IntegersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitFloats(ctx *FloatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTExprVisitor) VisitBooleans(ctx *BooleansContext) interface{} {
	return v.VisitChildren(ctx)
}
