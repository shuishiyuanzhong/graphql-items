package item

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type Delegate interface {
	Name() string
	// BuildItem 生成GraphQL Object
	BuildItem() *graphql.List
	Resolve() graphql.FieldResolveFn
}

type SqlHelper interface {
	Resolve() graphql.FieldResolveFn
}

type DefaultSqlHelper struct {
	tableName      string
	tableColumns   []*Column
	columnsByAlias map[string]*Column
	columnsByName  map[string]*Column
}

func NewDefaultSqlHelper(tableName string, columns []*Column) *DefaultSqlHelper {
	d := &DefaultSqlHelper{
		tableName:      tableName,
		tableColumns:   make([]*Column, 0, len(columns)),
		columnsByAlias: make(map[string]*Column),
		columnsByName:  make(map[string]*Column),
	}

	for _, column := range columns {
		d.tableColumns = append(d.tableColumns, column)
		d.columnsByAlias[column.Alias] = column
		d.columnsByName[column.Name] = column
	}

	return d
}

func (d *DefaultSqlHelper) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// TODO * 应该替换成具体的列表
		sql := "SELECT * from %s"
		sql = fmt.Sprintf(sql, d.tableName)

		// TODO query from db

		return []map[string]interface{}{{"id": "1", "name": "Example Product", "price": 99.99}}, nil
	}
}

// DTO to GraphQL Object

type Column struct {
	Type  ColumnType
	Name  string
	Alias string
}

type ColumnType string

const (
	Int    ColumnType = "Int"
	Float             = "Float"
	String            = "String"
)
