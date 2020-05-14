package compiler

import (
	"fmt"
	"testing"
)

type valuegetter map[string]interface{}

func MakeVG(name string, val interface{}) valuegetter {
	m := make(map[string]interface{})
	m[name] = val
	return m
}

func (vg valuegetter) Get(name string) interface{} {
	v, ok := vg[name]
	if ok {
		return v
	}
	return nil
}

func TestEval(t *testing.T) {
	expr, err := Compile("15")
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	v, err := expr.Eval(nil)
	fmt.Println(v)

	expr, err = Compile("1 + 2")
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	v, err = expr.Eval(nil)
	fmt.Println(v)

	expr, err = Compile("1 + $v")
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	vg := MakeVG("$v", 15.5)
	v, err = expr.Eval(vg)
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	fmt.Printf("%s == %.1f\n", expr, v)

	expr, err = Compile("$v - 10")
	v, err = expr.Eval(vg)
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	fmt.Printf("%s == %.1f\n", expr, v)

	expr, _ = Compile("2 * 15 / 10")
	v, err = expr.Eval(nil)
	if v != 3.0 {
		t.FailNow()
	}

	expr, _ = Compile("( 3 + 5) * 10")
	v, err = expr.Eval(nil)
	if v != 80.0 {
		t.FailNow()
	}

	expr, _ = Compile("63 % 10")
	v, err = expr.Eval(nil)
	if v != 3.0 {
		t.FailNow()
	}

	expr, _ = Compile("2 ^ 3")
	v, err = expr.Eval(nil)
	if v != 8.0 {
		t.FailNow()
	}

	expr, _ = Compile("2.5 * 2.0")
	v, err = expr.Eval(nil)
	if v != 5.0 {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("true")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("true && false")
	v, err = expr.Eval(nil)
	if v != false {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("2 * 15 > 5 * 3")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("60 * 1 > 2 && 16 > 5 * 3")
	v, err = expr.Eval(nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	} else if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("60.1 * 1 > 2 * 15 > 5 * 3")
	v, err = expr.Eval(nil)
	if err == nil {
		t.FailNow()
	}

	expr, _ = Compile("60 is integer")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("true is not float")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("true is not float,invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("'hello world' =~ /^hello .*$/")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("'15.5.5' is not ip4")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("'15.5.5.5' is ip4")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("'chuanshi.zl' is host")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	vg = MakeVG("$t", true)
	expr, _ = Compile("$t")
	v, err = expr.Eval(vg)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	if v != true {
		fmt.Println("invalid value ..", v)
		t.FailNow()
	}

	vg = MakeVG("$f", false)
	expr, _ = Compile("$f")
	v, err = expr.Eval(vg)
	if v != false {
		fmt.Println("invalid value ,,", v)
		t.FailNow()
	}

	expr, _ = Compile("'vvv' in ['aaa', 'bbb', 'vvv']")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("'kkk' not in ['aaa', 'bbb', 'vvv']")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("33.3  in [33.3, 44.4, 55.5]")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr, _ = Compile("@source == 'world'")
	vg = MakeVG("@source", "world")
	v, err = expr.Eval(vg)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

}
