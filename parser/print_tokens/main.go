package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/z-division/ql/parser"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tok, err := parser.NewRawTokenizer(os.Args[1], bytes.NewReader(content))
	if err != nil {
		panic(err)
	}

	tok, err = parser.NewTerminatorProcessor(os.Args[1], tok)
	if err != nil {
		panic(err)
	}

	for {
		token, err := tok.Next()
		if err == io.EOF {
			fmt.Println("<EOF>")
			return
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(token)
	}
}
