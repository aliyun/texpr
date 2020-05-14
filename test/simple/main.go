package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/tauris-expression-engine/compiler"
)

func main() {
	expr, err := compiler.Compile(os.Args[1])

	if err != nil {
		log.Fatalf("expression err, %s", os.Args[1])
	}
	v, err := expr.Eval(nil)
	fmt.Println(v)
}
