package com.aliyun.tauris.expression;

/**
 * Class Context
 *
 * @author yundun-waf-dev
 * @date 2019-03-11
 */
public interface Context {

    /**
     * get value by name
     * @param name context variable name, the first character of name is '$' or '@'
     * @return context variable value
     */
    Object get(String name);

    default void set(String name, Object value) {

    }
}
