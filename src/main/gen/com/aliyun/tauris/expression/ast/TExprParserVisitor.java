// Generated from TExprParser.g4 by ANTLR 4.6

    package com.aliyun.tauris.expression.ast;

import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link TExprParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface TExprParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link TExprParser#parse}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParse(TExprParser.ParseContext ctx);
	/**
	 * Visit a parse tree produced by the {@code binaryExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryExpression(TExprParser.BinaryExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code matchExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatchExpression(TExprParser.MatchExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code inExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInExpression(TExprParser.InExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isTypeExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsTypeExpression(TExprParser.IsTypeExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code calcExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCalcExpression(TExprParser.CalcExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isNotTypeExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsNotTypeExpression(TExprParser.IsNotTypeExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code notExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNotExpression(TExprParser.NotExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code parenExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParenExpression(TExprParser.ParenExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code notInExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNotInExpression(TExprParser.NotInExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code comparatorExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitComparatorExpression(TExprParser.ComparatorExpressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code variableExpression}
	 * labeled alternative in {@link TExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVariableExpression(TExprParser.VariableExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#variable}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVariable(TExprParser.VariableContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#comparator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitComparator(TExprParser.ComparatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#binary}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinary(TExprParser.BinaryContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#bool}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBool(TExprParser.BoolContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteral(TExprParser.LiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#container}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitContainer(TExprParser.ContainerContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#array}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray(TExprParser.ArrayContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#calc}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCalc(TExprParser.CalcContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#bit}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBit(TExprParser.BitContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#shift}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitShift(TExprParser.ShiftContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#plus}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPlus(TExprParser.PlusContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#multiplying}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMultiplying(TExprParser.MultiplyingContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAtom(TExprParser.AtomContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#scientific}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitScientific(TExprParser.ScientificContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#function}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction(TExprParser.FunctionContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#funcname}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncname(TExprParser.FuncnameContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#parameters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParameters(TExprParser.ParametersContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#number}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumber(TExprParser.NumberContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#regex}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRegex(TExprParser.RegexContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType(TExprParser.TypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#strings}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStrings(TExprParser.StringsContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#integers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegers(TExprParser.IntegersContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#floats}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFloats(TExprParser.FloatsContext ctx);
	/**
	 * Visit a parse tree produced by {@link TExprParser#booleans}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBooleans(TExprParser.BooleansContext ctx);
}