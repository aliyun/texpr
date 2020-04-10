package com.aliyun.tauris.expression;

import org.junit.Assert;
import org.junit.Test;

/**
 * Created by ZhangLei on 2017/11/24.
 */
public class IsHostTypeTest {

    @Test
    public void test() {
        IsHostType isHostType = new IsHostType();

        Assert.assertTrue(isHostType.check("www.taobao.com"));
        Assert.assertTrue(isHostType.check("*.taobao.com"));
        Assert.assertTrue(isHostType.check("localhost"));
        Assert.assertTrue(isHostType.check("a.b.c.d.e.f"));
    }
}
