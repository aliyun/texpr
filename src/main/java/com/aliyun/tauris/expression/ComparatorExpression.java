package com.aliyun.tauris.expression;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class ComparatorExpression extends TExpression {

    private TExpression left;
    private TExpression right;

    private ComparatorOp op;

    public ComparatorExpression(TExpression left, ComparatorOp op, TExpression right) {
        this.left = left;
        this.op = op;
        this.right = right;
    }

    @Override
    public Object eval(Context e) {
        return op.eval(e, left, right);
    }

    @Override
    public String toString() {
        return String.format("%s %s %s", left, op, right);
    }
}
