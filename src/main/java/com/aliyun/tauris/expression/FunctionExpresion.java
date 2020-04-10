package com.aliyun.tauris.expression;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Class FunctionExpresion
 *
 * @author yundun-waf-dev
 * @date 2019-12-16
 */
public class FunctionExpresion extends TExpression {


    private List<TExpression> params;

    private EmbedFunction function;

    public FunctionExpresion(EmbedFunction function, List<TExpression> params) {
        this.function = function;
        this.params = params;
    }

    @Override
    public Object eval(Context ctx) {
        List<Object> ps = params.stream().map((p) -> p.eval(ctx)).collect(Collectors.toList());
        return function.execute(ctx, ps);
    }

    @Override
    public String toString() {
        List<String> ps = params.stream().map(TExpression::toString).collect(Collectors.toList());
        return String.format("%s(%s)", function.name(), String.join(", ", ps));
    }
}
