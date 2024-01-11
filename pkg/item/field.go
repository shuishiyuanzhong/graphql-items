package item

import "github.com/graphql-go/graphql"

type Field struct {
	fieldName string
	fieldType string

	resolver graphql.FieldResolveFn
}

func (f *Field) Convert() *graphql.Field {
	return &graphql.Field{
		Name:    f.fieldName,
		Type:    nil, // f.aliasType 根据这个值从Hub加载相应的Output
		Resolve: f.resolver,
	}
}

func (f *Field) SetResolver(resolver graphql.FieldResolveFn) {
	f.resolver = resolver
}

func (f *Field) Resolver() graphql.FieldResolveFn {
	return f.resolver
}

func NewItemField(fieldName, fieldType string) *Field {
	return &Field{
		fieldName: fieldName,
		fieldType: fieldType,
	}
}
