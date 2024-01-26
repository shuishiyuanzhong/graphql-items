package item

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"github.com/shuishiyuanzhong/graphql-items/app"
	"github.com/shuishiyuanzhong/graphql-items/conf"
)

var (
	Hub *ItemHub
)

type ItemHub struct {
	delegates []Delegate

	DB *sql.DB
}

func InitGraphQL() (*graphql.Schema, error) {
	Hub = new(ItemHub)
	Hub.DB = conf.C().Mysql.GetDB()

	Hub.Register(app.NewUserDelegate())

	return Hub.BuildSchema()
}

func (h *ItemHub) Register(delegate Delegate) {
	h.delegates = append(h.delegates, delegate)
}

func (h *ItemHub) BuildSchema() (*graphql.Schema, error) {
	fields := make(graphql.Fields)
	for _, delegate := range h.delegates {
		item, err := h.buildItem(delegate)
		if err != nil {
			return nil, err
		}

		fields[delegate.Name()] = item
	}

	// 生成schema(逻辑不变)
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		},
	)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

// 一个field依赖其他delegate对象，就发生在这里
func (h *ItemHub) initItemField(delegate Delegate) (graphql.Fields, error) {
	rawFields := delegate.BuildField2()
	fields := make(graphql.Fields)
	for _, f := range rawFields {
		convert, err := f.Convert()
		if err != nil {
			return nil, err
		}
		fields[f.fieldName] = convert
	}
	return fields, nil
}

// 这个方法最终应该输出一个能够被Hub直接使用的field字段
func (h *ItemHub) buildItem(delegate Delegate) (*graphql.Field, error) {
	fields, err := h.initItemField(delegate)
	if err != nil {
		return nil, err
	}

	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   delegate.Name(),
		Fields: fields,
	})

	var result graphql.Output
	result = obj
	if delegate.IsList() {
		result = graphql.NewList(obj)
	}

	return &graphql.Field{
		Type:    result,
		Resolve: delegate.Resolve(),
	}, nil
}
