package parser

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

const (
	chunkSize = 1024
)

var (
	// tokens that can be determined by a single char look ahead
	singleCharTokens = map[byte]int{
		'{':  L_BRACE,
		'}':  R_BRACE,
		'(':  L_PAREN,
		')':  R_PAREN,
		';':  SEMICOLON,
		'\n': NEWLINE,
		',':  COMMA,
		'/':  DIV, // don't have to worry about comments
		'%':  MOD,

		// TODO(patrick): deal with ambiguous symbols
		// '=':  // = vs ==
		// '<': // < vs <= vs <<
		// '>': // > vs >= vs >>
		// '&': // & vs &&
		// '|': // | vs ||
		// '!': // ! vs !=
		// '.': DOT vs float
		// '-': SUB vs number literal
		// '+': ADD vs number literal
		// '*': MUL vs STAR STAR ?
		'=': ASSIGN,
		'<': LT,
		'>': GT,
		'&': AND,
		'|': OR,
		'!': NOT,
		'.': DOT,
		'-': SUB,
		'+': ADD,
		'*': MUL,
	}

	// keywords are case insensitive
	keywords = map[string]int{
		"let": LET,
	}
)

func newlineSensitive(token int) bool {
	return false
}

type tokenizer struct {
	filename string
	reader   io.Reader

	Position

	// \n is context sensitive
	prevTokenType int
	prevToken     *Token
	parenCount    int

	buffered []byte
}

func newTokenizer(filename string, reader io.Reader) (*tokenizer, error) {
	return &tokenizer{
		filename: filename,
		reader:   reader,
		Position: Position{
			Column: 1,
			Line:   1,
		},
	}, nil
}

func (tok *tokenizer) fillN(n int) error {
	if n < 1 {
		panic(fmt.Sprintf("Invalid peek: %d", n))
	}

	for len(tok.buffered) < n {
		buffer := make([]byte, chunkSize)
		n, err := tok.reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		tok.buffered = append(tok.buffered, buffer[:n]...)
	}

	if len(tok.buffered) < n {
		return io.EOF
	}

	return nil
}

func (tok *tokenizer) peekN(n int) ([]byte, error) {
	err := tok.fillN(n)
	if err != nil {
		return nil, err
	}

	return tok.buffered[:n], nil
}

func (tok *tokenizer) peek() (byte, error) {
	err := tok.fillN(1)
	if err != nil {
		return 0, err
	}

	return tok.buffered[0], nil
}

func (tok *tokenizer) consumeN(n int) {
	if len(tok.buffered) < n {
		panic(fmt.Sprintf("Invalid consume: %d", n))
	}

	for _, char := range tok.buffered[:n] {
		if char == '(' {
			tok.parenCount += 1
		} else if char == ')' {
			tok.parenCount -= 1
		}

		if char == '\n' {
			tok.Column = 1
			tok.Line += 1
		} else {
			tok.Column += 1
		}
	}

	tok.buffered = tok.buffered[n:]
}

func (tok *tokenizer) consume() {
	tok.consumeN(1)
}

func (tok *tokenizer) nextToken(lval *qlSymType) (int, error) {
	// TODO(patrick): parse leading comment

	tokenType, err := tok.parseNextToken(lval)
	if err != nil {
		return 0, err
	}

	// TODO(patrick): parse trailing comment

	err = tok.stripMeaninglessWhitespaces()
	if err != nil {
		return 0, err
	}

	tok.prevTokenType = tokenType
	tok.prevToken = lval.Token

	return tokenType, nil
}

func (tok *tokenizer) stripMeaninglessWhitespaces() error {
	for {
		char, err := tok.peek()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		switch char {
		case ' ', '\t':
			tok.consume()
		case '\n':
			if tok.parenCount == 0 && newlineSensitive(tok.prevTokenType) {
				return nil
			}

			tok.consume()
		default:
			return nil
		}
	}
}

// Identifiers (/ keywords) are of the form [_a-zA-Z][_a-zA-Z0-9]*
func (tok *tokenizer) parseIdentifier(lval *qlSymType) (int, error) {
	var value []byte

	start := tok.Position
	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return LEX_ERROR, err
		}

		if !(('a' <= char && char <= 'z') ||
			('A' <= char && char <= 'Z') ||
			'_' == char) {

			break
		}

		value = append(value, char)
		tok.consume()
	}

	if len(value) == 0 {
		// Not an identifier
		return 0, nil
	}

	lval.Token = &Token{
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}

	return IDENTIFIER, nil
}

func (tok *tokenizer) parseSingleCharToken(lval *qlSymType) (int, error) {
	char, err := tok.peek()
	if err != nil {
		return LEX_ERROR, err
	}

	token, ok := singleCharTokens[char]
	if !ok {
		return 0, nil
	}

	start := tok.Position

	lval.Token = &Token{
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(char),
	}

	tok.consume()
	return token, nil
}

func (tok *tokenizer) parseNextToken(lval *qlSymType) (int, error) {
	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return 0, nil
		}
		return LEX_ERROR, err
	}

	token, err := tok.parseIdentifier(lval)
	if err != nil {
		return LEX_ERROR, err
	}

	if token == IDENTIFIER {
		// The identifier may be a keyword
		kwToken, ok := keywords[strings.ToLower(lval.Token.Value)]
		if ok {
			return kwToken, nil
		}

		return token, nil
	}

	token, err = tok.parseSingleCharToken(lval)
	if token != 0 || err != nil {
		return token, err
	}

	return LEX_ERROR, fmt.Errorf(
		"%s:%v: unknown character: %c",
		filepath.Base(tok.filename),
		tok.Position,
		char)
}
