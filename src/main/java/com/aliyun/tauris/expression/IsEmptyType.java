package com.aliyun.tauris.expression;



import java.util.Collection;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsEmptyType implements IsType {

    @Override
    public boolean check(Object value) {
        if (value == null) {
            return true;
        }
        if (value instanceof String) {
            return ((String)value).trim().isEmpty();
        }
        if (value instanceof Collection) {
            return ((Collection)value).isEmpty();
        }
        return false;
    }

    @Override
    public String toString() {
        return "empty";
    }
}
