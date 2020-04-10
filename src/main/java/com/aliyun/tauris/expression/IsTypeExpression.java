package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsTypeExpression extends TExpression {

    private TExpression left;
    private IsType      type;

    private boolean not;

    public IsTypeExpression(TExpression left, IsType type) {
        this.left = left;
        this.type = type;
    }

    public IsTypeExpression(TExpression left, IsType type, boolean not) {
        this.left = left;
        this.type = type;
        this.not = not;
    }

    @Override
    public Object eval(Context e) {
        Object bo = left.eval(e);
        boolean r;
        if (bo == null) {
            r = type.check(null);
        } else {
            r = type.check(bo);
        }
        return not ? !r : r;
    }

    @Override
    public String toString() {
        return String.format("%s is %s%s", left, (not ? "not " : ""), type);
    }
}
