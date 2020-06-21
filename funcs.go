package texpr

import (
	"fmt"
	"reflect"
	"strings"
)

var funcMap map[string]Function

func init() {
	funcMap = make(map[string]Function)
	RegisterFunc(&EvalFunc{})
}

func RegisterFunc(f Function) {
	funcMap[f.Name()] = f
}

type Function interface {
	Execute(vg ValueGetter, params []interface{}) (interface{}, error)
	Name() string
}


type FuncExpr struct {
	name string
	function Function
	params []Expression
	repr string
}

func NewFuncExpr(name string, function Function, params []Expression) *FuncExpr {
	vs := make([]string, len(params))
	for i, _ := range params {
		vs[i] = params[i].String()
	}
	return &FuncExpr{
		name: name,
		function: function,
		params: params,
		repr: strings.Join(vs, ", "),
	}
}

func (e *FuncExpr) Eval(vg ValueGetter) (interface{}, error) {
	pvs := make([]interface{}, len(e.params))
	for i, _ := range e.params {
		p, err := e.params[i].Eval(vg)
		if err != nil {
			return nil, err
		}
		pvs[i] = p
	}
	return e.function.Execute(vg, pvs)
}

func (e *FuncExpr) String() string {
	return e.repr
}

type EvalFunc struct {
	Function
}

func (f *EvalFunc) Name() string {
	return "eval"
}

func (f *EvalFunc) Execute(vg ValueGetter, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, fmt.Errorf(`eval() takes exactly 1 argument (" + %d + " given)`, len(params))
	}
	param := params[0]
	pType := reflect.TypeOf(param)
	if pType.Kind() != reflect.String {
		return nil, fmt.Errorf(`eval() need string, %s found`, pType.String())
	}
	expr, err := Compile(param.(string))
	if err != nil {
		return nil, err
	}
	return expr.Eval(vg)
}
