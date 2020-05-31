# TExpr 
Go语言布尔及数值运算表达式解析器

[![Build Status](https://travis-ci.org/aliyun/texpr.svg?branch=master)](https://travis-ci.org/aliyun/texpr)
[![Coverage status](https://img.shields.io/codecov/c/github/aliyun/texpr/master.svg)](https://codecov.io/github/aliyun/texpr)

[source](https://github.com/aliyun/texpr)

## Quick Start

```bash
git clone git@github.com:aliyun/texpr.git
cd texpr
go run test/simple/main.go 1+1
go run test/simple/main.go "(100 * 11) / 25"
go run test/simple/main.go "99 in [99, 88, 77]"
go run test/simple/main.go "99 > 100"
go run test/simple/main.go "33 > 22 && 22 > 11"
go run test/simple/main.go "'world' =~ /.+orl.+/"
```



## 数值计算

| 操作符 | 描述 | 例子   |
| ------ | ---- | ------ |
| +      | 加   | 1+1=2  |
| -      | 减   | 2-1=1  |
| *      | 乘   | 2*2=4  |
| /      | 除   | 2/2=1  |
| %      | 取余 | 5/2=1  |
| ^      | 幂   | 2^4=16 |

*不支持位运算*

## 布尔计算

### 关系运算符

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| =     | 两值相等     | A=B     |
| >   | 左值大于右值 | A>B  |
| <    | 右值大于左值 | A<B |
| >=     | 左值大于等于右值     | A>=B    |
| <=   | 右值大于等于左值 | A<=B |

*左值和右值必须为同类型*


### 逻辑运算符

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| &&     | 称为逻辑与运算符。当且仅当两个操作数都为真，条件才为真。     | A && B       |
| \|\|   | 称为逻辑或操作符。如果任何两个操作数任何一个为真，条件为真。 | A \|\| B     |
| not    | 称为逻辑非运算符。用来反转操作数的逻辑状态。如果条件为true，则逻辑非运算符将得到false。 | not（A && B) |

### 包含表达式

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| in     | 判断左值是否被在右值中，右值需要为数组类型    | A in [A, B, C, D]       |



### 正则匹配

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| =~    | 判断左值是匹配右值，右值是正则表达式    | 'world' =~ /.+orl.+/       |



### IS 判断

| 操作符 | 描述                         | 例子       |
| ------ | ---------------------------- | ---------- |
| is     | 判断左值是否是右值代表的类型 | A is empty |
| is not   | 判断左值是否*不是*右值代表的类型 | A is not null |

支持的类型如下

| 名称 | 描述                         | 例子       |
| ------ | ---------------------------- | ---------- |
| string    | 字符串 | A is string |
| integer   | 整数 | A is integer |
| float   | 浮点数 | A is float |
| boolean   | 布尔值 | A is boolean && not A |
| host    | 主机名 | A is host |
| ip4    | ipv4地址 | A is ip4 |
| empty    | null、空字符、空数组 | A is not empty |
| null    | null | A is not null |


## Development

### 简单表达式计算

```go
import (
	"fmt"
	"log"
	"os"

	t "github.com/aliyun/texpr"
)

func main() {
    input := "99 in "33 > 22 && 22 > 11"
	expr, err := t.Compile(input)

	if err != nil {
		log.Fatalf("expression err, %s", os.Args[1])
	}
	v, err := expr.Eval(nil)
	fmt.Println(v)
}
```

### 带有变量的表达式计算

表达式中可以带有变量，表达式中的变量必须以'$'或者'@'开头，如：

```
expr, err := t.Compile("($value - 10) / @meta")
m := <make an object>
v, err := expr.Eval(m)
if err != nil {
log.Fatalf("expression err, %s", err)
}
```

其中  **expr.Eval()** 的参数m，必须实现下面的interface

```
type ValueGetter interface {
	Get(name string) interface{}
}
```

在执行时，解释器会调用参数的Get方法取得表达式中变量的值，详细实现可以参考范例 test/variable/main.go

