// A standard (mostly) regex based tokenizer.

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

type parseTokenFunc func(tok *rawTokenizer, lval *qlSymType) (int, error)

var (
	// leading character -> parse function
	leadingCharParseTokenFunc = map[byte]parseTokenFunc{
		'a': parseIdentifierOrKeyword,
		'b': parseIdentifierOrKeyword,
		'c': parseIdentifierOrKeyword,
		'd': parseIdentifierOrKeyword,
		'e': parseIdentifierOrKeyword,
		'f': parseIdentifierOrKeyword,
		'g': parseIdentifierOrKeyword,
		'h': parseIdentifierOrKeyword,
		'i': parseIdentifierOrKeyword,
		'j': parseIdentifierOrKeyword,
		'k': parseIdentifierOrKeyword,
		'l': parseIdentifierOrKeyword,
		'm': parseIdentifierOrKeyword,
		'n': parseIdentifierOrKeyword,
		'o': parseIdentifierOrKeyword,
		'p': parseIdentifierOrKeyword,
		'q': parseIdentifierOrKeyword,
		'r': parseIdentifierOrKeyword,
		's': parseIdentifierOrKeyword,
		't': parseIdentifierOrKeyword,
		'u': parseIdentifierOrKeyword,
		'v': parseIdentifierOrKeyword,
		'w': parseIdentifierOrKeyword,
		'x': parseIdentifierOrKeyword,
		'y': parseIdentifierOrKeyword,
		'z': parseIdentifierOrKeyword,
		'A': parseIdentifierOrKeyword,
		'B': parseIdentifierOrKeyword,
		'C': parseIdentifierOrKeyword,
		'D': parseIdentifierOrKeyword,
		'E': parseIdentifierOrKeyword,
		'F': parseIdentifierOrKeyword,
		'G': parseIdentifierOrKeyword,
		'H': parseIdentifierOrKeyword,
		'I': parseIdentifierOrKeyword,
		'J': parseIdentifierOrKeyword,
		'K': parseIdentifierOrKeyword,
		'L': parseIdentifierOrKeyword,
		'M': parseIdentifierOrKeyword,
		'N': parseIdentifierOrKeyword,
		'O': parseIdentifierOrKeyword,
		'P': parseIdentifierOrKeyword,
		'Q': parseIdentifierOrKeyword,
		'R': parseIdentifierOrKeyword,
		'S': parseIdentifierOrKeyword,
		'T': parseIdentifierOrKeyword,
		'U': parseIdentifierOrKeyword,
		'V': parseIdentifierOrKeyword,
		'W': parseIdentifierOrKeyword,
		'X': parseIdentifierOrKeyword,
		'Y': parseIdentifierOrKeyword,
		'Z': parseIdentifierOrKeyword,
		'_': parseIdentifierOrKeyword,

		'1': parseNumber,
		'2': parseNumber,
		'3': parseNumber,
		'4': parseNumber,
		'5': parseNumber,
		'6': parseNumber,
		'7': parseNumber,
		'8': parseNumber,
		'9': parseNumber,
		'0': parseNumber,

		'.': parseDot, // Could be a dot or a float

		'/': parseSlash, // comments or divide

		'\'': parseChar,

		'"': parseString,

		'=': parseSymbol(map[string]int{"=": ASSIGN, "==": EQ}),
		'!': parseSymbol(map[string]int{"!": NOT, "!=": NE}),
		'<': parseSymbol(map[string]int{"<": LT, "<=": LE, "<<": L_SHIFT}),
		'>': parseSymbol(map[string]int{">": GT, ">=": GE, ">>": R_SHIFT}),

		'&': parseSymbol(map[string]int{"&": BITWISE_AND, "&&": AND}),
		'|': parseSymbol(map[string]int{"|": BITWISE_OR, "||": OR}),
		'^': parseSymbol(map[string]int{"^": XOR}),

		'+': parseSymbol(map[string]int{"+": ADD}),
		'-': parseSymbol(map[string]int{"-": SUB}),
		'*': parseSymbol(map[string]int{"*": MUL, "**": STAR_STAR}),
		'%': parseSymbol(map[string]int{"%": MOD}),

		',': parseSymbol(map[string]int{",": COMMA}),

		'\n': parseSymbol(map[string]int{"\n": NEWLINE}),

		':': parseSymbol(map[string]int{":": COLON}),
		';': parseSymbol(map[string]int{";": SEMICOLON}),

		'(': parseSymbol(map[string]int{"(": L_PAREN}),
		')': parseSymbol(map[string]int{")": R_PAREN}),
		'{': parseSymbol(map[string]int{"{": L_BRACE}),
		'}': parseSymbol(map[string]int{"}": R_BRACE}),
		'[': parseSymbol(map[string]int{"[": L_BRACKET}),
		']': parseSymbol(map[string]int{"]": R_BRACKET}),

		// Unused symbols: @ # $ ~ ` ? \

		// NOTE(patrick): ' ' and '\t' are ignored
	}

	// keywords are case insensitive
	keywords = map[string]int{
		"let": LET,
	}
)

func parseIdentifierOrKeyword(tok *rawTokenizer, lval *qlSymType) (int, error) {
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
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: invalid identifier",
			filepath.Base(tok.filename),
			tok.Position)
	}

	lval.Token = &Token{
		Type: IDENTIFIER,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}

	// The identifier may be a keyword
	kwToken, ok := keywords[strings.ToLower(lval.Token.Value)]
	if ok {
		lval.Token.Type = kwToken
		return kwToken, nil
	}

	return IDENTIFIER, nil
}

func parseNumber(tok *rawTokenizer, lval *qlSymType) (int, error) {
	return LEX_ERROR, fmt.Errorf(
		"%s:%v: TODO(patrick): implement number parser",
		filepath.Base(tok.filename),
		tok.Position)
}

func parseDot(tok *rawTokenizer, lval *qlSymType) (int, error) {
	return LEX_ERROR, fmt.Errorf(
		"%s:%v: TODO(patrick): implement dot parser",
		filepath.Base(tok.filename),
		tok.Position)
}

func parseChar(tok *rawTokenizer, lval *qlSymType) (int, error) {
	return LEX_ERROR, fmt.Errorf(
		"%s:%v: TODO(patrick): implement char parser",
		filepath.Base(tok.filename),
		tok.Position)
}

func parseString(tok *rawTokenizer, lval *qlSymType) (int, error) {
	return LEX_ERROR, fmt.Errorf(
		"%s:%v: TODO(patrick): implement string parser",
		filepath.Base(tok.filename),
		tok.Position)
}

func parseSlash(tok *rawTokenizer, lval *qlSymType) (int, error) {
	return LEX_ERROR, fmt.Errorf(
		"%s:%v: TODO(patrick): implement slash parser",
		filepath.Base(tok.filename),
		tok.Position)
}

func parseSymbol(symbolTokens map[string]int) parseTokenFunc {
	return func(tok *rawTokenizer, lval *qlSymType) (int, error) {
		matchedSymbol := ""
		matchedTokenType := 0
		for symbol, tokenType := range symbolTokens {
			if len(matchedSymbol) >= len(symbol) {
				continue
			}

			bytes, err := tok.peekN(len(symbol))
			if err != nil {
				if err != io.EOF {
					return LEX_ERROR, err
				}

				continue
			}

			if string(bytes) == symbol {
				matchedSymbol = symbol
				matchedTokenType = tokenType
			}
		}

		if matchedSymbol == "" {
			return LEX_ERROR, fmt.Errorf(
				"%s:%v: invalid symbol",
				filepath.Base(tok.filename),
				tok.Position)
		}

		start := tok.Position
		tok.consumeN(len(matchedSymbol))

		lval.Token = &Token{
			Type: matchedTokenType,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: matchedSymbol,
		}

		return matchedTokenType, nil
	}
}

type rawTokenizer struct {
	filename string
	reader   io.Reader

	Position

	// \n is context sensitive
	prevTokenType int
	prevToken     *Token
	parenCount    int

	buffered []byte
}

func newRawTokenizer(filename string, reader io.Reader) (*rawTokenizer, error) {
	return &rawTokenizer{
		filename: filename,
		reader:   reader,
		Position: Position{
			Column: 1,
			Line:   1,
		},
	}, nil
}

func (tok *rawTokenizer) fillN(n int) error {
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

func (tok *rawTokenizer) peekN(n int) ([]byte, error) {
	err := tok.fillN(n)
	if err != nil {
		return nil, err
	}

	return tok.buffered[:n], nil
}

func (tok *rawTokenizer) peek() (byte, error) {
	err := tok.fillN(1)
	if err != nil {
		return 0, err
	}

	return tok.buffered[0], nil
}

func (tok *rawTokenizer) consumeN(n int) {
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

func (tok *rawTokenizer) consume() {
	tok.consumeN(1)
}

func (tok *rawTokenizer) nextToken(lval *qlSymType) (int, error) {
	tokenType, err := tok.parseNextToken(lval)
	if err != nil {
		return 0, err
	}

	err = tok.stripMeaninglessWhitespaces()
	if err != nil {
		return 0, err
	}

	// TODO(patrick): remove this hack.  Newlines are meaningful.
	if tokenType == NEWLINE {
		return tok.nextToken(lval)
	}

	tok.prevTokenType = tokenType
	tok.prevToken = lval.Token

	return tokenType, nil
}

func (tok *rawTokenizer) stripMeaninglessWhitespaces() error {
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
		default:
			return nil
		}
	}
}

func (tok *rawTokenizer) parseNextToken(lval *qlSymType) (int, error) {
	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return 0, nil
		}
		return LEX_ERROR, err
	}

	parseToken, ok := leadingCharParseTokenFunc[char]
	if !ok {
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: unknown character: %c",
			filepath.Base(tok.filename),
			tok.Position,
			char)
	}

	return parseToken(tok, lval)
}
