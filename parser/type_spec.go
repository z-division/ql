package parser

type TypeSpec interface {
	Node
	isTypeSpec()
}

type typeSpec struct{}

func (typeSpec) isTypeSpec() {}

type ScalarType struct {
	Location
	typeSpec

	Type *Token
}

type NamedType struct {
	Location
	typeSpec

	Type *Token
}
