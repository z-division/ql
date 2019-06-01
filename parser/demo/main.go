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

	parsed, err := parser.Parse("file", bytes.NewReader(content))
	if err != nil {
		panic(err)
	}

	fmt.Println("OK")
	for _, node := range parsed {
		fmt.Println(node)
	}
}
