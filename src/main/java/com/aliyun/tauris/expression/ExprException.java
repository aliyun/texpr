package com.aliyun.tauris.expression;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class ExprException extends RuntimeException {

    public ExprException(String message) {
        super(message);
    }

    public ExprException(String message, Throwable cause) {
        super(message, cause);
    }
}
