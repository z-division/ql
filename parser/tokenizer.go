package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type tokenizer struct {
	tokens []string
}

func newTokenizer(reader io.Reader) (*tokenizer, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return &tokenizer{
		tokens: strings.Split(
			strings.ToLower(strings.Trim(string(content), "\n")),
			" "),
	}, nil
}

func (tok *tokenizer) nextToken(lval *qlSymType) (int, error) {
	if len(tok.tokens) == 0 {
		return 0, nil
	}

	head := tok.tokens[0]
	tok.tokens = tok.tokens[1:]
	switch head {
	case "select":
		return SELECT, nil
	case "from":
		return FROM, nil
	case "where":
		return WHERE, nil
	default:
		return LEX_ERROR, fmt.Errorf("Unknown token: %s", head)
	}
}
