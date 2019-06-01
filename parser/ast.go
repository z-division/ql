package parser

import (
	"fmt"
	"reflect"
)

const (
	indentLevel = "  "
)

func formatIdent(indent int) string {
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += indentLevel
	}

	return indentStr
}

func prettyFormatNode(prefix string, node Node, indent int) string {
	indentStr := formatIdent(indent)

	nodeStruct := reflect.ValueOf(node).Elem()
	structType := nodeStruct.Type()

	result := fmt.Sprintf(
		"%s%s[%s (%v)",
		indentStr,
		prefix,
		structType.Name(),
		node.Position())

	for i := 0; i < nodeStruct.NumField(); i++ {
		field := structType.Field(i)
		if field.PkgPath != "" { // private field
			continue
		}

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

	result += "\n" + indentStr + "]"

	return result
}

type Location struct {
	Filename    string
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
}

func (loc Location) String() string {
	return fmt.Sprintf("%d:%d", loc.StartLine, loc.StartColumn)
}

type Node interface {
	Position() Location
	String() string
	isNode()
}

type node struct {
	Location
}

func (n node) Position() Location {
	return n.Location
}

func (node) isNode() {
}

type Expr interface {
	Node
	isExpr()
}

type expr struct{}

func (expr) isExpr() {}

type Identifier struct {
	node
	expr

	Value string
}

func (id *Identifier) String() string {
	return prettyFormatNode("", id, 0)
}

type AssignExpr struct {
	node
	expr

	Name       string
	Expression Expr
}

func (assign *AssignExpr) String() string {
	return prettyFormatNode("", assign, 0)
}
