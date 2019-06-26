package parser

type Declaration interface {
	Node
	isDeclaration()
}

type declaration struct{}

func (declaration) isDeclaration() {}

// TODO(patrick): Remove once we have defined other real declarations
type FakeDeclaration struct {
	Location
	declaration

	Expression Expr
}

func (fake *FakeDeclaration) String() string {
	return prettyFormatNode("", fake, 0)
}

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
