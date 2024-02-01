package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
	"os"
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
        io.WriteString(out, program.String())
        io.WriteString(out, "\n")
	}
}

func printParserErrors(errors []string) {
    for _, msg := range errors {
        io.WriteString(os.Stderr, "\t" + msg + "\n")
    }
}
