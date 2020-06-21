package texpr

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/aliyun/texpr/ast"
)

type ValueGetter interface {
	Get(name string) interface{}
}

func Compile(expr string) (Expression, error) {
	input := antlr.NewInputStream(expr)
	lexer := ast.NewTExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := ast.NewTExprParser(stream)
	visitor := &ExprVisitor{}
	return p.Parse().Accept(visitor).(Expression), nil
}

func MustCompile(expr string) Expression {
	exp, err := Compile(expr)
	if err != nil {
		panic(`texpr: Compile("` + expr + `"): ` + err.Error())
	}
	return exp
}

type Expression interface {
	Eval(vg ValueGetter) (interface{}, error)
	String() string
}

type ValueExpression struct {
	value interface{}
	kind  reflect.Kind
}

func (e *ValueExpression) Eval(vg ValueGetter) (interface{}, error) {
	return e.value, nil
}

func (e *ValueExpression) String() string {
	if e.kind == reflect.Bool {
		return fmt.Sprintf("%t", e.value.(bool))
	}
	if e.kind == reflect.Int64 {
		return fmt.Sprintf("%d", e.value.(int64))
	}
	if e.kind == reflect.Float64 {
		return fmt.Sprintf("%.2f", e.value.(float64))
	}
	if e.kind == reflect.String {
		return fmt.Sprintf("%s", e.value.(string))
	}
	panic("...")
}

func NewValueExpression(value interface{}) *ValueExpression {
	kind := reflect.TypeOf(value).Kind()
	return &ValueExpression{value, kind}
}

type opFunc func(left, right interface{}) interface{}

type OP struct {
	Symbol string
	Func   opFunc
}

func (op *OP) Eval(left, right interface{}) (interface{}, error) {
	return op.Func(left, right), nil
}

func NewOP(symbol string, fun opFunc) *OP {
	return &OP{symbol, fun}
}

func and(left, right interface{}) interface{} {
	return left.(bool) && right.(bool)
}

func or(left, right interface{}) interface{} {
	return left.(bool) || right.(bool)
}

func plus(left, right interface{}) interface{} {
	return left.(float64) + right.(float64)
}

func minus(left, right interface{}) interface{} {
	return left.(float64) - right.(float64)
}

func div(left, right interface{}) interface{} {
	return left.(float64) / right.(float64)
}

func mul(left, right interface{}) interface{} {
	return left.(float64) * right.(float64)
}

func pow(left, right interface{}) interface{} {
	return math.Pow(left.(float64), right.(float64))
}

func mod(left, right interface{}) interface{} {
	return math.Mod(left.(float64), right.(float64))
}

func comparable(items ...interface{}) bool {
	for _, i := range items {
		if !reflect.TypeOf(i).Comparable() {
			return false
		}
	}
	return true
}

func compare(left, right interface{}) int {
	if ls, ok := left.(string); ok {
		if rs, ok := right.(string); ok {
			if ls == rs {
				return 0
			}
			if ls < rs {
				return -1
			}
			return 1
		}
	}
	if ls, ok := left.(float64); ok {
		if rs, ok := right.(float64); ok {
			return int(ls - rs)
		}
	}
	if ls, ok := left.(int64); ok {
		if rs, ok := right.(int64); ok {
			return int(ls - rs)
		}
	}
	if ls, ok := left.(uint64); ok {
		if rs, ok := right.(uint64); ok {
			return int(ls - rs)
		}
	}
	return 1
}

func gt(left, right interface{}) interface{} {
	return compare(left, right) > 0
}

func ge(left, right interface{}) interface{} {
	return compare(left, right) >= 0
}

func lt(left, right interface{}) interface{} {
	return compare(left, right) < 0
}

func le(left, right interface{}) interface{} {
	return compare(left, right) <= 0
}

func eq(left, right interface{}) interface{} {
	return compare(left, right) == 0
}

type BiExpr struct {
	left  Expression
	right Expression
	op    *OP
}

func (e *BiExpr) Eval(vg ValueGetter) (interface{}, error) {
	lv, err := e.left.Eval(vg)
	if err != nil {
		return nil, err
	}
	rv, err := e.right.Eval(vg)
	if err != nil {
		return nil, err
	}
	if reflect.TypeOf(lv) != reflect.TypeOf(rv) {
		msg := fmt.Sprintf("%s is %s, but %s is %s", e.left.String(), reflect.TypeOf(lv), e.right.String(), reflect.TypeOf(rv))
		return nil, fmt.Errorf("invalid expression `%s %s %s`, %s", e.left.String(), e.op.Symbol, e.right.String(), msg)
	}
	if !comparable(lv, rv) {
		return nil, fmt.Errorf("cannot compared %s with %s", reflect.TypeOf(lv), reflect.TypeOf(rv))
	}
	ret, err := e.op.Eval(lv, rv)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (e *BiExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.left.String(), e.op.Symbol, e.right.String())
}

func NewBiBoolExpr(left, right Expression, op *OP) Expression {
	return &BiExpr{left, right, op}
}

func NewBiCalcExpr(left, right Expression, op *OP) Expression {
	return &BiExpr{left, right, op}
}

type NotExpr struct {
	right Expression
}

func (e *NotExpr) Eval(vg ValueGetter) (interface{}, error) {
	b, err := e.right.Eval(vg)
	if err != nil {
		return nil, err
	}
	return !(b.(bool)), nil
}

func (e *NotExpr) String() string {
	return fmt.Sprintf("not %s", e.right.String())
}

type isType func(interface{}) bool

var isTypeFuncs map[string]isType

func isString(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.ConvertibleTo(reflect.TypeOf("hello"))
}

func isInteger(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.ConvertibleTo(reflect.TypeOf(int64(0)))
}

func isFloat(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.ConvertibleTo(reflect.TypeOf(float64(0)))
}

func isBoolean(v interface{}) bool {
	t := reflect.TypeOf(v).Kind()
	return t == reflect.Bool
}

func isNull(v interface{}) bool {
	return v == nil
}

func isEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	s := reflect.TypeOf(v).Kind()
	if s != reflect.String {
		return false
	}
	return strings.TrimSpace(v.(string)) == ""
}

var ip4MatchPattern = regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+$")
var hostMatchPattern = regexp.MustCompile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")

func isIP4(v interface{}) bool {
	if v == nil {
		return true
	}
	s := reflect.TypeOf(v).Kind()
	if s != reflect.String {
		return false
	}
	return ip4MatchPattern.Match([]byte(v.(string)))
}

func isHost(v interface{}) bool {
	if v == nil {
		return true
	}
	s := reflect.TypeOf(v).Kind()
	if s != reflect.String {
		return false
	}
	host := v.(string)
	if strings.HasPrefix(host, "*.") {
		host = host[2:]
	}
	return ip4MatchPattern.Match([]byte(v.(string))) || hostMatchPattern.Match([]byte(v.(string)))
}

func init() {
	isTypeFuncs = make(map[string]isType)
	isTypeFuncs["string"] = isString
	isTypeFuncs["integer"] = isInteger
	isTypeFuncs["float"] = isFloat
	isTypeFuncs["boolean"] = isBoolean
	isTypeFuncs["host"] = isHost
	isTypeFuncs["ip4"] = isIP4
	isTypeFuncs["empty"] = isEmpty
	isTypeFuncs["null"] = isNull
}

func valueIsType(value interface{}, t string) (bool, error) {
	f, ok := isTypeFuncs[t]
	if !ok {
		return false, fmt.Errorf("invalid type %s", t)
	}
	return f(value), nil
}

type IsTypeExpr struct {
	left Expression
	not  bool
	kind string
}

func (e *IsTypeExpr) Eval(vg ValueGetter) (interface{}, error) {
	b, err := e.left.Eval(vg)
	if err != nil {
		return nil, err
	}
	v, err := valueIsType(b, e.kind)
	if e.not {
		return !v, err
	}
	return v, err
}

func (e *IsTypeExpr) String() string {
	return fmt.Sprintf("%s is %s", e.left.String(), e.kind)
}

type VariableExpr struct {
	name string
}

func (e *VariableExpr) Eval(vg ValueGetter) (interface{}, error) {
	val := vg.Get(e.name)
	if val == nil {
		return 0, fmt.Errorf("value %s is nil", e.name)
	}
	return val, nil
}

func (e *VariableExpr) String() string {
	return e.name
}

type ArrayExpr struct {
	Array []Expression
	Type  reflect.Type
}

func (e *ArrayExpr) Eval(vg ValueGetter) (interface{}, error) {
	array := make([]interface{}, len(e.Array))
	for idx, _ := range array {
		v, err := e.Array[idx].Eval(vg)
		if err != nil {
			return nil, err
		}
		array[idx] = v
	}
	return array, nil
}

func (e *ArrayExpr) String() string {
	array := make([]string, len(e.Array))
	for idx := range array {
		array[idx] = e.Array[idx].String()
	}
	return fmt.Sprintf("[%s]", strings.Join(array, ","))
}

type InExpr struct {
	left  Expression
	array *ArrayExpr
	not   bool
}

func (e *InExpr) Eval(vg ValueGetter) (interface{}, error) {
	val, err := e.left.Eval(vg)
	if err != nil {
		return nil, err
	}
	ary, err := e.array.Eval(vg)
	if err != nil {
		return nil, err
	}
	ret := ListContainsValue(val, ary.([]interface{}), e.array.Type)
	if e.not {
		return !ret, nil
	}
	return ret, nil
}

func (e *InExpr) String() string {
	return fmt.Sprintf("%s in %s", e.left.String(), e.array.String())
}

func ListContainsValue(val interface{}, array []interface{}, typ reflect.Type) bool {
	t := reflect.TypeOf(val)
	var nv interface{}
	if t.ConvertibleTo(typ) {
		nv = reflect.ValueOf(val).Convert(typ).Interface()
	}
	for _, x := range array {
		if x == nv {
			return true
		}
	}
	return false
}

type MatchExpr struct {
	left  Expression
	regex *regexp.Regexp
}

func (e *MatchExpr) Eval(vg ValueGetter) (interface{}, error) {
	v, err := e.left.Eval(vg)
	if err != nil {
		return nil, err
	}
	if reflect.TypeOf(v).Kind() != reflect.String {
		return nil, fmt.Errorf("cannot match %s(%s) to float64", reflect.TypeOf(v), v)
	}
	s := v.(string)
	return e.regex.Match([]byte(s)), nil
}

func (e *MatchExpr) String() string {
	return fmt.Sprintf("%s =~ /%s/", e.left.String(), e.regex)
}

func NewMatchExpr(left Expression, regex string) (Expression, error) {
	r := regexp.MustCompile(regex)
	return &MatchExpr{left, r}, nil
}

type ExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *ExprVisitor) VisitParse(ctx *ast.ParseContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *ExprVisitor) VisitBinaryExpression(ctx *ast.BinaryExpressionContext) interface{} {
	op := ctx.Binary().Accept(v).(*OP)
	left := ctx.Expression(0).Accept(v).(Expression)
	right := ctx.Expression(1).Accept(v).(Expression)
	return NewBiBoolExpr(left, right, op)
}

func (v *ExprVisitor) VisitMatchExpression(ctx *ast.MatchExpressionContext) interface{} {
	left := ctx.Expression().Accept(v).(Expression)
	regex := ctx.Regex().Accept(v).(*regexp.Regexp)
	return &MatchExpr{left, regex}
}

func (v *ExprVisitor) VisitInExpression(ctx *ast.InExpressionContext) interface{} {
	left := ctx.Expression().Accept(v).(Expression)
	array := ctx.Array().Accept(v).(*ArrayExpr)
	return &InExpr{left, array, false}
}

func (v *ExprVisitor) VisitIsTypeExpression(ctx *ast.IsTypeExpressionContext) interface{} {
	kind := ctx.Kind().GetText()
	val := ctx.Expression().Accept(v).(Expression)
	return &IsTypeExpr{val, false, kind}
}

func (v *ExprVisitor) VisitCalcExpression(ctx *ast.CalcExpressionContext) interface{} {
	return ctx.Calc().Accept(v)
}

func (v *ExprVisitor) VisitIsNotTypeExpression(ctx *ast.IsNotTypeExpressionContext) interface{} {
	kind := ctx.Kind().GetText()
	val := ctx.Expression().Accept(v).(Expression)
	return &IsTypeExpr{val, true, kind}
}

func (v *ExprVisitor) VisitNotExpression(ctx *ast.NotExpressionContext) interface{} {
	expr := ctx.Expression().Accept(v).(Expression)
	return &NotExpr{expr}
}

func (v *ExprVisitor) VisitParenExpression(ctx *ast.ParenExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *ExprVisitor) VisitNotInExpression(ctx *ast.NotInExpressionContext) interface{} {
	left := ctx.Expression().Accept(v).(Expression)
	array := ctx.Array().Accept(v).(*ArrayExpr)
	return &InExpr{left, array, true}
}

func (v *ExprVisitor) VisitComparatorExpression(ctx *ast.ComparatorExpressionContext) interface{} {
	op := ctx.Comparator().Accept(v).(*OP)
	left := ctx.Expression(0).Accept(v).(Expression)
	right := ctx.Expression(1).Accept(v).(Expression)
	return NewBiBoolExpr(left, right, op)
}

func (v *ExprVisitor) VisitVariableExpression(ctx *ast.VariableExpressionContext) interface{} {
	return ctx.Variable().Accept(v)
}

func (v *ExprVisitor) VisitVariable(ctx *ast.VariableContext) interface{} {
	if ctx.VARIABLE() != nil {
		name := ctx.VARIABLE().GetText()
		return &VariableExpr{name}
	}
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	}
	if ctx.Array() != nil {
		return ctx.Array().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitComparator(ctx *ast.ComparatorContext) interface{} {
	if ctx.GT() != nil {
		return &OP{">", gt}
	}
	if ctx.GE() != nil {
		return &OP{">=", ge}
	}
	if ctx.LT() != nil {
		return &OP{"<", lt}
	}
	if ctx.LE() != nil {
		return &OP{"<=", le}
	}
	if ctx.EQ() != nil {
		return &OP{"==", eq}
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitBinary(ctx *ast.BinaryContext) interface{} {
	if ctx.AND() != nil {
		return &OP{"&&", and}
	}
	if ctx.OR() != nil {
		return &OP{"||", or}
	}
	return fmt.Errorf("invalid binary op type %s", ctx.GetText())
}

func (v *ExprVisitor) VisitBoolean(ctx *ast.BooleanContext) interface{} {
	return NewValueExpression(ctx.GetText() == "true")
}

func (v *ExprVisitor) VisitLiteral(ctx *ast.LiteralContext) interface{} {
	if ctx.Boolean() != nil {
		vt, _ := strconv.ParseBool(ctx.Boolean().GetText())
		return NewValueExpression(vt)
	}
	if ctx.Integer() != nil {
		vt, _ := strconv.ParseFloat(ctx.Integer().GetText(), 64)
		return NewValueExpression(vt)
	}
	if ctx.Float() != nil {
		vt, _ := strconv.ParseFloat(ctx.Float().GetText(), 64)
		return NewValueExpression(vt)
	}
	if ctx.Varchar() != nil {
		s := ctx.GetText()
		l := len(s)
		return NewValueExpression(s[1 : l-1])
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitArray(ctx *ast.ArrayContext) interface{} {
	if ctx.Strings() != nil {
		return ctx.Strings().Accept(v)
	}
	if ctx.Integers() != nil {
		return ctx.Integers().Accept(v)
	}
	if ctx.Floats() != nil {
		return ctx.Floats().Accept(v)
	}
	if ctx.Booleans() != nil {
		return ctx.Booleans().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitCalc(ctx *ast.CalcContext) interface{} {
	return ctx.Plus().Accept(v)
}

func (v *ExprVisitor) VisitPlus(ctx *ast.PlusContext) interface{} {
	if ctx.MINUS() != nil {
		left := ctx.Plus().Accept(v).(Expression)
		right := ctx.Multiplying().Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("-", minus))
	} else if ctx.PLUS() != nil {
		left := ctx.Plus().Accept(v).(Expression)
		right := ctx.Multiplying().Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("+", plus))
	}
	return ctx.Multiplying().Accept(v)
}

func (v *ExprVisitor) VisitMultiplying(ctx *ast.MultiplyingContext) interface{} {
	if ctx.DIV() != nil {
		left := ctx.Multiplying().Accept(v).(Expression)
		right := ctx.Pow().Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("/", div))
	}
	if ctx.MUL() != nil {
		left := ctx.Multiplying().Accept(v).(Expression)
		right := ctx.Pow().Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("*", mul))
	}
	if ctx.MOD() != nil {
		left := ctx.Multiplying().Accept(v).(Expression)
		right := ctx.Pow().Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("%", mod))
	}
	return ctx.Pow().Accept(v)
}

func (v *ExprVisitor) VisitPow(ctx *ast.PowContext) interface{} {
	if ctx.POW() != nil {
		left := ctx.Atom(0).Accept(v).(Expression)
		right := ctx.Atom(1).Accept(v).(Expression)
		return NewBiCalcExpr(left, right, NewOP("^", pow))
	}
	return ctx.Atom(0).Accept(v)
}

func (v *ExprVisitor) VisitAtom(ctx *ast.AtomContext) interface{} {
	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		return ctx.Plus().Accept(v)
	}
	if ctx.Variable() != nil {
		return ctx.Variable().Accept(v)
	}
	if ctx.Scientific() != nil {
		return ctx.Scientific().Accept(v)
	}
	if ctx.Function() != nil {
		return ctx.Function().Accept(v)
	}

	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitScientific(ctx *ast.ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitFunction(ctx *ast.FunctionContext) interface{} {
	params := ctx.Parameters().Accept(v).([]interface{})
	fname := ctx.Funcname().Accept(v).(string)
	f, b := funcMap[fname]
	if !b {
		return errors.New("unknown function " + fname)
	}
	exprs := make([]Expression, len(params))
	for i, _ := range params {
		exprs[i] = params[i].(Expression)
	}
	return NewFuncExpr(fname, f, exprs)
}

func (v *ExprVisitor) VisitFuncname(ctx *ast.FuncnameContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (v *ExprVisitor) VisitParameters(ctx *ast.ParametersContext) interface{} {
	var params []interface{}
	allExprs := ctx.AllExpression()
	for i, _ := range allExprs {
		params = append(params, allExprs[i].Accept(v))
	}
	return params
}


func (v *ExprVisitor) VisitNumber(ctx *ast.NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitRegex(ctx *ast.RegexContext) interface{} {
	regex := ctx.Regex().GetText()
	l := len(regex)
	return regexp.MustCompile(regex[1 : l-1])
}

func (v *ExprVisitor) VisitKind(ctx *ast.KindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitStrings(ctx *ast.StringsContext) interface{} {
	ss := make([]Expression, len(ctx.AllVarchar()))
	for idx, _ := range ss {
		s := ctx.Varchar(idx).GetText()
		l := len(s)
		ss[idx] = NewValueExpression(s[1 : l-1])
	}
	return &ArrayExpr{ss, reflect.TypeOf("")}
}

func (v *ExprVisitor) VisitIntegers(ctx *ast.IntegersContext) interface{} {
	ss := make([]Expression, len(ctx.AllInteger()))
	for idx, _ := range ss {
		s := ctx.Integer(idx).GetText()
		f, _ := strconv.ParseInt(s, 10, 64)
		ss[idx] = NewValueExpression(f)
	}
	return &ArrayExpr{ss, reflect.TypeOf(int64(0))}
}

func (v *ExprVisitor) VisitFloats(ctx *ast.FloatsContext) interface{} {
	ss := make([]Expression, len(ctx.AllFloat()))
	for idx, _ := range ss {
		s := ctx.Float(idx).GetText()
		f, _ := strconv.ParseFloat(s, 64)
		ss[idx] = NewValueExpression(f)
	}
	return &ArrayExpr{ss, reflect.TypeOf(float64(0))}
}

func (v *ExprVisitor) VisitBooleans(ctx *ast.BooleansContext) interface{} {
	ss := make([]Expression, len(ctx.AllBoolean()))
	for idx, _ := range ss {
		s := ctx.Boolean(idx).GetText()
		f, _ := strconv.ParseBool(s)
		ss[idx] = NewValueExpression(f)
	}
	return &ArrayExpr{ss, reflect.TypeOf(true)}
}
