# TExpr


[![Build Status](https://travis-ci.org/aliyun/texpr.svg?branch=master)](https://travis-ci.org/aliyun/texpr)
[![Coverage status](https://img.shields.io/codecov/c/github/aliyun/texpr/master.svg)](https://codecov.io/github/aliyun/texpr)

Go语言布尔及数值运算表达式解析器

[Java语言版本](https://github.com/tauris-io/expression)

布尔表达式：

```
$country in ['cn', 'us', 'jp'] && $length > 100
$host == 'cs.qa.com' && 'Webkit' not in $ua
```

数值运算表达式:

```
( 1 ^ 5) + (2 << 3)
4 | 3 + 3
```



# Example

```
package main

import (
	"fmt"
	"log"

	"github.com/aliyun/texpr"
)

type context map[string]interface{}

func (vg context) Get(name string) interface{} {
	v, ok := vg[name]
	if ok {
		return v
	}
	return nil
}

func main() {
	m := context {
		"$value" : 15.0,
	}
	expr := texpr.MustCompile("$value - 10")
	v, err := expr.Eval(m)
	if err != nil {
		log.Fatalf("expression err, %s", err)
	}
	fmt.Println(v) // 5
}

```





# 表达式语法

## 标量变量

| 类型   | 范例               | 说明                                    |
| ------ | ------------------ | --------------------------------------- |
| 字符串 | 'hello' , "world"  | 字符串可以用单引号或双引号              |
| 整数   | 1,3,64             |                                         |
| 浮点数 | 1.1, 3.14          |                                         |
| 布尔值 | true/false         |                                         |
| 变量   | $name, @meta       | 变量名必须用$或@作为前缀.               |
| 数组   | ["hello", "world"] | 数组元素的类型可以是上面5种类型中的一种 |



## 数值计算

| 操作符 | 描述   | 例子        |
| ------ | ------ | ----------- |
| +      | 加     | 1+1=2       |
| -      | 减     | 2-1=1       |
| *      | 乘     | 2*2=4       |
| /      | 除     | 2/2=1       |
| %      | 取余   | 5/2=1       |
| &      | 位与   | 33&44=34    |
| \|     | 位或   | 33\|44=45   |
| ^      | 异或   | 33\|44=13   |
| <<     | 左位移 | 1 << 3 = 8  |
| >>     | 右位移 | 16 >> 3 = 2 |



## 布尔计算

### 关系运算符

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| ==    | 两值相等     | A==B    |
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
| in     | 判断左值是否被包含在右值中，右值为数组类型  | A in [A, B, C, D]       |
| in | 判断左值是否被包含在右值中，左值右值都是字符串或字符串类型变量 | A in B |

example

```
$name in ['John', 'Mary', "Jackson"]
'world' in 'hello world'
```



### 正则匹配

| 操作符 | 描述                                                         | 例子         |
| ------ | ------------------------------------------------------------ | ------------ |
| =~    | 判断左值是匹配右值，右值是正则表达式    | 'world' =~ /.+orl.+/       |

正则表达式首尾需要用'/'包含

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

范例

```
$name is string
$nil is null
'192.168.1.1' is ip4
```



## 函数运算

### 表达式函数

当前版本内置了一个函数'eval', 用来解析运算变量中的表达式, 范例:

```
// $expr = "$left + $right"
type context map[string]interface{}
func (vg context) Get(name string) interface{} {
	v, ok := vg[name]
	if ok {
		return v
	}
	return nil
}


m := context {
		"$left" : 2.0,
		"$right" : 3.0,
		"$ex" : "$left + $right"
}

expr, err := t.Compile("eval($ex)")

v, _ := expr.Eval(m)
// v == 5.0
```

### 自定义函数

可以在程序中自定义扩展函数, 只需要两个步骤

1. 定义一个结构体, 实现Function接口
2. 定义一个结构体变量并注册.

下面是自定义has_prefix函数的范例

```

type HasPrefixFunc struct {
	Function
}

func (f *HasPrefixFunc) Name() string {
	return "has_prefix"
}

func (f *HasPrefixFunc) Execute(vg ValueGetter, params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf(`HasPrefixFunc() takes exactly 2 argument (" + %d + " given)`, len(params))
	}
	return strings.HasPrefix(params[0].(string), params[1].(string)), nil
}

func main() {
  RegisterFunc(&HasPrefixFunc{})
	...
}

```



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
// 定义一个类型, 实现Get(string)interface{} 接口
type context map[string]interface{}
func (vg context) Get(name string) interface{} {
	v, ok := vg[name]
	if ok {
		return v
	}
	return nil
}

expr, err := t.Compile("($value - 10) / @meta")

m := context {
		"$value" : 15.0,
		"@meta" : 9.0,
}

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

