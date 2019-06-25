// NOTE(patrick): the concrete syntax parse tree intentionally preserves all
// parsed tokens, which may be use to reconstruct / format the file.

package parser

type ControlFlowExpr interface {
	Expr
	isControlFlowExpr()
}

type controlFlowExpr struct {
	expr
}

func (controlFlowExpr) isControlFlowExpr() {}

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
	TrueClause  *ExprBlock
	Else        *Token          // Optional
	FalseClause ControlFlowExpr // Optional
}

func (cond *ConditionalExpr) String() string {
	return prettyFormatNode("", cond, 0)
}

type ForExpr struct {
	Location
	controlFlowExpr

	*ScopeDef // Optional
	For       *Token
	Predicate Expr
	Body      *ExprBlock
}

func (fe *ForExpr) String() string {
	return prettyFormatNode("", fe, 0)
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

	Let        *Token // Optional
	Name       *Token
	TypeSpec   TypeSpec // Optional (disallow implicit embedded field access)
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
