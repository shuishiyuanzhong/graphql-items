package item

import "github.com/graphql-go/graphql"

type Delegate interface {
	Name() string
	BuildItem() *graphql.Object
	Resolve() graphql.FieldResolveFn
}

type DefaultDelegate struct{}

func (d *DefaultDelegate) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return Product{ID: "1", Name: "Example Product", Price: 99.99}, nil
	}
}

// DTO to GraphQL Object

type ItemDemo struct {
	*DefaultDelegate
}

func (d *ItemDemo) Name() string {
	return "item"
}

// 构造当前的Item的GraphQL Object
// but 这个对象自己Resolve还没有生成出来
func (d *ItemDemo) BuildItem() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   d.Name(),
			Fields: d.BuildField(),
		},
	)
}

// 声明当前Item所拥有的字段，以及各自字段的的Resolve
func (d *ItemDemo) BuildField() graphql.Fields {
	fields := make(graphql.Fields)
	fields["id"] = &graphql.Field{
		Type: graphql.String,
	}
	fields["name"] = &graphql.Field{
		Type: graphql.String,
	}
	fields["price"] = &graphql.Field{
		Type: graphql.Float,
	}
	fields["test"] = &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "test", nil
		},
	}
	return fields
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
