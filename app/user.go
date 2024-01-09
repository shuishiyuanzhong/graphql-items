package app

import (
	"github.com/graphql-go/graphql"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

type UserDelegate struct {
	item.SqlHelper
}

func NewUserDelegate() (d *UserDelegate) {
	d = &UserDelegate{}
	d.SqlHelper = item.NewDefaultSqlHelper("user", d.initItemTable())
	return
}

func (d *UserDelegate) initItemTable() []*item.Column {
	columns := make([]*item.Column, 0)
	columns = append(columns, &item.Column{
		Type:  "",
		Name:  "id",
		Alias: "id",
	})
	columns = append(columns, &item.Column{
		Type:  "",
		Name:  "name",
		Alias: "name",
	})
	columns = append(columns, &item.Column{
		Type:  "",
		Name:  "email",
		Alias: "email",
	})

	return columns
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
