package com.aliyun.tauris.expression.functions;

import com.aliyun.tauris.expression.Context;
import com.aliyun.tauris.expression.EmbedFunction;
import com.aliyun.tauris.expression.ExprException;
import com.aliyun.tauris.expression.TExpression;

import java.util.List;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentLinkedQueue;

/**
 * Class EvalFunction
 *
 * @author yundun-waf-dev
 * @date 2019-12-16
 */
public class EvalFunction implements EmbedFunction {

    public static final String MAX_CACHE_SIZE = "texpr.max_cache_size";

    private ConcurrentHashMap<String, TExpression> cache     = new ConcurrentHashMap<>();
    private ConcurrentLinkedQueue<String>          cacheKeys = new ConcurrentLinkedQueue<>();

    private long minCacheSize = 0;
    private long maxCacheSize = 0;

    public EvalFunction() {
        try {
            maxCacheSize = Long.parseLong(System.getProperty(MAX_CACHE_SIZE, "0"));
        } catch (NumberFormatException e) {
            System.err.println("illegal number format of -Dtexpr.max_cache_size:" + System.getProperty(MAX_CACHE_SIZE));
        }
        if (maxCacheSize > 0 && maxCacheSize < 100) {
            maxCacheSize = 100;
        }
        minCacheSize = (int)(maxCacheSize * 0.25);
    }

    @Override
    public String name() {
        return "eval";
    }

    @Override
    public Object execute(Context ctx, List<Object> params) {
        if (params.size() != 1) {
            throw new ExprException("eval() takes exactly 1 argument (" + params.size() + " given)");
        }
        Object param = params.get(0);
        if (param == null) {
            throw new ExprException("eval() need string, but null found");
        }
        if (!(param instanceof String)) {
            throw new ExprException("eval() need string, " + param.getClass().getName() + " found");
        }
        TExpression expr = cache.computeIfAbsent((String) param, (s) -> {
            try {
                TExpression e = TExpression.compile(s);
                cacheKeys.add(s);
                return e;
            } catch (Exception e) {
                throw new IllegalArgumentException(e.getMessage(), e);
            }
        });
        if (maxCacheSize > 0 && cacheKeys.size() > maxCacheSize) {
            // clear cache
            while (cacheKeys.size() > minCacheSize) {
                String k = cacheKeys.peek();
                if (k == null) {
                    break;
                }
                cacheKeys.remove(k);
            }
        }
        if (expr != null) {
            return expr.eval(ctx);
        }
        return null;
    }

    public int getCacheSize() {
        return cacheKeys.size();
    }

    @Override
    public String toString() {
        return "eval()";
    }
}
