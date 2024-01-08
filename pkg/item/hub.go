package item

import "github.com/graphql-go/graphql"

type ItemHub struct {
	delegates []Delegate
}

func (h *ItemHub) Register(delegate Delegate) {
	h.delegates = append(h.delegates, delegate)
}

func (h *ItemHub) BuildSchema() (graphql.Schema, error) {
	fields := make(graphql.Fields)
	for _, delegate := range h.delegates {
		fields[delegate.Name()] = &graphql.Field{
			Type:    delegate.BuildItem(),
			Resolve: delegate.Resolve(),
		}
	}
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		},
	)

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		})
}

var Hub *ItemHub

func InitGraphQL() (graphql.Schema, error) {
	Hub = new(ItemHub)
	Hub.Register(new(ItemDemo))

	return Hub.BuildSchema()
}
