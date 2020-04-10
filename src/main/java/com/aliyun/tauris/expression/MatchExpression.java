package com.aliyun.tauris.expression;



import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Created by ZhangLei on 17/5/26.
 */
public class MatchExpression extends TExpression {

    private TExpression left;
    private Pattern     pattern;

    private Set<String> namedGroups;

    public MatchExpression(TExpression left, Pattern pattern) {
        this.left = left;
        this.pattern = pattern;
        this.namedGroups = getNamedGroupCandidates(pattern.pattern());
    }

    @Override
    public Object eval(Context ctx) {
        Object bo = left.eval(ctx);
        if (bo == null) {
            return false;
        }
        Matcher m = pattern.matcher(bo.toString());
        if (m.matches()) {
            if (!namedGroups.isEmpty()) {
                for (String name: namedGroups) {
                    String value = m.group(name);
                    if (value != null) {
                        ctx.set(name, value);
                    }
                }
            }

            return true;
        }
        return false;
    }

    private static Set<String> getNamedGroupCandidates(String regex) {
        Set<String> namedGroups = new TreeSet<>();
        Matcher m = Pattern.compile("\\(\\?<([a-zA-Z][a-zA-Z0-9]*)>").matcher(regex);
        while (m.find()) {
            namedGroups.add(m.group(1));
        }
        return Collections.unmodifiableSet(namedGroups);
    }

    public String toString() {
        return String.format("%s =~ /%s/", left, pattern.toString());
    }

}
