package app

import (
	"github.com/graphql-go/graphql"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

type UserDelegate struct {
	*item.DefaultDelegate
}

func (d *UserDelegate) Name() string {
	return "users"
}

func (d *UserDelegate) BuildItem() *graphql.List {
	return graphql.NewList(graphql.NewObject(
		graphql.ObjectConfig{

			Name:   d.Name(),
			Fields: d.BuildField(),
		},
	))
}

// 声明当前Item所拥有的字段，以及各自字段的的Resolve
func (d *UserDelegate) BuildField() graphql.Fields {
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
