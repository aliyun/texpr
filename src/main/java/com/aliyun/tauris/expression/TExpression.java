package com.aliyun.tauris.expression;


import com.aliyun.tauris.expression.ast.TExprLexer;
import com.aliyun.tauris.expression.ast.TExprParser;
import com.aliyun.tauris.expression.ast.TExprParserBaseVisitor;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.ParseTree;
import org.antlr.v4.runtime.tree.TerminalNode;

import java.util.*;
import java.util.function.BiFunction;
import java.util.regex.Pattern;

/**
 * Created by yundun-waf-dev
 */
public abstract class TExpression {

    private static IsTypeFactory        IS_TYPE_FACTORY    = new IsTypeFactory();
    private static EmbedFunctionFactory EMBED_FUNC_FACTORY = new EmbedFunctionFactory();

    protected String expression;

    protected Set<String> variables;

    public abstract Object eval(Context e);

    public Number calc(Context e) {
        Object ret = this.eval(e);
        if (ret instanceof Number) {
            return ((Number) ret);
        }
        if (ret == null) {
            throw new ExprException(String.format("`%s` is null, not a number", this));
        }
        throw new ExprException(String.format("`%s` returns %s which is not number type", this, ret.getClass().getSimpleName()));
    }

    public Boolean check(Context e) {
        Object ret = this.eval(e);
        if (ret == null) {
            return false;
        }
        if (ret instanceof Boolean) {
            return ((Boolean) ret);
        }
        throw new ExprException(String.format("`%s` returns %s which is not boolean type", this, ret.getClass().getSimpleName()));
    }

    /**
     * get variable names within expression
     *
     * @return variable names set
     */
    public Set<String> getVariables() {
        return Collections.unmodifiableSet(variables);
    }

    public static TExpression compile(String expr) {
        CharStream  charStream = new ANTLRInputStream(expr);
        TExprLexer  lexer      = new TExprLexer(charStream);
        TokenStream tokens = new CommonTokenStream(lexer);
        TExprParser parser = new TExprParser(tokens);
        parser.removeErrorListeners();
        ExpressionErrorListener errorListener = new ExpressionErrorListener(expr);
        parser.addErrorListener(errorListener);
        ExpressionVisitor visitor = new ExpressionVisitor();
        try {
            TExpression e = visitor.visit(parser.parse());
            if (errorListener.hasError()) {
                throw new ExprException(String.format("expression `%s` syntax error, %s", expr, errorListener.errorMessage));
            }
            e.expression = expr;
            return e;
        } catch (ExprException e) {
            throw e;
        } catch (Exception e) {
            throw new ExprException(String.format("parse expression `%s` failed", expr), e);
        }
    }

    public static void throwIfNull(Object arg, String name) throws IllegalArgumentException {
        if (arg == null) {
            throw new IllegalArgumentException(name + " is null");
        }
    }

    @Override
    public String toString() {
        return expression;
    }

    private static class ExpressionVisitor extends TExprParserBaseVisitor<TExpression> {

        private Set<String> variables = new HashSet<>();

        @Override
        public TExpression visit(ParseTree tree) {
            TExpression expr = super.visit(tree);
            expr.variables = new HashSet<>(variables);
            return expr;
        }

        @Override
        public TExpression visitParse(TExprParser.ParseContext ctx) {
            return ctx.expression().accept(this);
        }

        @Override
        public TExpression visitBinaryExpression(TExprParser.BinaryExpressionContext ctx) {
            BinaryOp    op   = BinaryOp.build(ctx.binary().getText());
            TExpression left = super.visit(ctx.expression().get(0));
            TExpression right = super.visit(ctx.expression().get(1));
            return new BinaryExpression(left, op, right);
        }

        @Override
        public TExpression visitLiteral(TExprParser.LiteralContext ctx) {
            if (ctx.Boolean() != null) {
                return new LiteralExpression.BOOLEAN(Boolean.valueOf(ctx.getText()));
            }
            if (ctx.Float() != null) {
                return new LiteralExpression.FLOAT(Double.valueOf(ctx.getText()));
            }
            if (ctx.Integer() != null) {
                return new LiteralExpression.INT(Long.valueOf(ctx.getText()));
            }
            if (ctx.String() != null) {
                return new LiteralExpression.STRING(ctx.getText());
            }
            throw new IllegalArgumentException("unknown literal");
        }

        @Override
        public TExpression visitArray(TExprParser.ArrayContext ctx) {
            if (ctx.booleans() != null) {
                return ctx.booleans().accept(this);
            }
            if (ctx.integers() != null) {
                return ctx.integers().accept(this);
            }
            if (ctx.floats() != null) {
                return ctx.floats().accept(this);
            }
            if (ctx.strings() != null) {
                return ctx.strings().accept(this);
            }
            if (ctx.getText().replaceAll(" ", "").equals("[]")) {
                return new SetExpression(new HashSet<>());
            }
            throw new IllegalArgumentException("unknown array type");
        }

        @Override
        public TExpression visitStrings(TExprParser.StringsContext ctx) {
            Set<String> set = new HashSet<>();
            for (TerminalNode n : ctx.String()) {
                set.add(n.getText().substring(1, n.getText().length() - 1));
            }
            return new SetExpression(set);
        }

        @Override
        public TExpression visitIntegers(TExprParser.IntegersContext ctx) {
            Set<Long> set = new HashSet<>();
            for (TerminalNode n : ctx.Integer()) {
                set.add(Long.valueOf(n.getText()));
            }
            return new SetExpression(set);
        }

        @Override
        public TExpression visitFloats(TExprParser.FloatsContext ctx) {
            Set<Double> set = new HashSet<>();
            for (TerminalNode n : ctx.Float()) {
                set.add(Double.valueOf(n.getText()));
            }
            return new SetExpression(set);
        }

        @Override
        public TExpression visitBooleans(TExprParser.BooleansContext ctx) {
            Set<Boolean> set = new HashSet<>();
            for (TerminalNode n : ctx.Boolean()) {
                set.add(Boolean.valueOf(n.getText()));
            }
            return new SetExpression(set);
        }

        @Override
        public TExpression visitInExpression(TExprParser.InExpressionContext ctx) {
            TExpression left = this.visit(ctx.expression());
            return new InExpression(left, (ContainerExpression) ctx.container().accept(this));
        }

        @Override
        public TExpression visitNotInExpression(TExprParser.NotInExpressionContext ctx) {
            TExpression left = this.visit(ctx.expression());
            return new InExpression(left, (ContainerExpression) ctx.container().accept(this), true);
        }

        @Override
        public TExpression visitContainer(TExprParser.ContainerContext ctx) {
            if (ctx.array() != null) {
                return ctx.array().accept(this);
            }
            if (ctx.variable() != null) {
                return ctx.variable().accept(this);
            }
            if (ctx.String() != null) {
                return new LiteralExpression.STRING(ctx.String().getText());
            }
            return super.visitContainer(ctx);
        }

        @Override
        public TExpression visitIsTypeExpression(TExprParser.IsTypeExpressionContext ctx) {
            TExpression expression = this.visit(ctx.expression());
            IsType      type       = (IsType) IS_TYPE_FACTORY.newInstance(ctx.type().getText());
            if (type == null) {
                throw new IllegalArgumentException("Unknown type:" + ctx.type().getText());
            }
            return new IsTypeExpression(expression, type);
        }


        @Override
        public TExpression visitIsNotTypeExpression(TExprParser.IsNotTypeExpressionContext ctx) {
            TExpression expression = this.visit(ctx.expression());
            IsType      type       = (IsType) IS_TYPE_FACTORY.newInstance(ctx.type().getText());
            if (type == null) {
                throw new IllegalArgumentException("Unknown type:" + ctx.type().getText());
            }
            return new IsTypeExpression(expression, type, true);
        }

        @Override
        public TExpression visitNotExpression(TExprParser.NotExpressionContext ctx) {
            return new NotExpression(super.visit(ctx.expression()));
        }

        @Override
        public TExpression visitMatchExpression(TExprParser.MatchExpressionContext ctx) {
            String  regex = ctx.regex().getText().trim();
            Pattern p     = Pattern.compile(regex.substring(1, regex.length() - 1));
            return new MatchExpression(this.visit(ctx.expression()), p);
        }

        @Override
        public TExpression visitParenExpression(TExprParser.ParenExpressionContext ctx) {
            return new ParenExpression(ctx.expression().accept(this));
        }

        @Override
        public TExpression visitComparatorExpression(TExprParser.ComparatorExpressionContext ctx) {
            ComparatorOp op    = ComparatorOp.build(ctx.comparator().getText());
            TExpression  left  = ctx.expression().get(0).accept(this);
            TExpression  right = ctx.expression().get(1).accept(this);
            return new ComparatorExpression(left, op, right);
        }

        @Override
        public TExpression visitVariable(TExprParser.VariableContext ctx) {
            if (ctx.VARIABLE() != null) {
                variables.add(ctx.VARIABLE().toString());
                return new VariableExpression(ctx.VARIABLE().toString());
            }
            if (ctx.literal() != null) {
                return ctx.literal().accept(this);
            }
            if (ctx.array() != null) {
                return ctx.array().accept(this);
            }
            return super.visitVariable(ctx);
        }

        @Override
        public TExpression visitType(TExprParser.TypeContext ctx) {
            return super.visitType(ctx);
        }

        @Override
        public TExpression visitComparator(TExprParser.ComparatorContext ctx) {
            return super.visitComparator(ctx);
        }

        @Override
        public TExpression visitBinary(TExprParser.BinaryContext ctx) {
            return super.visitBinary(ctx);
        }

        @Override
        public TExpression visitCalcExpression(TExprParser.CalcExpressionContext ctx) {
            return super.visitCalcExpression(ctx);
        }

        @Override
        public TExpression visitVariableExpression(TExprParser.VariableExpressionContext ctx) {
            return super.visitVariableExpression(ctx);
        }

        @Override
        public TExpression visitBool(TExprParser.BoolContext ctx) {
            return super.visitBool(ctx);
        }

        @Override
        public TExpression visitCalc(TExprParser.CalcContext ctx) {
            return ctx.bit().accept(this);
        }

        @Override
        public TExpression visitBit(TExprParser.BitContext ctx) {
            TExpression left = null;
            if (ctx.bit() != null) {
                left = ctx.bit().accept(this);
            }
            TExpression right = ctx.shift().accept(this);
            if (ctx.BAND() != null) {
                return new CalcExpression(left, right, "&", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() & r.longValue();
                    }
                });
            }
            if (ctx.BEOR() != null) {
                return new CalcExpression(left, right, "^", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() ^ r.longValue();
                    }
                });
            }
            if (ctx.BIOR() != null) {
                return new CalcExpression(left, right, "|", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() | r.longValue();
                    }
                });
            }
            return right;
        }

        @Override
        public TExpression visitShift(TExprParser.ShiftContext ctx) {
            TExpression left = null;
            if (ctx.shift() != null) {
                left = ctx.shift().accept(this);
            }
            TExpression right = ctx.plus().accept(this);
            if (ctx.LSHIFT() != null) {
                return new CalcExpression(left, right, "<<", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() << r.intValue();
                    }
                });
            }
            if (ctx.RSHIFT() != null) {
                return new CalcExpression(left, right, ">>", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() >> r.intValue();
                    }
                });
            }
            if (ctx.RSHIFT3() != null) {
                return new CalcExpression(left, right, ">>>", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.longValue() >>> r.intValue();
                    }
                });
            }
            return right;
        }

        @Override
        public TExpression visitPlus(TExprParser.PlusContext ctx) {
            TExpression left = null;
            if (ctx.plus() != null) {
                left = ctx.plus().accept(this);
            }
            TExpression right = ctx.multiplying().accept(this);
            if (ctx.MINUS() != null) {
                return new CalcExpression(left, right, "-", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.doubleValue() - r.doubleValue();
                    }
                });
            } else if (ctx.PLUS() != null) {
                return new CalcExpression(left, right, "+", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.doubleValue() + r.doubleValue();
                    }
                });
            }
            return ctx.multiplying().accept(this);
        }

        @Override
        public TExpression visitMultiplying(TExprParser.MultiplyingContext ctx) {
            TExpression left = null;
            if (ctx.multiplying() != null) {
                left = ctx.multiplying().accept(this);
            }
            TExpression right = ctx.atom().accept(this);
            if (ctx.DIV() != null) {
                return new CalcExpression(left, right, "/", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.doubleValue() / r.doubleValue();
                    }
                });
            } else if (ctx.MUL() != null) {
                return new CalcExpression(left, right, "*", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.doubleValue() * r.doubleValue();
                    }
                });
            } else if (ctx.MOD() != null) {
                return new CalcExpression(left, right, "%", new BiFunction<Number, Number, Number>() {
                    @Override
                    public Number apply(Number l, Number r) {
                        return l.doubleValue() % r.longValue();
                    }
                });
            }
            return right;
        }

        @Override
        public TExpression visitAtom(TExprParser.AtomContext ctx) {
            if (ctx.LPAREN() != null && ctx.RPAREN() != null) {
                return ctx.bit().accept(this);
            }
            if (ctx.variable() != null) {
                return ctx.variable().accept(this);
            }
            if (ctx.scientific() != null) {
                return ctx.scientific().accept(this);
            }
            if (ctx.function() != null) {
                return ctx.function().accept(this);
            }
            return this.visitChildren(ctx);
        }

        @Override
        public TExpression visitScientific(TExprParser.ScientificContext ctx) {
            return super.visitScientific(ctx);
        }

        @Override
        public TExpression visitFunction(TExprParser.FunctionContext ctx) {
            List<TExpression> params = new ArrayList<>(ctx.parameters().expression().size());
            for (TExprParser.ExpressionContext expr : ctx.parameters().expression()) {
                params.add(expr.accept(this));
            }
            EmbedFunction func = (EmbedFunction) EMBED_FUNC_FACTORY.newInstance(ctx.funcname().getText());
            return new FunctionExpresion(func, params);
        }

        @Override
        public TExpression visitFuncname(TExprParser.FuncnameContext ctx) {
            return super.visitFuncname(ctx);
        }

        @Override
        public TExpression visitParameters(TExprParser.ParametersContext ctx) {
            return super.visitParameters(ctx);
        }

        @Override
        public TExpression visitNumber(TExprParser.NumberContext ctx) {
            return super.visitNumber(ctx);
        }

        @Override
        public TExpression visitRegex(TExprParser.RegexContext ctx) {
            return super.visitRegex(ctx);
        }
    }

    private static class ExpressionErrorListener extends BaseErrorListener {

        String expression;

        String errorMessage;

        public ExpressionErrorListener(String expression) {
            this.expression = expression;
        }

        public boolean hasError() {
            return errorMessage != null;
        }

        public String getErrorMessage() {
            return errorMessage;
        }

        @Override
        public void syntaxError(Recognizer<?, ?> recognizer, Object offendingSymbol, int line, int charPositionInLine, String msg, RecognitionException e) {
            errorMessage = "line " + line + ":" + charPositionInLine + " " + msg;
        }
    }
}
