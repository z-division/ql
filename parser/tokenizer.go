package parser

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/dropbox/godropbox/bufio2"
)

const (
	maxLookAheadSize = 256
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
		/*
		   '=':  // = vs ==
		   '<': // < vs <= vs <<
		   '>': // > vs >= vs >>
		   '&': // & vs &&
		   '|': // | vs ||
		   '!': // ! vs !=
		   '.': DOT vs float
		   '-': SUB vs number literal
		   '+': ADD vs number literal
		   '*': MUL vs STAR STAR ?
		*/
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

	currentLine   int
	currentColumn int

	// \n is context sensitive
	prevToken  int
	parenCount int

	reader *bufio2.LookAheadBuffer
}

func newTokenizer(filename string, reader io.Reader) (*tokenizer, error) {
	// TODO(patrick): strip leading whitespaces before running

	tok := &tokenizer{
		filename:      filename,
		currentLine:   1,
		currentColumn: 1,
		prevToken:     0,
		parenCount:    0,
		reader:        bufio2.NewLookAheadBuffer(reader, maxLookAheadSize),
	}

	err := tok.stripMeaninglessWhitespaces()
	if err != nil {
		return nil, err
	}

	return tok, nil
}

// This assumes that leading whitespace are stripped
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

	tok.prevToken = tokenType
	return tokenType, nil
}

func (tok *tokenizer) stripMeaninglessWhitespaces() error {
	for {
		bytes, err := tok.reader.Peek(1)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		switch bytes[0] {
		case ' ', '\t':
			err = tok.reader.Consume(1)
			if err != nil {
				return err
			}

			tok.currentColumn += 1
		case '\n':
			if tok.parenCount == 0 && newlineSensitive(tok.prevToken) {
				return nil
			}

			err = tok.reader.Consume(1)
			if err != nil {
				return err
			}

			tok.currentLine += 1
			tok.currentColumn = 1
		default:
			return nil
		}
	}
}

// Identifiers (/ keywords) are of the form [_a-zA-Z][_a-zA-Z0-9]*
func (tok *tokenizer) parseIdentifier(lval *qlSymType) (int, error) {
	bytes, err := tok.reader.Peek(1)
	if err != nil {
		return LEX_ERROR, err
	}

	if !(('a' <= bytes[0] && bytes[0] <= 'z') ||
		('A' <= bytes[0] && bytes[0] <= 'Z') ||
		'_' == bytes[0]) {

		// Not an identifier
		return 0, nil
	}

	bytes, err = tok.reader.PeekAll()
	if err != nil {
		if err == io.EOF {
			// Not enough bytes left to full the entire buffer.  just use
			// whatever is in the buffer
			bytes = tok.reader.Buffer()
		} else {
			return LEX_ERROR, err
		}
	}

	count := 1
	for count < len(bytes) {
		if ('a' <= bytes[count] && bytes[count] <= 'z') ||
			('A' <= bytes[count] && bytes[count] <= 'Z') ||
			('0' <= bytes[count] && bytes[count] <= '9') ||
			'_' == bytes[count] {

			count += 1
		} else {
			lval.StartLine = tok.currentLine
			lval.StartColumn = tok.currentColumn

			lval.strVal = string(bytes[:count])

			err = tok.reader.Consume(count)
			if err != nil {
				return 0, err
			}
			tok.currentColumn += count

			lval.EndLine = tok.currentLine
			lval.EndColumn = tok.currentColumn

			return IDENTIFIER, nil
		}
	}

	// XXX(patrick): maybe handle this correctly ...
	return LEX_ERROR, fmt.Errorf(
		"%s:%s:%s: identifier is too long",
		filepath.Base(tok.filename),
		tok.currentLine,
		tok.currentColumn)
}

func (tok *tokenizer) parseSingleCharToken(lval *qlSymType) (int, error) {
	bytes, err := tok.reader.Peek(1)
	if err != nil {
		return LEX_ERROR, err
	}

	token, ok := singleCharTokens[bytes[0]]
	if !ok {
		return 0, nil
	}

	lval.StartLine = tok.currentLine
	lval.StartColumn = tok.currentColumn

	err = tok.reader.Consume(1)
	if err != nil {
		return LEX_ERROR, err
	}
	tok.currentColumn += 1

	lval.EndLine = tok.currentLine
	lval.EndColumn = tok.currentColumn

	return token, nil
}

func (tok *tokenizer) parseNextToken(lval *qlSymType) (int, error) {
	bytes, err := tok.reader.Peek(1)
	if err != nil {
		if err == io.EOF {
			return 0, nil
		}
		return 0, err
	}

	token, err := tok.parseIdentifier(lval)
	if err != nil {
		return token, err
	}

	if token == IDENTIFIER {
		// The identifier may be a keyword
		kwToken, ok := keywords[strings.ToLower(lval.strVal)]
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
		"%s:%d:%d: unknown character: %c",
		filepath.Base(tok.filename),
		tok.currentLine,
		tok.currentColumn,
		bytes[0])
}
