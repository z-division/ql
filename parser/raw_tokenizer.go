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

type parseTokenFunc func(tok *rawTokenizer) (*Token, error)

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

		// Unused leading characters: @ # $ ~ ` ? \

		// NOTE(patrick): ' ' and '\t' are ignored
	}

	// keywords are case insensitive
	keywords = map[string]int{
		"let": LET,
	}

	// NOTE(patrick): This is a subset of the standard C escape sequence.
	// https://en.wikipedia.org/wiki/Escape_sequences_in_C
	validEscapeChars = map[byte]struct{}{
		'n':  struct{}{}, // \n
		't':  struct{}{}, // \t
		'\\': struct{}{}, // \\
		'\'': struct{}{}, // \'
		'"':  struct{}{}, // \"
	}

	// TODO(patrick): perform stricter check.
	invalidLiteralChars = map[byte]struct{}{
		'\a': struct{}{},
		'\b': struct{}{},
		'\f': struct{}{},
		'\n': struct{}{},
		'\r': struct{}{},
		'\t': struct{}{},
		'\v': struct{}{},
	}
)

func parseIdentifierOrKeyword(tok *rawTokenizer) (*Token, error) {
	var value []byte

	start := tok.Position
	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
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
		return nil, fmt.Errorf(
			"%s:%v: invalid identifier",
			filepath.Base(tok.filename),
			tok.Position)
	}

	token := &Token{
		Type: IDENT,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}

	// The identifier may be a keyword
	kwToken, ok := keywords[strings.ToLower(token.Value)]
	if ok {
		token.Type = kwToken
	}

	return token, nil
}

// return (nil, nil) if it doesn't match.
func maybeParseDecimal(tok *rawTokenizer) ([]byte, error) {
	var value []byte
	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if !('0' <= char && char <= '9') {
			break
		}

		value = append(value, char)
		tok.consume()
	}

	return value, nil
}

// return (nil, nil) if it doesn't match.
func maybeParseFloatFractional(tok *rawTokenizer) ([]byte, error) {
	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	if char != '.' {
		return nil, nil
	}

	tok.consume()

	value, err := maybeParseDecimal(tok)
	if err != nil {
		return nil, err
	}

	return append([]byte{'.'}, value...), nil
}

// return (nil, nil) if it doesn't match.
func maybeParseFloatExponent(tok *rawTokenizer) ([]byte, error) {
	bytes, err := tok.peekN(2)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	if bytes[0] != 'e' && bytes[0] != 'E' {
		return nil, nil
	}

	if bytes[1] == '+' || bytes[1] == '-' {
		bytes, err = tok.peekN(3)
		if err != nil {
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
	}

	dig := bytes[len(bytes)-1]
	if !('0' <= dig && dig <= '9') {
		return nil, nil
	}

	tok.consumeN(len(bytes) - 1)

	value, err := maybeParseDecimal(tok)
	if err != nil {
		return nil, err
	}

	if len(value) == 0 {
		return nil, fmt.Errorf(
			"%s:%v: invaild exponent",
			filepath.Base(tok.filename),
			tok.Position)
	}

	return append(bytes[:len(bytes)-1], value...), nil
}

func parseHex(tok *rawTokenizer) (*Token, error) {
	start := tok.Position

	value, err := tok.peekN(2)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf(
				"%s:%v: invaild hexadecimal",
				filepath.Base(tok.filename),
				tok.Position)
		}

		return nil, err
	}

	if !(value[0] == '0' &&
		(value[1] == 'x' || value[1] == 'X')) {

		return nil, fmt.Errorf(
			"%s:%v: invaild hexadecimal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	tok.consumeN(2)

	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if ('0' <= char && char <= '9') ||
			('a' <= char && char <= 'f') ||
			('A' <= char && char <= 'F') {

			value = append(value, char)
			tok.consume()
		} else {
			break
		}
	}

	if len(value) == 2 {
		return nil, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
	}

	return &Token{
		Type: INT,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}, nil
}

func parseOct(tok *rawTokenizer) (*Token, error) {
	start := tok.Position

	var value []byte
	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if !('0' <= char && char <= '7') {
			break
		}

		value = append(value, char)
		tok.consume()
	}

	if len(value) == 0 {
		return nil, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
	}

	return &Token{
		Type: INT,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}, nil
}

func parseNumber(tok *rawTokenizer) (*Token, error) {
	bytes, err := tok.peekN(2)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if err != io.EOF && bytes[0] == '0' {
		if bytes[1] == 'x' || bytes[1] == 'X' { // hex
			return parseHex(tok)
		} else { // oct
			return parseOct(tok)
		}
	}

	start := tok.Position

	value, err := maybeParseDecimal(tok)
	if err != nil {
		return nil, err
	}

	if len(value) == 0 {
		return nil, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
	}

	fractional, err := maybeParseFloatFractional(tok)
	if err != nil {
		return nil, err
	}

	exponent, err := maybeParseFloatExponent(tok)
	if err != nil {
		return nil, err
	}

	numType := INT
	if len(fractional) > 0 || len(exponent) > 0 {
		numType = FLOAT
		value = append(value, fractional...)
		value = append(value, exponent...)
	}

	return &Token{
		Type: numType,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}, nil
}

func parseDot(tok *rawTokenizer) (*Token, error) {
	start := tok.Position

	fractional, err := maybeParseFloatFractional(tok)
	if err != nil {
		return nil, err
	}

	if len(fractional) == 0 {
		return nil, fmt.Errorf(
			"%s:%v: invaild symbol",
			filepath.Base(tok.filename),
			tok.Position)
	} else if len(fractional) == 1 { // single dot
		return &Token{
			Type: DOT,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: ".",
		}, nil
	}

	// dot follow by digit is a float
	exponent, err := maybeParseFloatExponent(tok)
	if err != nil {
		return nil, err
	}

	return &Token{
		Type: FLOAT,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(append(fractional, exponent...)),
	}, nil
}

func parseChar(tok *rawTokenizer) (*Token, error) {
	value, err := tok.peekN(3)
	if err != nil || value[0] != '\'' {
		return nil, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	// '' is invalid
	if value[1] == '\'' {
		return nil, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	start := tok.Position

	if value[1] == '\\' {
		_, ok := validEscapeChars[value[2]]
		if !ok {
			return nil, fmt.Errorf(
				"%s:%v: invalid escaped character",
				filepath.Base(tok.filename),
				tok.Position)
		}

		value, err = tok.peekN(4)
		if err != nil || value[3] != '\'' {
			return nil, fmt.Errorf(
				"%s:%v: invalid char literal",
				filepath.Base(tok.filename),
				tok.Position)
		}

		tok.consumeN(4)
		return &Token{
			Type: CHAR,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: string(value),
		}, nil
	}

	_, ok := invalidLiteralChars[value[1]]
	if ok || value[2] != '\'' {
		return nil, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	tok.consumeN(3)
	return &Token{
		Type: CHAR,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}, nil
}

func parseString(tok *rawTokenizer) (*Token, error) {
	start := tok.Position

	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf(
				"%s:%v: invalid string literal",
				filepath.Base(tok.filename),
				tok.Position)
		}
		return nil, err
	}

	if char != '"' {
		return nil, fmt.Errorf(
			"%s:%v: invalid string literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	tok.consume()

	value := []byte{'"'}
	escaped := false
	for {
		char, err = tok.peek()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf(
					"%s:%v: invalid string literal",
					filepath.Base(tok.filename),
					tok.Position)
			}

			return nil, err
		}

		_, ok := invalidLiteralChars[char]
		if ok {
			return nil, fmt.Errorf(
				"%s:%v: invalid string literal",
				filepath.Base(tok.filename),
				tok.Position)
		}

		end := false
		if escaped {
			_, ok := validEscapeChars[char]
			if !ok {
				return nil, fmt.Errorf(
					"%s:%v: invalid escaped character",
					filepath.Base(tok.filename),
					tok.Position)
			}

			escaped = false
		} else if char == '"' {
			end = true
		} else if char == '\\' {
			escaped = true
		}

		value = append(value, char)
		tok.consume()

		if end {
			return &Token{
				Type: STRING,
				Location: Location{
					Filename: tok.filename,
					Start:    start,
					End:      tok.Position,
				},
				Value: string(value),
			}, nil
		}
	}
}

// return (nil, nil) if it doesn't match.
func maybeParseLineComment(tok *rawTokenizer) ([]byte, error) {
	value, err := tok.peekN(2)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}

		return nil, err
	}

	if value[0] != '/' || value[1] != '/' {
		return nil, nil
	}

	tok.consumeN(2)

	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if char == '\n' {
			break
		}

		value = append(value, char)
		tok.consume()
	}

	return value, nil
}

// return (nil, nil) if it doesn't match.
func maybeParseCommentBlock(tok *rawTokenizer) ([]byte, error) {
	value, err := tok.peekN(2)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}

		return nil, err
	}

	if value[0] != '/' || value[1] != '*' {
		return nil, nil
	}

	tok.consumeN(2)

	scope := 1
	for {
		char, err := tok.peek()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf(
					"%s:%v: invalid block comment",
					filepath.Base(tok.filename),
					tok.Position)
			}
		}

		if char == '/' {
			bytes, err := tok.peekN(2)
			if err != nil {
				if err == io.EOF {
					return nil, fmt.Errorf(
						"%s:%v: invalid block comment",
						filepath.Base(tok.filename),
						tok.Position)
				}

				return nil, err
			}

			if bytes[1] == '*' {
				value = append(value, '/', '*')
				tok.consumeN(2)
				scope += 1
				continue
			}
		} else if char == '*' {
			bytes, err := tok.peekN(2)
			if err != nil {
				if err == io.EOF {
					return nil, fmt.Errorf(
						"%s:%v: invalid block comment",
						filepath.Base(tok.filename),
						tok.Position)
				}

				return nil, err
			}

			if bytes[1] == '/' {
				value = append(value, '*', '/')
				tok.consumeN(2)
				scope -= 1

				if scope == 0 {
					break
				}

				continue
			}
		}

		value = append(value, char)
		tok.consume()
	}

	return value, nil
}

func parseSlash(tok *rawTokenizer) (*Token, error) {
	start := tok.Position

	bytes, err := maybeParseLineComment(tok)
	if err != nil {
		return nil, err
	}

	if len(bytes) > 0 {
		return &Token{
			Type: COMMENT,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: string(bytes),
		}, nil
	}

	bytes, err = maybeParseCommentBlock(tok)
	if err != nil {
		return nil, err
	}

	if len(bytes) > 0 {
		return &Token{
			Type: COMMENT,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: string(bytes),
		}, nil
	}

	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf(
				"%s:%v: invalid symbol",
				filepath.Base(tok.filename),
				tok.Position)
		}

		return nil, err
	}

	if char != '/' {
		return nil, fmt.Errorf(
			"%s:%v: invalid symbol",
			filepath.Base(tok.filename),
			tok.Position)
	}

	tok.consume()

	return &Token{
		Type: DIV,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: "/",
	}, nil
}

func parseSymbol(symbolTokens map[string]int) parseTokenFunc {
	return func(tok *rawTokenizer) (*Token, error) {
		matchedSymbol := ""
		matchedTokenType := 0
		for symbol, tokenType := range symbolTokens {
			if len(matchedSymbol) >= len(symbol) {
				continue
			}

			bytes, err := tok.peekN(len(symbol))
			if err != nil {
				if err != io.EOF {
					return nil, err
				}

				continue
			}

			if string(bytes) == symbol {
				matchedSymbol = symbol
				matchedTokenType = tokenType
			}
		}

		if matchedSymbol == "" {
			return nil, fmt.Errorf(
				"%s:%v: invalid symbol",
				filepath.Base(tok.filename),
				tok.Position)
		}

		start := tok.Position
		tok.consumeN(len(matchedSymbol))

		return &Token{
			Type: matchedTokenType,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: matchedSymbol,
		}, nil
	}
}

type rawTokenizer struct {
	filename string
	reader   io.Reader

	Position

	buffered []byte
}

func NewRawTokenizer(filename string, reader io.Reader) (Tokenizer, error) {
	tok := &rawTokenizer{
		filename: filename,
		reader:   reader,
		Position: Position{
			Column: 1,
			Line:   1,
		},
	}

	err := tok.stripMeaninglessWhitespaces()
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func (tok *rawTokenizer) Filename() string {
	return tok.filename
}

func (tok *rawTokenizer) Pos() Position {
	return tok.Position
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

func (tok *rawTokenizer) Next() (*Token, error) {
	token, err := tok.parseNextToken()
	if err != nil {
		return nil, err
	}

	err = tok.stripMeaninglessWhitespaces()
	if err != nil {
		return nil, err
	}

	return token, nil
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

func (tok *rawTokenizer) parseNextToken() (*Token, error) {
	char, err := tok.peek()
	if err != nil {
		return nil, err
	}

	parseToken, ok := leadingCharParseTokenFunc[char]
	if !ok {
		return nil, fmt.Errorf(
			"%s:%v: unknown character: %c",
			filepath.Base(tok.filename),
			tok.Position,
			char)
	}

	return parseToken(tok)
}
