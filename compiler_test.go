package texpr

import (
	"fmt"
	"math"
	"reflect"
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
	var err error
	var expr Expression
	expr, err = Compile("15")
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	v, err := expr.Eval(nil)
	if v != 15.0 {
		t.Fatalf("expect 15, actual %.1f", v)
	}
}

func TestEval2(t *testing.T) {
	var err error
	var expr Expression
	var v interface{}
	expr, err = Compile("1 + 2")
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	v, _ = expr.Eval(nil)
	if v != 3.0 {
		t.Fatalf("expect 3, actual %.1f", v)
	}
}

func TestEval3(t *testing.T) {
	var err error
	var expr Expression
	var v interface{}
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

	expr = MustCompile("$v - 10")
	v, err = expr.Eval(vg)
	if err != nil {
		t.Fatalf("expression err, %s", err)
	}
	fmt.Printf("%s == %.1f\n", expr, v)

	expr = MustCompile("2 * 15 / 10")
	v, err = expr.Eval(nil)
	if v != 3.0 {
		t.FailNow()
	}

	expr = MustCompile("( 3 + 5) * 10")
	v, err = expr.Eval(nil)
	if v != 80.0 {
		t.Fatalf("expect 80, actual %.1f", v)
	}

	expr = MustCompile("63 % 10")
	v, err = expr.Eval(nil)
	if v != 3.0 {
		t.Fatalf("expect 3.0, actual %.1f", v)
	}

	expr = MustCompile("2 ^ 3")
	v, err = expr.Eval(nil)
	if v != int64(1) {
		t.Fatalf("expect 1, actual %d", v)
	}

	expr = MustCompile("2.5 * 2.0")
	v, err = expr.Eval(nil)
	if v != 5.0 {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("true")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("true && false")
	v, err = expr.Eval(nil)
	if v != false {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("2 * 15 > 5 * 3")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("60 * 1 > 2 && 16 > 5 * 3")
	v, err = expr.Eval(nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	} else if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("60.1 * 1 > 2 * 15 > 5 * 3")
	v, err = expr.Eval(nil)
	if err == nil {
		t.FailNow()
	}

	expr = MustCompile("60 is integer")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("true is not float")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("true is not float,invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("'hello world' =~ /^hello .*$/")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("'15.5.5' is not ip4")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("'15.5.5.5' is ip4")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("'chuanshi.zl' is host")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	vg = MakeVG("$t", true)
	expr = MustCompile("$t")
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
	expr = MustCompile("$f")
	v, err = expr.Eval(vg)
	if v != false {
		fmt.Println("invalid value ,,", v)
		t.FailNow()
	}

	expr, err = Compile("'vvv' in ['aaa', 'bbb', 'vvv']")
	if err != nil {
		panic(err)
	}
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("'kkk' not in ['aaa', 'bbb', 'vvv']")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("33.3  in [33.3, 44.4, 55.5]")
	v, err = expr.Eval(nil)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

	expr = MustCompile("@source == 'world'")
	vg = MakeVG("@source", "world")
	v, err = expr.Eval(vg)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}

}

func TestRegex(t *testing.T) {
	expr := MustCompile("$val =~ /wo.*/ && $val =~/.*ld$/")
	vg := MakeVG("$val", "world")
	v, _ := expr.Eval(vg)
	if v != true {
		fmt.Println("invalid value", v)
		t.FailNow()
	}
}

func TestEvalFuncPlus(t *testing.T) {
	s := `eval('1+1')`
	expr := MustCompile(s)
	r, err := expr.Eval(nil)
	if err != nil {
		panic(err)
	}
	v := r.(float64)
	if float64(2) != r {
		t.Fatalf("expect 2, actual %.1f", v)
	}
}

type PowFunc struct {
	Function
}

func (f *PowFunc) Name() string {
	return "pow"
}

func (f *PowFunc) Execute(vg ValueGetter, params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf(`pow() takes exactly 2 argument (" + %d + " given)`, len(params))
	}
	x, err := toFloat64(params[0])
	if err != nil {
		return nil, err
	}
	y, err := toFloat64(params[1])
	if err != nil {
		return nil, err
	}
	return math.Pow(x, y), nil
}

func toFloat64(v interface{}) (float64, error) {
	switch v.(type) {
	case float64:
		return v.(float64), nil
	case int:
		return float64(v.(int)), nil
	case int8:
		return float64(v.(int8)), nil
	case int16:
		return float64(v.(int16)), nil
	case int32:
		return float64(v.(int32)), nil
	case int64:
		return float64(v.(int64)), nil
	case float32:
		return float64(v.(float32)), nil
	default:
		return 0, fmt.Errorf(`pow() need string, %s found`, reflect.TypeOf(v).String())
	}
}

func TestEvalFuncPow(t *testing.T) {
	RegisterFunc(&PowFunc{})
	s := `pow(2,3)`
	expr := MustCompile(s)
	r, err := expr.Eval(nil)
	if err != nil {
		panic(err)
	}
	v := r.(float64)
	if float64(8) != r {
		t.Fatalf("expect 8, actual %.1f", v)
	}
}