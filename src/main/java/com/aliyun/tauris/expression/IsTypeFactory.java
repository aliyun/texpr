package com.aliyun.tauris.expression;

import com.aliyun.tauris.expression.IsType;

import java.util.Map;
import java.util.ServiceLoader;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsTypeFactory  {

    private Map<String, IsType> types = new ConcurrentHashMap<>();

    public IsTypeFactory() {
        ServiceLoader<? extends IsType> cvt = ServiceLoader.load(IsType.class);
        for (IsType t : cvt) {
            types.put(nameOfType(t), t);
        }
    }

    public Object newInstance(String name) {
        IsType t = types.get(name);
        if (t == null) {
            throw new IllegalArgumentException("unknown type " + name);
        }
        return t;
    }

    private String nameOfType(IsType type) {
        String name = type.getClass().getSimpleName();
        if (name.startsWith("Is")) {
            name = name.substring(2);
        }
        if (name.endsWith("Type")) {
            name = name.substring(0, name.length() - 4);
        }
        return camelToUnderscore(name);
    }

    public static String camelToUnderscore(String name) {
        char[] cs = name.toCharArray();
        StringBuilder sb = new StringBuilder();
        boolean addLowdash = false;
        for (int i = 0; i < cs.length; i++) {
            char c = cs[i];

            if (c >= 'A' && c <= 'Z') {
                if (addLowdash) {
                    sb.append('_');
                    addLowdash = false;
                }
                int delta = c - 'A';
                sb.append((char)('a' + delta));
            } else {
                sb.append(c);
                addLowdash = true;
            }
        }
        return sb.toString();
    }
}
