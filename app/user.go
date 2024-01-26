package app

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

func NewUserDelegate() (d *UserDelegate) {
	d = &UserDelegate{}
	d.SqlHelper = item.NewDefaultSqlHelper(d.Name(), d.initItemTable())
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

func (d *UserDelegate) BuildItem(fields graphql.Fields) *graphql.List {
	return graphql.NewList(graphql.NewObject(
		graphql.ObjectConfig{
			Name:   d.Name(),
			Fields: fields,
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
	// TODO 当前的item依赖于别的item的时候，需要重新设计一下这里。BuildField目前并没有足够的上下文去处理好这个问题
	fields["test"] = &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "test", nil
		},
	}
	/** TODO 如果说这个delegate中有一个字段对应了其他的delegate，即
	users{
		id
		name
		department{
			id
			name
		}
	}
	这种情况下，这个BuildField方法所拥有的上下文并不足够创建一个department field。
	这个方法的职责，应该只是声明delegate所拥有的字段，以及各自字段的Resolve，但不负责输出graphql.Field。
	输出的东西应该是一个抽象的对象，由ItemHub来负责将这些抽象的对象转换成graphql.Field。
	*/

	return fields
}

func (d *UserDelegate) BuildField2() []*item.Field {
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
	fields = append(fields,
		item.NewItemField("test", item.FieldTypeString),
	)
	return fields
}
