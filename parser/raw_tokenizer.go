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
		Type: IDENT,
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

	return IDENT, nil
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

    tok.consumeN(len(bytes)-1)

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

func parseHex(tok *rawTokenizer, lval *qlSymType) (int, error) {
    start := tok.Position

    value, err := tok.peekN(2)
    if err != nil {
        if err != io.EOF {
            return LEX_ERROR, fmt.Errorf(
                "%s:%v: invaild hexadecimal",
                filepath.Base(tok.filename),
                tok.Position)
        }

        return LEX_ERROR, err
    }

    if !(value[0] == '0' &&
        (value[1] == 'x' || value[1] == 'X')) {

        return LEX_ERROR, fmt.Errorf(
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
            return LEX_ERROR, err
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
        return LEX_ERROR, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
    }

    lval.Token = &Token{
        Type: INT,
        Location: Location{
            Filename: tok.filename,
            Start:    start,
            End:      tok.Position,
        },
        Value: string(value),
    }

    return INT, nil
}

func parseOct(tok *rawTokenizer, lval *qlSymType) (int, error) {
    start := tok.Position

    var value []byte
    for {
        char, err := tok.peek()
        if err != nil {
            if err == io.EOF {
                break
            }
            return LEX_ERROR, err
        }

        if !('0' <= char && char <= '7') {
            break
        }

        value = append(value, char)
        tok.consume()
    }

    if len(value) == 0 {
        return LEX_ERROR, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
    }

    lval.Token = &Token{
        Type: INT,
        Location: Location{
            Filename: tok.filename,
            Start:    start,
            End:      tok.Position,
        },
        Value: string(value),
    }

    return INT, nil
}

func parseNumber(tok *rawTokenizer, lval *qlSymType) (int, error) {
    bytes, err := tok.peekN(2)
    if err != nil && err != io.EOF {
        return LEX_ERROR, err
    }

    if err != io.EOF && bytes[0] == '0' {
        if bytes[1] == 'x' || bytes[1] == 'X' {  // hex
            return parseHex(tok, lval)
        } else { // oct
            return parseOct(tok, lval)
        }
    }

    start := tok.Position

    value, err := maybeParseDecimal(tok)
    if err != nil {
        return LEX_ERROR, err
    }

    if len(value) == 0 {
        return LEX_ERROR, fmt.Errorf(
			"%s:%v: invaild number",
			filepath.Base(tok.filename),
			tok.Position)
    }

	fractional, err := maybeParseFloatFractional(tok)
	if err != nil {
		return LEX_ERROR, err
	}

    exponent, err := maybeParseFloatExponent(tok)
    if err != nil {
        return LEX_ERROR, err
    }

    numType := INT
    if len(fractional) > 0 || len(exponent) > 0 {
        numType = FLOAT
        value = append(value, fractional...)
        value = append(value, exponent...)
    }

    lval.Token = &Token{
        Type: numType,
        Location: Location{
            Filename: tok.filename,
            Start:    start,
            End:      tok.Position,
        },
        Value: string(value),
    }

    return numType, nil
}

func parseDot(tok *rawTokenizer, lval *qlSymType) (int, error) {
	start := tok.Position

	fractional, err := maybeParseFloatFractional(tok)
	if err != nil {
		return LEX_ERROR, err
	}

    if len(fractional) == 0 {
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: invaild symbol",
			filepath.Base(tok.filename),
			tok.Position)
    } else if len(fractional) == 1 {  // single dot
        lval.Token = &Token{
            Type: DOT,
            Location: Location{
                Filename: tok.filename,
                Start:    start,
                End:      tok.Position,
            },
            Value: ".",
        }

        return DOT, nil
    } else { // dot follow by digit is a float
		exponent, err := maybeParseFloatExponent(tok)
		if err != nil {
			return LEX_ERROR, err
		}

		lval.Token = &Token{
			Type: FLOAT,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: string(append(fractional, exponent...)),
		}

		return FLOAT, nil
	}
}

func parseChar(tok *rawTokenizer, lval *qlSymType) (int, error) {
	value, err := tok.peekN(3)
	if err != nil || value[0] != '\'' {
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	// '' is invalid
	if value[1] == '\'' {
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	start := tok.Position

	if value[1] == '\\' {
		_, ok := validEscapeChars[value[2]]
		if !ok {
			return LEX_ERROR, fmt.Errorf(
				"%s:%v: invalid escaped character",
				filepath.Base(tok.filename),
				tok.Position)
		}

		value, err = tok.peekN(4)
		if err != nil || value[3] != '\'' {
			return LEX_ERROR, fmt.Errorf(
				"%s:%v: invalid char literal",
				filepath.Base(tok.filename),
				tok.Position)
		}

		tok.consumeN(4)
		lval.Token = &Token{
			Type: CHAR,
			Location: Location{
				Filename: tok.filename,
				Start:    start,
				End:      tok.Position,
			},
			Value: string(value),
		}

		return CHAR, nil
	}

	_, ok := invalidLiteralChars[value[1]]
	if ok || value[2] != '\'' {
		return LEX_ERROR, fmt.Errorf(
			"%s:%v: invalid char literal",
			filepath.Base(tok.filename),
			tok.Position)
	}

	tok.consumeN(3)
	lval.Token = &Token{
		Type: CHAR,
		Location: Location{
			Filename: tok.filename,
			Start:    start,
			End:      tok.Position,
		},
		Value: string(value),
	}

	return CHAR, nil
}

func parseString(tok *rawTokenizer, lval *qlSymType) (int, error) {
	start := tok.Position

	char, err := tok.peek()
	if err != nil {
		if err == io.EOF {
			return LEX_ERROR, fmt.Errorf(
				"%s:%v: invalid string literal",
				filepath.Base(tok.filename),
				tok.Position)
		}
		return LEX_ERROR, err
	}

	if char != '"' {
		return LEX_ERROR, fmt.Errorf(
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
				return LEX_ERROR, fmt.Errorf(
					"%s:%v: invalid string literal",
					filepath.Base(tok.filename),
					tok.Position)
			}

			return LEX_ERROR, err
		}

		_, ok := invalidLiteralChars[char]
		if ok {
			return LEX_ERROR, fmt.Errorf(
				"%s:%v: invalid string literal",
				filepath.Base(tok.filename),
				tok.Position)
		}

		end := false
		if escaped {
			_, ok := validEscapeChars[char]
			if !ok {
				return LEX_ERROR, fmt.Errorf(
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
			lval.Token = &Token{
				Type: STRING,
				Location: Location{
					Filename: tok.filename,
					Start:    start,
					End:      tok.Position,
				},
				Value: string(value),
			}

			return STRING, nil
		}
	}
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
