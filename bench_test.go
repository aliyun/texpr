package texpr

import (
	"testing"

)

type Entity struct {
	m map[string]interface{}
}

func (e Entity) Get(name string) interface{} {
	v, ok := e.m[name]
	if ok {
		return v
	}
	return nil
}

func Benchmark_expr(b *testing.B) {
	params := make(map[string]interface{})
	params["$Origin"] = "MOW"
	params["$Country"] = "RU"
	params["$Adults"] = 1.0
	params["$Value"] = 100.0
	entity := &Entity{params}

	program, err := Compile(`($Origin == "MOW" || $Country == "RU") && ($Value >= 100 || $Adults == 1)`)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = program.Eval(entity)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
