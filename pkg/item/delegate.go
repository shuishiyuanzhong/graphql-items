package item

import "github.com/graphql-go/graphql"

type Delegate interface {
	// Name graphQL查询的字段名称
	Name() string
	// Type item对象类型
	Type() FieldType

	Resolve() graphql.FieldResolveFn
	BuildField() []*Field

	IsList() bool
}
