package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public class NotExpression extends TExpression {

    private TExpression right;

    public NotExpression(TExpression right) {
        if (right == null) throw new IllegalArgumentException("null value");
        this.right = right;
    }

    @Override
    public Comparable eval(Context e) {
        Object ret = right.eval(e);
        if (!(ret instanceof Boolean)) {
            throw new ExprException(String.format("`%s` is not a boolean expression", right.toString()));
        }
        return !((Boolean)ret);
    }

    @Override
    public String toString() {
        return "not " + right.toString();
    }
}
