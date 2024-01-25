package model

import (
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

const (
	FieldTypeDepartment = "department"
)

type DepartmentDelegate struct {
	item.SqlHelper
}

var _ item.Delegate = (*DepartmentDelegate)(nil)

func (d *DepartmentDelegate) Name() string {
	return "departments"
}

func (d *DepartmentDelegate) Type() item.FieldType {
	return FieldTypeDepartment
}

func (d *DepartmentDelegate) BuildField() []*item.Field {
	fields := make([]*item.Field, 0)

	fields = append(fields,
		item.NewItemField("id", item.FieldTypeString),
	)
	fields = append(fields,
		item.NewItemField("name", item.FieldTypeString),
	)
	fields = append(fields,
		item.NewItemField("test", item.FieldTypeString),
	)
	return fields
}

func (d *DepartmentDelegate) IsList() bool {
	return true
}

func NewDepartmentDelegate() (d *DepartmentDelegate) {
	d = &DepartmentDelegate{}
	d.SqlHelper = item.NewDefaultSqlHelper("department", d.initItemTable())
	return
}

func (d *DepartmentDelegate) initItemTable() []*item.Column {
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
