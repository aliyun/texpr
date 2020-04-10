package com.aliyun.tauris.expression;

import org.junit.Assert;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class FuncTest {

    @Test
    public void testEval() {
        MockContext e = new MockContext("");
        e.setField("name", "world");
        e.setField("empty", "");
        e.setField("status", 302);
        e.setField("flag", false);
        e.setField("expr", "$flag == true");
        Assert.assertTrue(TExpression.compile("eval('1 + 1') == 2").check(e));
        Assert.assertFalse(TExpression.compile("eval($expr)").check(e));
    }

    @Test
    public void testExcept() {
        try {
            TExpression.compile("eval('1 + $') == 2").check(new MockContext(""));
            Assert.assertTrue(false);
        } catch (IllegalArgumentException e) {
            Assert.assertTrue(true);
        }
    }

    @Test
    public void testCache() {
        System.setProperty("texpr.max_cache_size", "100");
        TExpression expr = TExpression.compile("eval($de) == $val + 1");
        MockContext e = new MockContext("");
        for (int i = 0; i < 102; i++) {
            e.setField("val", i);
            e.setField("de", "1 + " + i);
            Assert.assertTrue(expr.check(e));
        }
        System.clearProperty("texpr.max_cache_size");
    }

    @Test
    public void testCoCache() throws Exception {
        System.setProperty("texpr.max_cache_size", "100");
        TExpression expr = TExpression.compile("eval($de) == $val + 1");
        List<Thread> ex =  new ArrayList<>();
        for (int i = 0; i < 100; i++) {
            ex.add(new Thread(() -> {
                MockContext e = new MockContext("");
                for (int x = 0; x < 102; x++) {
                    e.setField("val", x);
                    e.setField("de", "1 + " + x);
                    boolean b = expr.check(e);
                    Assert.assertTrue(b);
                }
            }));
        }
        for (Thread t: ex) {
            t.start();
        }
        for (Thread t: ex) {
            t.join();
        }
        System.clearProperty("texpr.max_cache_size");
    }
}
