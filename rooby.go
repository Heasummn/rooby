package main

import (
	"flag"
	"fmt"
	"github.com/rooby-lang/rooby/ast"
	"github.com/rooby-lang/rooby/bytecode"
	"github.com/rooby-lang/rooby/lexer"
	"github.com/rooby-lang/rooby/parser"
	"github.com/rooby-lang/rooby/vm"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	compileOptionPtr := flag.Bool("c", false, "Compile to bytecode")

	flag.Parse()

	filepath := flag.Arg(0)

	var fileExt string
	dir, filename := path.Split(filepath)
	splitedFN := strings.Split(filename, ".")

	if len(splitedFN) <= 1 {
		fmt.Printf("Only support eval/compile single file now.")
		return
	}

	filename = splitedFN[0]
	fileExt = splitedFN[1]

	file, err := ioutil.ReadFile(filepath)
	check(err)

	switch fileExt {
	case "ro":
		program := buildAST(file)

		g := bytecode.NewGenerator(program)
		bytecodes := g.GenerateByteCode(program)

		if !*compileOptionPtr {
			execBytecode(bytecodes)
			return
		}

		writeByteCode(bytecodes, dir, filename)

	case "robc":
		bytecodes := string(file)
		execBytecode(bytecodes)
	default:
		fmt.Printf("Unknown file extension: %s", fileExt)
	}
}

func writeByteCode(bytecodes, dir, filename string) {
	f, err := os.Create(dir + filename + ".robc")

	if err != nil {
		panic(err)
	}

	f.WriteString(bytecodes)
}

func execBytecode(bytecodes string) {
	v := vm.New()
	v.ExecBytecodes(bytecodes)
}

func buildAST(file []byte) *ast.Program {
	input := string(file)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	p.CheckErrors()

	return program
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
