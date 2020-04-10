package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public abstract class ComparatorOp {


    public abstract Comparable eval(Context e, TExpression left, TExpression right);

    private static int compare(Context e, TExpression left, TExpression right) {
        Comparable lv = (Comparable)left.eval(e);
        Comparable rv = (Comparable)right.eval(e);
        if (lv == null && rv == null) {
            return 0;
        }
        if (lv == null) {
            return -1;
        }
        if (rv == null) {
            return 1;
        }
        if (lv instanceof Number) {
            lv = ((Number)lv).doubleValue();
        }
        if (rv instanceof Number) {
            rv = ((Number)rv).doubleValue();
        }
        return lv.compareTo(rv);
    }

    public static ComparatorOp build(String op) {
        if (op.equals(">")) {
            return new GT();
        }
        if (op.equals(">=")) {
            return new GT(true);
        }
        if (op.equals("<")) {
            return new LT();
        }
        if (op.equals("<=")) {
            return new LT(true);
        }
        if (op.equals("==")) {
            return new EQ();
        }
        if (op.equals("!=")) {
            return new NE();
        }
        throw new IllegalArgumentException("illegal comparator op :" + op);
    }

    private static class GT extends ComparatorOp {

        private boolean equals;

        public GT() {
        }

        public GT(boolean equals) {
            this.equals = equals;
        }

        @Override
        public Comparable eval(Context e, TExpression left, TExpression right) {
            int x = compare(e, left, right);
            return x > 0 || (equals && x == 0);
        }

        @Override
        public String toString() {
            return equals ? ">=" : ">";
        }
    }

    private static class LT extends ComparatorOp {

        private boolean equals;

        public LT() {
        }

        public LT(boolean equals) {
            this.equals = equals;
        }

        @Override
        public Comparable eval(Context e, TExpression left, TExpression right) {
            int x = compare(e, left, right);
            return x < 0 || (equals && x == 0);
        }

        @Override
        public String toString() {
            return equals ? "<=" : "<";
        }
    }

    private static class EQ extends ComparatorOp {

        @Override
        public Comparable eval(Context e, TExpression left, TExpression right) {
            int x = compare(e, left, right);
            return x == 0;
        }

        @Override
        public String toString() {
            return "==";
        }
    }

    private static class NE extends ComparatorOp {

        @Override
        public Comparable eval(Context e, TExpression left, TExpression right) {
            int x = compare(e, left, right);
            return x != 0;
        }

        @Override
        public String toString() {
            return "!=";
        }
    }



}
