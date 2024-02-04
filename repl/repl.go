package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
    
    "monkey/eval"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(p.Errors())
            continue
        }

        evaluated := eval.Eval(program)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func printParserErrors(errors []string) {
    io.WriteString(os.Stderr, " Parser errors:\n")
    for _, msg := range errors {
        io.WriteString(os.Stderr, "\t" + msg + "\n")
    }
}
