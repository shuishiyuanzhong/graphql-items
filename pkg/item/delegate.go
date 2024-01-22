package item

import "github.com/graphql-go/graphql"

type Delegate interface {
	Name() string
	Type() FieldType
	// BuildItem 生成GraphQL Object
	BuildItem(graphql.Fields) *graphql.List
	Resolve() graphql.FieldResolveFn
	BuildField2() []*Field

	IsList() bool
}
