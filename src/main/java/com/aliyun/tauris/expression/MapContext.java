package com.aliyun.tauris.expression;

import com.aliyun.tauris.expression.Context;

import java.util.concurrent.ConcurrentHashMap;

/**
 * Created by ZhangLei on 16/12/8.
 */
public class MapContext implements Context {

    private ConcurrentHashMap<String, Object> values = new ConcurrentHashMap<>();

    @Override
    public Object get(String name) {
        return values.get(name);
    }

    @Override
    public void set(String name, Object value) {
        values.put(name, value);
    }
}
