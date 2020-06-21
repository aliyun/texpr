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
