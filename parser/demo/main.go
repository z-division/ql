package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/z-division/ql/parser"
)

func main() {
	for _, path := range os.Args[1:] {
		fmt.Println(path, "===========================")

		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		parsed, err := parser.Parse(path, bytes.NewReader(content))
		if err != nil {
			panic(err)
		}

		for _, node := range parsed {
			fmt.Println(node)
		}
		fmt.Println("OK")
	}
}
