package com.aliyun.tauris.expression;


import java.util.regex.Pattern;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class IsHostType implements IsType {

    private static final Pattern IP4_PATTERN = Pattern.compile("^\\d+\\.\\d+\\.\\d+\\.\\d+$");

    private static final Pattern HOSTNAME_PATTERN = Pattern.compile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$");

    @Override
    public boolean check(Object value) {
        if (value == null || !(value instanceof String)) {
            return false;
        }
        String host = (String) value;
        if (IP4_PATTERN.matcher(host).matches()) {
            return true;
        }
        if (host.startsWith("*.")) {
            host = host.substring(2);
        }
        return HOSTNAME_PATTERN.matcher(host).matches();
    }

    @Override
    public String toString() {
        return "host";
    }
}
