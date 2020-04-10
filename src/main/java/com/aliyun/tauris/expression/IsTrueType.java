package com.aliyun.tauris.expression;


/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsTrueType implements IsType {

    @Override
    public boolean check(Object value) {
        return "true".equals(value);
    }

    @Override
    public String toString() {
        return "true";
    }
}
