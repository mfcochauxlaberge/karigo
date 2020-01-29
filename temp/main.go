package main

import (
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	_ "github.com/mfcochauxlaberge/karigo"
)

const src = `package foo

import (
	"strings"
	"bytes"
	"github.com/mfcochauxlaberge/karigo"
)

func Action(cp *karigo.Checkpoint) {
	cp.Apply([]karigo.Op{
		karigo.NewOpSet("0_meta", "action", "value", "It works!"),
	})
}

func Bar(s string) string {
	return strings.ToUpper(s+"-Foo!")
}
`

func main() {
	i := interp.New(interp.Options{
		GoPath: "/home/mfcl/Go",
	})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	println(r)
}
