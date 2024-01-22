package item

import (
	"github.com/graphql-go/graphql"
)

type Field struct {
	fieldName string
	fieldType FieldType

	asList bool

	resolver graphql.FieldResolveFn
}

func (f *Field) Convert(hub *ItemHub) (*graphql.Field, error) {
	// TODO 从Hub中加载FieldType，然后它们返回的type可能就是graphql框架原有的类型以及用户自己声明的item类型
	t, err := hub.loadFieldType(f.fieldType)
	if err != nil {
		return nil, err
	}

	return &graphql.Field{
		Name:    f.fieldName,
		Type:    t,
		Resolve: f.resolver,
	}, nil
}

func (f *Field) SetResolver(resolver graphql.FieldResolveFn) {
	f.resolver = resolver
}

func (f *Field) Resolver() graphql.FieldResolveFn {
	return f.resolver
}

func NewItemField(fieldName string, fieldType FieldType) *Field {
	return &Field{
		fieldName: fieldName,
		fieldType: fieldType,
		resolver:  nil,
	}
}
