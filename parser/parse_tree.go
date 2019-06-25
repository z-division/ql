// NOTE(patrick): the concrete syntax parse tree intentionally preserves all
// parsed tokens, which may be use to reconstruct / format the file.

package parser

import (
	"fmt"
	"path/filepath"
	"reflect"
)

const (
	indentLevel = "  "
)

type customFormatter interface {
	prettyFormat(prefix string, indent int) string
}

func formatIdent(indent int) string {
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += indentLevel
	}

	return indentStr
}

func prettyFormatNode(prefix string, node Node, indent int) string {
	formatter, ok := node.(customFormatter)
	if ok {
		return formatter.prettyFormat(prefix, indent)
	}

	indentStr := formatIdent(indent)

	nodeStruct := reflect.ValueOf(node).Elem()
	structType := nodeStruct.Type()

	result := fmt.Sprintf(
		"%s%s[%s (%v)",
		indentStr,
		prefix,
		structType.Name(),
		node.Loc())

	addNewLine := false
	for i := 0; i < nodeStruct.NumField(); i++ {
		field := structType.Field(i)
		if field.PkgPath != "" { // private field
			continue
		}

		if field.Name == "Location" {
			continue
		}

		addNewLine = true

		fieldValue := nodeStruct.Field(i)
		switch value := fieldValue.Interface().(type) {
		case Node:
			if fieldValue.IsNil() {
				result += fmt.Sprintf(
					"\n%s%s%s=<nil>",
					indentLevel,
					indentStr,
					field.Name)
			} else {
				result += "\n" + prettyFormatNode(
					field.Name+" = ",
					value,
					indent+1)
			}
		case []ControlFlowExpr:
			result += fmt.Sprintf(
				"\n%s%s%s = [",
				indentStr,
				indentLevel,
				field.Name)
			if len(value) == 0 {
				result += "]"
			} else {
				for _, argument := range value {
					result += "\n" + prettyFormatNode(
						"",
						argument,
						indent+2)
				}
				result += "\n" + indentStr + indentLevel + "]"
			}
		case []*Argument:
			result += fmt.Sprintf(
				"\n%s%s%s = [",
				indentStr,
				indentLevel,
				field.Name)
			if len(value) == 0 {
				result += "]"
			} else {
				for _, arg := range value {
					result += "\n" + prettyFormatNode(
						"",
						arg,
						indent+2)
				}
				result += "\n" + indentStr + indentLevel + "]"
			}
		default:
			result += fmt.Sprintf(
				"\n%s%s%s = %v",
				indentStr,
				indentLevel,
				field.Name,
				value)
		}
	}

	if addNewLine {
		result += "\n" + indentStr + "]"
	} else {
		result += "]"
	}

	return result
}

type Position struct {
	Line   int
	Column int
}

func (pos Position) String() string {
	return fmt.Sprintf("%d:%d", pos.Line, pos.Column)
}

type Location struct {
	Filename string
	Start    Position // inclusive
	End      Position // exclusive
}

func (loc Location) Loc() Location {
	return loc
}

func (Location) isNode() {
}

func (loc1 Location) Merge(loc2 Location) Location {
	return Location{
		Filename: loc1.Filename,
		Start:    loc1.Start,
		End:      loc2.End,
	}
}

func (loc Location) String() string {
	return fmt.Sprintf(
		"%s: %v-%v",
		filepath.Base(loc.Filename),
		loc.Start,
		loc.End)
}

type Node interface {
	Loc() Location
	String() string
	isNode()
}

// A list of comment token separated only by at most a single newline
type CommentGroup []string

// Similar to golang, a comment group g is associated with a non-comment token
// n if:
// - g starts on the same line as n ends
// - g starts on the line immediately following n, and there is at least one
//   empty line after g and before the next node
// - g starts before n and is not associated to the node before n via
//   the previous rules.
//
// NOTE: comments are never associated with NEWLINE, nor SEMICOLON
type Comments struct {
	Leading  []CommentGroup
	Trailing CommentGroup
}

func (comments Comments) prettyFormat(indent int) string {
	if len(comments.Leading) == 0 && len(comments.Trailing) == 0 {
		return ""
	}

	indentStr := formatIdent(indent)

	lines := "\n"

	if len(comments.Leading) > 0 {
		lines += indentStr + indentLevel + "Leading Comments:\n"
		for i, group := range comments.Leading {
			if i > 0 {
				lines += indentStr + indentLevel + "===\n"
			}

			for _, comment := range group {
				lines += indentStr + indentLevel + indentLevel + comment + "\n"
			}
		}
	}

	if len(comments.Trailing) > 0 {
		lines += indentStr + indentLevel + "Trailing Comment:\n"
		for _, comment := range comments.Trailing {
			lines += indentStr + indentLevel + indentLevel + comment + "\n"
		}
	}

	lines += indentStr
	return lines
}

type Token struct {
	Type int
	Location
	Value string

	Comments
}

func (token *Token) Name() string {
	diff := token.Type - LEX_ERROR + 3
	if 0 <= diff && diff < len(qlToknames) {
		return qlToknames[diff]
	}

	return fmt.Sprintf("UNKNOWN(%d)", token.Type)
}

func (token *Token) prettyFormat(prefix string, indent int) string {
	if token.Type == NEWLINE {
		return fmt.Sprintf(
			"%s%s[%s (%v)%v]",
			formatIdent(indent),
			prefix,
			token.Name(),
			token.Location,
			token.Comments.prettyFormat(indent))
	}

	return fmt.Sprintf(
		"%s%s[%s %s (%v)%v]",
		formatIdent(indent),
		prefix,
		token.Name(),
		token.Value,
		token.Location,
		token.Comments.prettyFormat(indent))
}

func (op *Token) String() string {
	return prettyFormatNode("", op, 0)
}

type Expr interface {
	Node
	isExpr()
}

type expr struct{}

func (expr) isExpr() {}

type ControlFlowExpr interface {
	Expr
	isControlFlowExpr()
}

type controlFlowExpr struct {
	expr
}

func (controlFlowExpr) isControlFlowExpr() {}

// <value>
type Identifier struct {
	Location
	expr

	Value *Token
}

func (id *Identifier) String() string {
	return prettyFormatNode("", id, 0)
}

// <value>
type Literal struct {
	Location
	expr

	Value *Token
}

func (literal *Literal) String() string {
	return prettyFormatNode("", literal, 0)
}

// <op> <expression>
type UnaryExpr struct {
	Location
	expr

	Op         *Token
	Expression Expr
}

func (unary *UnaryExpr) String() string {
	return prettyFormatNode("", unary, 0)
}

// <left> <op> <right>
type BinaryExpr struct {
	Location
	expr

	Left  Expr
	Op    *Token
	Right Expr
}

func (binary *BinaryExpr) String() string {
	return prettyFormatNode("", binary, 0)
}

// <primary expr> . <name>
type Accessor struct {
	Location
	expr

	PrimaryExpr Expr
	Dot         *Token
	Name        *Token
}

func (accessor *Accessor) String() string {
	return prettyFormatNode("", accessor, 0)
}

type Argument struct {
	Location
	Expression Expr
	Comma      *Token
}

type Invocation struct {
	Location
	expr

	Expression Expr
	LParen     *Token
	Arguments  []*Argument
	RParen     *Token
}

func (invoke *Invocation) String() string {
	return prettyFormatNode("", invoke, 0)
}

// Placeholder for comments at the end of file
type Noop struct {
	Location
	controlFlowExpr

	Value *Token
}

func (noop *Noop) String() string {
	return prettyFormatNode("", noop, 0)
}

type ScopeDef struct {
	Location

	Name *Token
	At   *Token
}

func (sd *ScopeDef) String() string {
	return prettyFormatNode("", sd, 0)
}

type ExprBlock struct {
	Location
	controlFlowExpr

	*ScopeDef  // Optional
	LBrace     *Token
	Statements []ControlFlowExpr
	RBrace     *Token
}

func (block *ExprBlock) String() string {
	return prettyFormatNode("", block, 0)
}

type ConditionalExpr struct {
	Location
	controlFlowExpr

	*ScopeDef   // Optional
	If          *Token
	Predicate   Expr
	TrueClause  ControlFlowExpr
	Else        *Token
	FalseClause ControlFlowExpr
}

func (cond *ConditionalExpr) String() string {
	return prettyFormatNode("", cond, 0)
}

// <expr>
type EvalExpr struct {
	Location
	controlFlowExpr

	Expression Expr
}

func (eval *EvalExpr) String() string {
	return prettyFormatNode("", eval, 0)
}

// let <name> = <expr>
type AssignExpr struct {
	Location
	controlFlowExpr

	Let        *Token
	Name       *Token
	Assign     *Token
	Expression Expr
}

func (assign *AssignExpr) String() string {
	return prettyFormatNode("", assign, 0)
}

type ReturnExpr struct {
	Location
	controlFlowExpr

	Return     *Token
	At         *Token // Optional
	Label      *Token // Optional
	Expression Expr
}

func (ret *ReturnExpr) String() string {
	return prettyFormatNode("", ret, 0)
}
