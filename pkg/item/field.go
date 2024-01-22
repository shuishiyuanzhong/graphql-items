package item

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type Field struct {
	fieldName string
	fieldType FieldType

	resolver graphql.FieldResolveFn
}

func (f *Field) Convert() (*graphql.Field, error) {
	// TODO 从Hub中加载FieldType，然后它们返回的type可能就是graphql框架原有的类型以及用户自己声明的item类型
	t, ok := FieldTypeMapping[f.fieldType]
	if !ok {
		return nil, fmt.Errorf("unsupported field type: %s", f.fieldType)
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
