package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func main() {
	// Common Expression Language
	d := cel.Declarations(decls.NewVar("name", decls.String))
	env, err := cel.NewEnv(d)
	if err != nil {
		log.Fatalf("environment creation error: %v\n", err)
	}
	ast, iss := env.Compile(`"Hello world! I'm " + name + "."`)
	// Check iss for compilation errors.
	if iss.Err() != nil {
		log.Fatalln(iss.Err())
	}
	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalln(err)
	}
	out, _, err := prg.Eval(map[string]interface{}{
		"name": "CEL",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
	// Output:Hello world! I'm CEL.
}
