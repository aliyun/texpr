package com.aliyun.tauris.expression;

import org.junit.Assert;
import org.junit.Test;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class EvalTest {

    @Test
    public void test() {
        MockContext e = new MockContext("");
        e.addMeta("expr", "$gc in ['1', '2', '3'] && $host == 'www.domain.com'");
        e.setField("host", "www.domain.com");
        e.setField("gc", "3");

        Assert.assertTrue(TExpression.compile("eval(@expr)").check(e));

        e.setField("gc", "5");
        Assert.assertFalse(TExpression.compile("eval(@expr)").check(e));

        e.setField("host", "www.domain1.com");
        e.setField("gc", "3");
        Assert.assertFalse(TExpression.compile("eval(@expr)").check(e));

        e.setField("host", "www.domain.com");
        e.setField("gc", "3");

        long now = System.currentTimeMillis();
        TExpression ex = TExpression.compile("eval(@expr)");
        for (int i = 0; i < 500; i++) {
            ex.check(e);
        }
        System.out.println("use " + (System.currentTimeMillis() - now));
        now = System.currentTimeMillis();
        for (int i = 0; i < 500; i++) {
            TExpression.compile("eval(@expr)").check(e);
        }
        System.out.println("use " + (System.currentTimeMillis() - now));


        e.setField("host", "cs.qa.com");
        e.setField("block", "true");

        Assert.assertTrue(TExpression.compile("$host == 'cs.qa.com' && $block == 'true'").check(e));
        String exp1 = "not ($matched_host == 'cs98.wafqa3.com' && $block_action != '')";
        String exp2 = "$block != ''";
        e.addMeta("default_output_expr", exp2);
        Assert.assertTrue(TExpression.compile(exp2).check(e));
//        Assert.assertFalse(TExpression.compile("(@default_output_expr is null || eval(@default_output_expr))").check(e));
    }
}


