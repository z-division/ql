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

type IterType struct {
	Location
	typeSpec

	Iter *Token

	// The following are optional if the element type is not binded yet.
	LBracket    *Token
	ElementType TypeSpec
	RBracket    *Token
}

type RecordType struct {
	Location
	typeSpec

	Iter *Token

	// The following are optional if the field names and types are not
	// binded yet.
	Lt *Token
	Fields []*Parameter
	Gt *Token
}
