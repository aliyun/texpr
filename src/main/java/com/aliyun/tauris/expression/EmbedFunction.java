package com.aliyun.tauris.expression;

import java.util.List;

/**
 * Created by zhanglei on 2019/12/16.
 */
public interface EmbedFunction {

    String name();
    Object execute(Context ctx, List<Object> params);

}
