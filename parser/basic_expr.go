package parser

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

type Argument struct {
	Location
	Expression Expr
	Comma      *Token
}

type TypeConversion struct {
	Location
	expr

	Type       *Token
	LParen     *Token
	Expression Expr
	RParen     *Token
}

func (tc *TypeConversion) String() string {
	return prettyFormatNode("", tc, 0)
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
