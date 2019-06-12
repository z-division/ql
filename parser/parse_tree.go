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

func getTokenName(token int) string {
	diff := token - LEX_ERROR + 3
	if 0 <= diff && diff < len(qlToknames) {
		return qlToknames[diff]
	}

	return fmt.Sprintf("UNKNOWN(%d)", token)
}

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

		fieldValue := nodeStruct.Field(i).Interface()
		switch value := fieldValue.(type) {
		case Node:
			result += "\n" + prettyFormatNode(
				field.Name+" = ",
				value,
				indent+1)
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

type Token struct {
	Type int
	Location
	Value string
}

func (token *Token) prettyFormat(prefix string, indent int) string {
	return fmt.Sprintf(
		"%s%s[%s %s (%v)]",
		formatIdent(indent),
		prefix,
		getTokenName(token.Type),
		token.Value,
		token.Location)
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

// ( <expression> )
type EvalOrderExpr struct {
	Location
	expr

	LParen     *Token
	Expression Expr
	RParen     *Token
}

func (evalOrder *EvalOrderExpr) String() string {
	return prettyFormatNode("", evalOrder, 0)
}

// let <name> = <expr>
type AssignExpr struct {
	Location
	expr

	Let        *Token
	Name       *Token
	Assign     *Token
	Expression Expr
}

func (assign *AssignExpr) String() string {
	return prettyFormatNode("", assign, 0)
}
