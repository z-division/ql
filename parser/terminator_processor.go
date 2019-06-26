package parser

import (
	"io"
)

var (
	// Tokens that are sensitive to terminators on either side of the token
	terminatorSensitive = map[int]struct{}{
		IDENT:          struct{}{},
		BYTE_LITERAL:   struct{}{},
		STRING_LITERAL: struct{}{},
		INT_LITERAL:    struct{}{},
		FLOAT_LITERAL:  struct{}{},
		BOOL_LITERAL:   struct{}{},
		NOOP:           struct{}{},
		R_BRACE:        struct{}{},
	}

	// Tokens that are sensitive to terminators on the leading side of the
	// token
	leadingTerminatorSensitive = map[int]struct{}{
		LET:    struct{}{},
		IF:     struct{}{},
		RETURN: struct{}{},
		TYPE:   struct{}{},
		FUNC:   struct{}{},

		// expression block needs special parser handling
		L_BRACE: struct{}{},

		// type specs are leading newline sensitive depending on context
		BOOL_TYPE:   struct{}{},
		INT_TYPE:    struct{}{},
		UINT_TYPE:   struct{}{},
		FLOAT_TYPE:  struct{}{},
		BYTE_TYPE:   struct{}{},
		STRING_TYPE: struct{}{},
		ITER_TYPE:   struct{}{},
		RECORD_TYPE: struct{}{},
	}

	// Tokens that are sensitive to terminators on the trailing side of the
	// token
	trailingTerminatorSensitive = map[int]struct{}{
		R_PAREN: struct{}{},
	}

	// Tokens that are implicitly lead by a terminator
	implicitLeadingTerminator = map[int]struct{}{
		// This ensures the last statement in the block terminates
		R_BRACE: struct{}{},
	}

	// Tokens that are implicitly trail by a terminator
	implicitTrailingTerminator = map[int]struct{}{}
)

// This tokenizer drops context insensitive newline tokens
type terminatorProcessor struct {
	Tokenizer

	prevToken *Token
	buffered  []*Token
}

func NewTerminatorProcessor(base Tokenizer) (Tokenizer, error) {
	return &terminatorProcessor{
		Tokenizer: base,
	}, nil
}

func (proc *terminatorProcessor) maybeFill() error {
	if len(proc.buffered) > 0 {
		return nil
	}

	var terminator *Token
	var nextNonTerminator *Token
	for {
		token, err := proc.Tokenizer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type == NEWLINE {
			if terminator == nil {
				terminator = token
			}

			continue
		}

		if token.Type == SEMICOLON {
			terminator = token
			continue
		}

		nextNonTerminator = token
		break
	}

	if proc.prevToken == nil {
		if nextNonTerminator != nil {
			proc.buffered = []*Token{nextNonTerminator}
			return nil
		}

		// "empty" file

		if terminator == nil {
			terminator = &Token{
				Type: NEWLINE,
				Location: Location{
					Filename: proc.Filename(),
					Start:    proc.Pos(),
					End:      proc.Pos(),
				},
			}
		}

		proc.buffered = []*Token{terminator}
		return nil
	}

	if nextNonTerminator == nil { // EOF
		if proc.prevToken.Type == NEWLINE ||
			proc.prevToken.Type == SEMICOLON {
			return io.EOF
		}

		if terminator == nil {
			terminator = &Token{
				Type:     NEWLINE,
				Location: proc.prevToken.Location,
			}
			terminator.Start = terminator.End
		}

		proc.buffered = []*Token{terminator}
		return nil
	}

	// Always preserve explicit terminator
	if terminator != nil && terminator.Type == SEMICOLON {
		proc.buffered = []*Token{terminator, nextNonTerminator}
		return nil
	}

	// Insert mandatory implicit terminator
	implicitTerminator := false
	_, ok := implicitTrailingTerminator[proc.prevToken.Type]
	if ok {
		implicitTerminator = true
	}

	_, ok = implicitLeadingTerminator[nextNonTerminator.Type]
	if ok {
		implicitTerminator = true
	}

	if implicitTerminator && terminator == nil {
		terminator = &Token{
			Type:     NEWLINE,
			Location: nextNonTerminator.Location,
		}
		terminator.End = terminator.Start
	}

	// No newline in between prev and next tokens.  Don't insert terminator.
	if terminator == nil {
		proc.buffered = []*Token{nextNonTerminator}
		return nil
	}

	// Preserve terminator if both prev and next token are terminator sensitive
	_, leading := leadingTerminatorSensitive[nextNonTerminator.Type]
	_, ok = terminatorSensitive[nextNonTerminator.Type]
	if ok {
		leading = true
	}

	_, trailing := trailingTerminatorSensitive[proc.prevToken.Type]
	_, ok = terminatorSensitive[proc.prevToken.Type]
	if ok {
		trailing = true
	}

	if leading && trailing {
		proc.buffered = []*Token{terminator, nextNonTerminator}
		return nil
	}

	proc.buffered = []*Token{nextNonTerminator}
	return nil
}

func (proc *terminatorProcessor) Next() (*Token, error) {
	err := proc.maybeFill()
	if err != nil {
		return nil, err
	}

	proc.prevToken = proc.buffered[0]
	proc.buffered = proc.buffered[1:]
	return proc.prevToken, nil
}
