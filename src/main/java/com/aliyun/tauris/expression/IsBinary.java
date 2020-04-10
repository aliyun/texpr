package com.aliyun.tauris.expression;


/**
 * 判断变量字符串是否包含字符NUL
 * Created by ZhangLei on 17/5/26.
 */
public class IsBinary implements IsType {

    private static final char NULL = '\0';

    @Override
    public boolean check(Object value) {
        if (value == null || !(value instanceof String)) {
            return false;
        }
        String str = (String)value;
        for (int i = 0; i < str.length(); i++) {
            if (str.charAt(i) == NULL) {
                return true;
            }
        }
        return false;
    }

    @Override
    public String toString() {
        return "binary";
    }
}
