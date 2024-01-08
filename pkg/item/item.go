package item

import "github.com/graphql-go/graphql"

type Delegate interface {
	Name() string
	// BuildItem 生成GraphQL Object
	BuildItem() *graphql.List
	Resolve() graphql.FieldResolveFn
}

type DefaultDelegate struct{}

func (d *DefaultDelegate) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return []map[string]interface{}{{"id": "1", "name": "Example Product", "price": 99.99}}, nil
	}
}

// DTO to GraphQL Object
