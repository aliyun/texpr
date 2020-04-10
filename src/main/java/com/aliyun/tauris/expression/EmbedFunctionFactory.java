package com.aliyun.tauris.expression;

import java.util.Map;
import java.util.ServiceLoader;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class EmbedFunctionFactory {

    private Map<String, EmbedFunction> funcs = new ConcurrentHashMap<>();

    public EmbedFunctionFactory() {
        ServiceLoader<? extends EmbedFunction> cvt = ServiceLoader.load(EmbedFunction.class);
        for (EmbedFunction t : cvt) {
            funcs.put(t.name(), t);
        }
    }

    public Object newInstance(String name) {
        EmbedFunction t = funcs.get(name);
        if (t == null) {
            throw new IllegalArgumentException("unknown function " + name);
        }
        return t;
    }
}
