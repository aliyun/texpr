package main

import (
	"fmt"
	"log"

	"github.com/aliyun/texpr"
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

func main() {
	m := MakeVG("$value", 15.0)

	expr, err := texpr.Compile("$value - 10")

	if err != nil {
		log.Fatalf("expression err, %s", err)
	}
	v, err := expr.Eval(m)
	if err != nil {
		log.Fatalf("expression err, %s", err)
	}
	fmt.Println(v)
}
