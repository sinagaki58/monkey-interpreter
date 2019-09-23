package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sinagaki58/monkey-interpreter/monkey/lexer"
	"github.com/sinagaki58/monkey-interpreter/monkey/parser"
)

func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		input := scanner.Text()
		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
