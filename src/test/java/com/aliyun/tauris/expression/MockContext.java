package com.aliyun.tauris.expression;

import java.lang.reflect.Array;
import java.lang.reflect.Method;
import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Created by ZhangLei on 16/12/8.
 */
public class MockContext implements Context {

    public static final String META_SOURCE    = "@source";
    public static final String META_TIMESTAMP = "@timestamp";

    private ConcurrentHashMap<String, Object> meta = new ConcurrentHashMap<>();

    private ConcurrentHashMap<String, Object> fields = new ConcurrentHashMap<>();

    public MockContext() {
        set(META_TIMESTAMP, new Date());
    }

    public MockContext(String source) {
        this();
        set(META_SOURCE, source);
    }

    public void addMeta(String key, Object value) {
        if (key == null || value == null) {
            throw new IllegalArgumentException("key or value is null");
        }
        this.meta.put(key, value);
    }

    public Map<String, Object> getMeta() {
        return Collections.unmodifiableMap(meta);
    }

    public Object getMeta(String name) {
        return get(meta, name);
    }

    public void set(String name, Object value) {
        if (name == null || name.isEmpty()) throw new IllegalArgumentException("name is null");
        if (value != null) {
            if (name.charAt(0) == '@') {
                meta.put(name.substring(1), value);
            } else {
                setField(name, value);
            }
        } else {
            remove(name);
        }
    }

    public Object get(String name) {
        if (name == null || name.isEmpty()) {
            return null;
        }
        if (name.charAt(0) == '@') {
            return get(meta, name.substring(1));
        } else {
            if (name.charAt(0) == '$') {
                return get(fields, name.substring(1));
            }
            return get(fields, name);
        }
    }

    public Object remove(String name) {
        if (name == null || name.isEmpty()) return null;
        if (name.charAt(0) == '@') {
            return meta.remove(name.substring(1));
        } else {
            return removeField(name);
        }
    }

    private Object get(Map<String, Object> data, String name) {
        if (data.containsKey(name)) {
            return data.get(name);
        } else {
            String[] parts = name.split("\\.");
            Object v = data;
            int elementIndex = 0;
            int leftBracket;

            for (int i = 0; i < parts.length; i++) {
                boolean last = parts.length - 1 == i;
                String part = parts[i];
                boolean valueIsArray = false;

                if (part.charAt(part.length() - 1) == ']') {
                    valueIsArray = true;
                    leftBracket = part.lastIndexOf('[');
                    try {
                        elementIndex = Integer.parseInt(part.substring(leftBracket + 1, part.length() - 1));
                    } catch (NumberFormatException e) {
                        return null;
                    }
                    part = part.substring(0, leftBracket);
                }

                if (v instanceof Map) {
                    v = ((Map) v).get(part);
                    if (v == null) {
                        return null;
                    }
                } else {
                    try {
                        char fc = part.charAt(0);
                        if (fc >= 'a' && fc <= 'z') {
                            fc += 'z' - 'a';
                            char[] cs = part.toCharArray();
                            cs[0] = fc;
                            part = new String(cs);
                        }
                        Method getter = v.getClass().getMethod("get" + part);
                        v = getter.invoke(v);
                    } catch (Exception e) {
                        return null;
                    }
                }
                if (v == null) {
                    return null;
                }
                if (valueIsArray) {
                    try {
                        try {
                            if (v instanceof List) {
                                if (elementIndex < 0) {
                                    elementIndex = ((List) v).size() + elementIndex;
                                }
                                v = ((List) v).get(elementIndex);
                            } else if (v.getClass().isArray()) {
                                if (elementIndex < 0) {
                                    elementIndex = Array.getLength(v) + elementIndex;
                                }
                                v = Array.get(v, elementIndex);
                            } else {
                                return null;
                            }
                        } catch (IndexOutOfBoundsException e) {
                            return null;
                        }

                    } catch (NumberFormatException e) {
                        return null;
                    }
                }
                if (last || v == null) {
                    return v;
                }
            }
            return v;
        }
    }

    public void setField(String name, Object value) {
        if (value == null) return;
        if (name.indexOf('.') >= 0) {
            throw new IllegalArgumentException(String.format("invalid key:`%s`", name));
        }
        fields.put(name, value);
    }

    public Object removeField(String name) {
        if (fields.containsKey(name)) {
            return fields.remove(name);
        } else if (name.indexOf('.') > 0) {
            String[] parts = name.split("\\.");
            Object v = fields;
            for (int i = 0; i < parts.length; i++) {
                boolean last = parts.length - 1 == i;
                if (!last) {
                    if (v == null) {
                        return null;
                    } else if (v instanceof Map) {
                        v = ((Map) v).get(parts[i]);
                    } else {
                        return null;
                    }
                } else {
                    if (v instanceof Map) {
                        return ((Map) v).remove(parts[i]);
                    }
                }
            }
        }
        return null;
    }

    public void setFields(Map<String, Object> fields) {
        this.fields.putAll(fields);
    }

    public Map<String, Object> getFields() {
        return fields;
    }

    public Date getTimestamp() {
        return (Date) get(META_TIMESTAMP);
    }

    public void setTimestamp(Date timestamp) {
        set(META_TIMESTAMP, timestamp);
    }

    public String getSource() {
        return (String) get(META_SOURCE);
    }

}
