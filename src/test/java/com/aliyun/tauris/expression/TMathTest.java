package com.aliyun.tauris.expression;

import org.joda.time.DateTime;
import org.joda.time.format.DateTimeFormat;
import org.joda.time.format.DateTimeFormatter;
import org.junit.Assert;
import org.junit.Test;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class TMathTest {

//    @Test
    public void test() {
        Assert.assertEquals(123 + 456, TExpression.compile("123 + 456").calc(null).longValue());
        Assert.assertEquals(1200 - 200, TExpression.compile("1200 - 200").calc(null).longValue());
        Assert.assertEquals(1200 - 2 * 500, TExpression.compile("1200 - 2 * 500").calc(null).longValue());
        Assert.assertEquals((1200 - 200) * 500, TExpression.compile("(1200 - 200) * 500").calc(null).longValue());
        Assert.assertEquals(11 % 2, TExpression.compile("11 % 2").calc(null).longValue());
        Assert.assertEquals((long)(((1978 - 124) * 256 + 123) / 2.5), TExpression.compile("((1978 - 124) * 256 + 123) / 2.5").calc(null).longValue());

        Assert.assertEquals(5 * 6 * 7, TExpression.compile("5 * 6 * 7").calc(null).longValue());
        MockContext ev = new MockContext("");
        ev.set("@abc", 456);
        ev.set("def", 123);

        Assert.assertEquals((456 + 123) * 2, TExpression.compile("(@abc + $def) * 2").calc(ev).longValue());

        long ts = ev.getTimestamp().getTime();
        ts = ts - ts % ( 15 * 60 * 1000);
        DateTimeFormatter sdf = DateTimeFormat.forPattern("'ds'=yyyyMMdd,'hh'=HH,'mi'=mm");
        String val = new DateTime(ts).toString(sdf);
        System.out.println(val);
        Assert.assertEquals(ts, TExpression.compile("@timestamp.time - @timestamp.time % (15 * 60 * 1000)").calc(ev).longValue());
    }

    @Test
    public void testMod() {
        String v = "1509422237661";
        Assert.assertEquals(Long.parseLong(v) % 15, TExpression.compile("1509422237661 % 15").calc(null).longValue());
    }
}
