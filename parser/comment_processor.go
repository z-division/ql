package parser

import (
	"fmt"
	"io"
	"strings"
)

type tokenOrCommentGroup struct {
	*Token
	CommentGroup
}

func (toc tokenOrCommentGroup) String() string {
	if toc.Token == nil {
		return fmt.Sprintf("%#v", strings.Join(toc.CommentGroup, "\n"))
	} else {
		return toc.Token.Name()
	}
}

type commentProcessor struct {
	Tokenizer

	buffered []*Token
}

func NewCommentProcessor(base Tokenizer) (Tokenizer, error) {
	proc := &commentProcessor{
		Tokenizer: base,
	}

	err := proc.maybeFill()
	if err != nil {
		return nil, err
	}

	return proc, nil
}

func (proc *commentProcessor) isTerminal(tok *Token) bool {
	return tok.Type == NEWLINE || tok.Type == SEMICOLON
}

func (proc *commentProcessor) mergeComments(
	lookAhead []*Token) []tokenOrCommentGroup {

	result := []tokenOrCommentGroup{}
	idx := 0
	for idx < len(lookAhead) {
		token := lookAhead[idx]

		if len(result) == 0 {
			if token.Type == COMMENT {
				result = append(
					result,
					tokenOrCommentGroup{
						CommentGroup: CommentGroup{token.Value},
					})
			} else {
				result = append(
					result,
					tokenOrCommentGroup{
						Token: token,
					})
			}

			idx++
			continue
		}

		if token.Type == COMMENT {
			if result[len(result)-1].Token == nil {
				// comment group continuation
				result[len(result)-1].CommentGroup =
					append(result[len(result)-1].CommentGroup, token.Value)
			} else { // new comment group
				result = append(
					result,
					tokenOrCommentGroup{
						CommentGroup: CommentGroup{token.Value},
					})
			}

			idx++
			continue
		}

		if token.Type == NEWLINE &&
			result[len(result)-1].Token == nil &&
			idx+1 < len(lookAhead) &&
			lookAhead[idx+1].Type == COMMENT {
			// comment group continuation

			result[len(result)-1].CommentGroup =
				append(
					result[len(result)-1].CommentGroup,
					lookAhead[idx+1].Value)

			idx += 2
			continue
		}

		// regular token
		result = append(
			result,
			tokenOrCommentGroup{
				Token: token,
			})

		idx++
	}

	return result
}

func (proc *commentProcessor) maybeFill() error {
	if len(proc.buffered) > 1 {
		return nil
	}

	var lookAhead []*Token
	for {
		token, err := proc.Tokenizer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		lookAhead = append(lookAhead, token)

		if token.Type != COMMENT && !proc.isTerminal(token) {
			break
		}
	}

	if len(lookAhead) == 0 {
		return nil
	}

	tokOrCom := proc.mergeComments(lookAhead)

	if len(proc.buffered) > 0 {
		// The current token may have trailing comments

		if len(proc.buffered) != 1 ||
			proc.isTerminal(proc.buffered[0]) {

			panic("Programming error")
		}

		curr := proc.buffered[0]

		if tokOrCom[0].Token == nil {
			// Comment starts on same line as curr.
			// This is always a trailing comment

			curr.Trailing = tokOrCom[0].CommentGroup
			tokOrCom = tokOrCom[1:]
		} else if tokOrCom[0].Token.Type == NEWLINE &&
			len(tokOrCom) > 1 &&
			tokOrCom[1].Token == nil {

			// potential trailing comment

			if len(tokOrCom) == 2 {
				curr.Trailing = tokOrCom[1].CommentGroup
				tokOrCom = tokOrCom[2:]
			} else if len(tokOrCom) == 3 &&
				proc.isTerminal(tokOrCom[2].Token) {

				curr.Trailing = tokOrCom[1].CommentGroup
				tokOrCom = tokOrCom[2:]
			} else if proc.isTerminal(tokOrCom[2].Token) &&
				proc.isTerminal(tokOrCom[3].Token) {

				curr.Trailing = tokOrCom[1].CommentGroup
				tokOrCom = tokOrCom[2:]
			}
		}
	}

	var comments []CommentGroup
	for _, toc := range tokOrCom {
		if toc.Token == nil {
			comments = append(comments, toc.CommentGroup)
		} else {
			proc.buffered = append(proc.buffered, toc.Token)
		}
	}

	if len(comments) > 0 {
		if len(proc.buffered) == 0 ||
			proc.isTerminal(proc.buffered[len(proc.buffered)-1]) {

			// comments at end of file.  Associate comment to a dummy token
			proc.buffered = append(
				proc.buffered,
				&Token{
					Type: NOOP,
					Location: Location{
						Filename: proc.Filename(),
						Start:    proc.Pos(),
						End:      proc.Pos(),
					},
				})
		}

		proc.buffered[len(proc.buffered)-1].Leading = comments
	}

	return nil
}

func (proc *commentProcessor) Next() (*Token, error) {
	err := proc.maybeFill()
	if err != nil {
		return nil, err
	}

	if len(proc.buffered) == 0 {
		return nil, io.EOF
	}

	tok := proc.buffered[0]
	proc.buffered = proc.buffered[1:]
	return tok, nil
}
