package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/z-division/ql/parser"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	_, err = parser.Parse(bytes.NewReader(content))
	if err != nil {
		panic(err)
	}

	fmt.Println("OK")
}
