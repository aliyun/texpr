package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsString implements IsType {

    @Override
    public boolean check(Object value) {
        return value instanceof String;
    }

    @Override
    public String toString() {
        return "string";
    }
}
