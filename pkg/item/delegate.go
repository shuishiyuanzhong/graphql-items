package item

import "github.com/graphql-go/graphql"

type Delegate interface {
	Name() string
	// BuildItem 生成GraphQL Object
	BuildItem() *graphql.List
	Resolve() graphql.FieldResolveFn
}
