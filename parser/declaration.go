package parser

type Declaration interface {
	Node
	isDeclaration()
}

type declaration struct{}

func (declaration) isDeclaration() {}

type TypeDef struct {
	Location
	declaration

	Type     *Token
	Name     *Token
	TypeSpec TypeSpec
}

func (td *TypeDef) String() string {
	return prettyFormatNode("", td, 0)
}

type Parameter struct {
	Location
	Name     *Token
	TypeSpec TypeSpec
	Comma    *Token // Optional
}

type FuncDef struct {
	Location
	declaration

	Func       *Token
	Name       *Token
	LParen     *Token
	Parameters []*Parameter
	RParen     *Token
	ReturnType TypeSpec // Optional
	Body       Expr
}

func (fd *FuncDef) String() string {
	return prettyFormatNode("", fd, 0)
}
