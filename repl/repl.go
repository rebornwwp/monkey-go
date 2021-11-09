package repl

import (
	"bufio"
	"fmt"
	"io"

	"example.com/m/lexer"
	"example.com/m/token"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	for {
		fmt.Fprintf(w, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			continue
		}

		l := lexer.New(scanner.Text())

		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tk)
		}
	}

}
