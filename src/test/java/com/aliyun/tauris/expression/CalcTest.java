package com.aliyun.tauris.expression;

import org.junit.Assert;
import org.junit.Test;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class CalcTest {

    @Test
    public void test() {

        MockContext e = new MockContext("");
        e.setField("name", "world");
        e.setField("empty", "");
        e.setField("status", 302);
        e.setField("flag", false);
        e.addMeta("author", "chuanshi.zl");
        Assert.assertTrue(TExpression.compile("1 + 1 == 2").check(e));
        Assert.assertTrue(TExpression.compile("11 - 1 == 10").check(e));
        Assert.assertTrue(TExpression.compile("10 * 2 == 20").check(e));
        Assert.assertTrue(TExpression.compile("20 / 2 == 10").check(e));
        Assert.assertTrue(TExpression.compile("19 % 10 == 9").check(e));
        Assert.assertTrue(TExpression.compile("11 - 10 != 2").check(e));

        Assert.assertEquals(TExpression.compile("19 % 10").calc(e).intValue(), 19 % 10);
        Assert.assertEquals(TExpression.compile("1 + 2 << 3").calc(e).intValue(), 1 + 2 << 3);
        Assert.assertEquals(TExpression.compile("1 + (2 << 3)").calc(e).intValue(), 1 + (2 << 3));
        Assert.assertEquals(TExpression.compile("( 1 ^ 5) + (2 << 3)").calc(e).intValue(),( 1 ^ 5) + (2 << 3));
        Assert.assertEquals(TExpression.compile("2 * 3 + 2").calc(e).intValue(), 2 * 3 + 2);
        Assert.assertEquals(TExpression.compile("4 | 3 + 3").calc(e).intValue(), 4 | 3 + 3);
        Assert.assertEquals(TExpression.compile("4 | 3 * 3").calc(e).intValue(), 4 | 3 * 3);
    }


}
