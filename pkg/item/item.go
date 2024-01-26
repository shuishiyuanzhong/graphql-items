package item

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/shuishiyuanzhong/graphql-items/conf"
)

type SqlHelper interface {
	Resolve() graphql.FieldResolveFn
}

type DefaultSqlHelper struct {
	tableName      string
	tableColumns   []*Column
	columnsByAlias map[string]*Column
	columnsByName  map[string]*Column

	db *sql.DB
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

	d.db = conf.C().Mysql.GetDB()
	return d
}

func (d *DefaultSqlHelper) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// TODO * 应该替换成具体的列表
		sql := "SELECT * from %s"
		sql = fmt.Sprintf(sql, d.tableName)

		// TODO query from db

		rows, err := d.db.QueryContext(context.Background(), sql)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		data := make([]map[string]interface{}, 0)
		for rows.Next() {
			ins := make(map[string]interface{})

			rows.Scan(&ins)

			data = append(data, ins)
		}
		return data, nil
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
