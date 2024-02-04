package item

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
)

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

		rows, err := HUB().GetDB().QueryContext(context.Background(), sql)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		columns, _ := rows.Columns()

		cache := make([]interface{}, len(columns)) //临时存储每行数据
		for index, _ := range cache {              //为每一列初始化一个指针
			var a interface{}
			cache[index] = &a
		}
		var list []map[string]interface{} //返回的切片
		for rows.Next() {
			_ = rows.Scan(cache...)

			item := make(map[string]interface{})
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
			list = append(list, item)
		}

		return list, nil
		//return []map[string]interface{}{{"id": "1", "name": "Example Product", "price": 99.99}}, nil
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
	Float  ColumnType = "Float"
	String ColumnType = "String"
)
