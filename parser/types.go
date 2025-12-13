package parser

type Type interface {
	TypeKind() string
}

type PrimitiveType struct {
	Name string
}

func (p PrimitiveType) TypeKind() string { return "primitive" }

type Field struct {
	Name string
	ID   uint16
	Type Type
}

type StructType struct {
	ID     uint16
	Name   string
	Fields []Field
}

func (s StructType) TypeKind() string { return "struct" }

type ListType struct {
	ElementType Type
}

func (l ListType) TypeKind() string { return "list" }

type Schema struct {
	Name    string
	Version int
	Structs []StructType
}

type MapType struct {
	KeyType   Type
	ValueType Type
}

func (m MapType) TypeKind() string { return "map" }
