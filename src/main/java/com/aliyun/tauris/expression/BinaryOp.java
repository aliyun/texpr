package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public abstract class BinaryOp {

    public abstract boolean eval(Context e, TExpression left, TExpression right);

    public static BinaryOp build(String op) {
        if (op.equals("&&") || op.equals("and")) {
            return new And();
        }
        if (op.equals("||") || op.equals("or")) {
            return new Or();
        }
        throw new IllegalArgumentException("illegal binary op :" + op);
    }

    private static class Or extends BinaryOp {
        @Override
        public boolean eval(Context e, TExpression left, TExpression right) {
            boolean lv = left.check(e);
            if (lv) {
                return true;
            }
            return right.check(e);
        }

        @Override
        public String toString() {
            return "||";
        }
    }

    private static class And extends BinaryOp {
        @Override
        public boolean eval(Context e, TExpression left, TExpression right) {
            boolean lv = left.check(e);
            if (!lv) return false;
            return right.check(e);
        }
        @Override
        public String toString() {
            return "&&";
        }
    }

}
