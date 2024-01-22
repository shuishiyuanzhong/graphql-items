package item

import "github.com/graphql-go/graphql"

type FieldType string

const (
	FieldTypeString  FieldType = "String"
	FieldTypeInt               = "Int"
	FieldTypeFloat             = "Float"
	FieldTypeBoolean           = "Boolean"
)

// TODO 这个结构应该维护在Hub中, 对应的校验逻辑也不应该放到NewField阶段
var FieldTypeMapping = map[FieldType]graphql.Output{
	FieldTypeString:  graphql.String,
	FieldTypeInt:     graphql.Int,
	FieldTypeFloat:   graphql.Float,
	FieldTypeBoolean: graphql.Boolean,
}
