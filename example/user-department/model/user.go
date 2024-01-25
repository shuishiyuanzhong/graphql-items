package model

import (
	"github.com/graphql-go/graphql"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

const (
	FieldTypeUser = "user"
)

type UserDelegate struct {
	item.SqlHelper
}

var _ item.Delegate = &UserDelegate{}

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

func (d *UserDelegate) Type() item.FieldType {
	return FieldTypeUser
}

func (d *UserDelegate) IsList() bool {
	return true
}

func (d *UserDelegate) BuildField() []*item.Field {
	fields := make([]*item.Field, 0)

	fields = append(fields,
		item.NewItemField("id", item.FieldTypeString),
	)
	fields = append(fields,
		item.NewItemField("name", item.FieldTypeString),
	)
	fields = append(fields,
		item.NewItemField("price", item.FieldTypeFloat),
	)

	fields = append(fields, d.departmentField())
	return fields
}

func (d *UserDelegate) departmentField() *item.Field {
	field := item.NewItemField("department", FieldTypeDepartment)
	field.SetResolver(func(p graphql.ResolveParams) (interface{}, error) {
		return []map[string]interface{}{{"id": "1", "name": "Example Product", "price": 99.99}}, nil
	})
	return field
}
