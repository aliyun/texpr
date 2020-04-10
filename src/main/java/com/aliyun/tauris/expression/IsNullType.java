package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsNullType implements IsType {

    @Override
    public boolean check(Object value) {
        return value == null;
    }

    @Override
    public String toString() {
        return "null";
    }
}
