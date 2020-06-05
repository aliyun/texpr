package texpr

import (
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
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := ast.NewTExprParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlr.NewConsoleErrorListener())
	visitor := &ExprVisitor{}
	exp := parser.Parse().Accept(visitor)
	return exp.(Expression), nil
}

func MustCompile(expr string) Expression {
	exp, err := Compile(expr)
	if err != nil {
		return nil
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

func band(left, right interface{}) interface{} {
	return left.(int64) & right.(int64)
}

func beor(left, right interface{}) interface{} {
	return int64(left.(float64)) ^ int64(right.(float64))
}

func bior(left, right interface{}) interface{} {
	return int64(left.(float64)) | int64(right.(float64))
}

func lshift(left, right interface{}) interface{} {
	return int64(left.(float64)) << int64(right.(float64))
}

func rshift(left, right interface{}) interface{} {
	return int64(left.(float64)) >> int64(right.(float64))
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

func (e *VariableExpr) contain(vg ValueGetter, sub interface{}) bool {
	val := vg.Get(e.name)
	if val == nil {
		return false
	}
	if val == sub {
		return true
	}
	vt := reflect.TypeOf(val)
	rt := reflect.TypeOf(sub)
	if rt.Kind() == reflect.String && vt.Kind() == reflect.String {
		return strings.Contains(val.(string), sub.(string))
	}
	return false
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
	Array []interface{}
	Type  reflect.Type
	Set   map[interface{}]int
}

func NewArrayExpr(es []Expression, t reflect.Type) *ArrayExpr {
	set := make(map[interface{}]int)
	ary := make([]interface{}, len(es))
	for _, e := range es {
		v, _ := e.Eval(nil)
		set[v] = 0
		ary = append(ary, v)
	}
	return &ArrayExpr{
		Array: ary,
		Type:  t,
		Set:   set,
	}
}

func (e *ArrayExpr) Eval(vg ValueGetter) (interface{}, error) {
	return e.Array, nil
}

func (e *ArrayExpr) contain(vg ValueGetter, val interface{}) bool {
	_, b := e.Set[val]
	return b
}

func (e *ArrayExpr) String() string {
	array := make([]string, len(e.Array))
	for idx := range array {
		array[idx] = fmt.Sprintf("%s", e.Array[idx])
	}
	return fmt.Sprintf("[%s]", strings.Join(array, ","))
}

type ContainerExpr interface {
	contain(vg ValueGetter, v interface{}) bool
}

type InExpr struct {
	left      Expression
	container ContainerExpr
	not       bool
}

func (e *InExpr) Eval(vg ValueGetter) (interface{}, error) {
	val, err := e.left.Eval(vg)
	if err != nil || val == nil {
		return nil, err
	}
	b := e.container.contain(vg, val)
	if e.not {
		return !b, nil
	}
	return b, nil
}

func (e *InExpr) String() string {
	return fmt.Sprintf("%s in %s", e.left.String(), e.String())
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
	*ast.BaseTExprParserVisitor
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
	regex := strings.TrimSpace(ctx.Regex().GetText())
	r := regexp.MustCompile(regex[1 : len(regex)-1])
	left := ctx.Expression().Accept(v)
	return &MatchExpr{left.(Expression), r}
}

func (v *ExprVisitor) VisitInExpression(ctx *ast.InExpressionContext) interface{} {
	left := ctx.Expression().Accept(v).(Expression)
	return &InExpr{
		left:      left.(Expression),
		container: ctx.Container().Accept(v).(ContainerExpr),
		not:       false,
	}
}

func (v *ExprVisitor) VisitNotInExpression(ctx *ast.NotInExpressionContext) interface{} {
	left := ctx.Expression().Accept(v).(Expression)
	return &InExpr{
		left:      left.(Expression),
		container: ctx.Container().Accept(v).(ContainerExpr),
		not:       true,
	}
}

func (v *ExprVisitor) VisitContainer(ctx *ast.ContainerContext) interface{} {
	if ctx.Array() != nil {
		return ctx.Array().Accept(v)
	}
	if ctx.Variable() != nil {
		return ctx.Variable().Accept(v)
	}
	if ctx.TString() != nil {
		c := ctx.TString().GetText()
		return NewValueExpression(c[1 : len(c)-1])
	}
	return v.VisitContainer(ctx)
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
	if ctx.TString() != nil {
		s := ctx.GetText()
		l := len(s)
		return NewValueExpression(s[1 : l-1])
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitCalc(ctx *ast.CalcContext) interface{} {
	return ctx.Bit().Accept(v)
}

func (v *ExprVisitor) VisitBit(ctx *ast.BitContext) interface{} {
	var left Expression
	if ctx.Bit() != nil {
		left = ctx.Bit().Accept(v).(Expression)
	}
	right := ctx.Shift().Accept(v).(Expression)
	if ctx.BAND() != nil {
		return NewBiCalcExpr(left, right, NewOP("&", band))
	}
	if ctx.BEOR() != nil {
		return NewBiCalcExpr(left, right, NewOP("^", beor))
	}
	if ctx.BIOR() != nil {
		return NewBiCalcExpr(left, right, NewOP("|", bior))
	}
	return right
}

func (v *ExprVisitor) VisitShift(ctx *ast.ShiftContext) interface{} {
	var left Expression
	if ctx.Shift() != nil {
		left = ctx.Shift().Accept(v).(Expression)
	}
	right := ctx.Plus().Accept(v).(Expression)
	if ctx.LSHIFT() != nil {
		return NewBiCalcExpr(left, right, NewOP("<<", lshift))
	}
	if ctx.RSHIFT() != nil {
		return NewBiCalcExpr(left, right, NewOP(">>", rshift))
	}
	return right
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
	var left Expression
	if ctx.Multiplying() != nil {
		left = ctx.Multiplying().Accept(v).(Expression)
	}
	right := ctx.Atom().Accept(v).(Expression)
	if ctx.DIV() != nil {
		return NewBiCalcExpr(left, right, NewOP("/", div))
	}
	if ctx.MUL() != nil {
		return NewBiCalcExpr(left, right, NewOP("*", mul))
	}
	if ctx.MOD() != nil {
		return NewBiCalcExpr(left, right, NewOP("%", mod))
	}
	return right
}

//
//func (v *ExprVisitor) VisitPow(ctx *ast.PowContext) interface{} {
//	if ctx.POW() != nil {
//		left := ctx.Atom(0).Accept(v).(Expression)
//		right := ctx.Atom(1).Accept(v).(Expression)
//		return NewBiCalcExpr(left, right, NewOP("^", pow))
//	}
//	return ctx.Atom(0).Accept(v)
//}

func (v *ExprVisitor) VisitScientific(ctx *ast.ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitFunction(ctx *ast.FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitFuncname(ctx *ast.FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitNumber(ctx *ast.NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitKind(ctx *ast.KindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitStrings(ctx *ast.StringsContext) interface{} {
	ss := make([]Expression, len(ctx.AllTString()))
	for idx, _ := range ss {
		s := ctx.TString(idx).GetText()
		l := len(s)
		ss[idx] = NewValueExpression(s[1 : l-1])
	}
	return NewArrayExpr(ss, reflect.TypeOf(""))
}

func (v *ExprVisitor) VisitIntegers(ctx *ast.IntegersContext) interface{} {
	ss := make([]Expression, len(ctx.AllInteger()))
	for idx, _ := range ss {
		s := ctx.Integer(idx).GetText()
		f, _ := strconv.ParseInt(s, 10, 64)
		ss[idx] = NewValueExpression(f)
	}
	return NewArrayExpr(ss, reflect.TypeOf(int64(0)))
}

func (v *ExprVisitor) VisitFloats(ctx *ast.FloatsContext) interface{} {
	ss := make([]Expression, len(ctx.AllFloat()))
	for idx, _ := range ss {
		s := ctx.Float(idx).GetText()
		f, _ := strconv.ParseFloat(s, 64)
		ss[idx] = NewValueExpression(f)
	}
	return NewArrayExpr(ss, reflect.TypeOf(float64(0)))
}

func (v *ExprVisitor) VisitBooleans(ctx *ast.BooleansContext) interface{} {
	ss := make([]Expression, len(ctx.AllBoolean()))
	for idx, _ := range ss {
		s := ctx.Boolean(idx).GetText()
		f, _ := strconv.ParseBool(s)
		ss[idx] = NewValueExpression(f)
	}
	return NewArrayExpr(ss, reflect.TypeOf(true))
}

func (v *ExprVisitor) VisitAtom(ctx *ast.AtomContext) interface{} {
	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		return ctx.Bit().Accept(v)
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

type TExprErrorListener struct {
	antlr.ErrorListener
}
