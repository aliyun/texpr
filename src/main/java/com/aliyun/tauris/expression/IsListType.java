package com.aliyun.tauris.expression;

import java.util.List;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsListType implements IsType {

    @Override
    public boolean check(Object value) {
        return value != null && ((value instanceof List) || value.getClass().isArray());
    }

    @Override
    public String toString() {
        return "list";
    }
}
