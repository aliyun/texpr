package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/texpr"
)

func main() {
	expr, err := texpr.Compile(os.Args[1])

	if err != nil {
		log.Fatalf("expression err, %s", os.Args[1])
	}
	v, err := expr.Eval(nil)
	fmt.Println(v)
}
