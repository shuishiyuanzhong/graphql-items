package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/shuishiyuanzhong/graphql-items/app"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
	"net/http"
)

// 定义GraphQL Schema
func createSchema() graphql.Schema {

	// 定义Product类型
	productType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Product",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.String},
				"name":  &graphql.Field{Type: graphql.String},
				"price": &graphql.Field{Type: graphql.Float},
				"test": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "test", nil
					},
				},
			},
		},
	)

	// 定义User类型
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.String},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
				"product": &graphql.Field{
					Type:    productType,
					Resolve: productResolver,
				},
			},
		},
	)

	// 定义Query类型
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:    userType,
					Resolve: userResolver,
				},
				"product": &graphql.Field{
					Type:    productType,
					Resolve: productResolver,
				},
			},
		},
	)

	// 创建并返回Schema
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	return schema
}

// User 定义用户数据模型
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product 定义产品数据模型
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// userResolver 解析器：返回一个用户
func userResolver(p graphql.ResolveParams) (interface{}, error) {
	// 在实际应用中，您可能需要从数据库或其他数据源获取用户数据
	return User{ID: "1", Name: "John Doe", Email: "john@example.com"}, nil
}

// productResolver 解析器：返回一个产品
func productResolver(p graphql.ResolveParams) (interface{}, error) {
	// 在实际应用中，您可能需要从数据库或其他数据源获取产品数据
	return Product{ID: "1", Name: "Example Product", Price: 99.99}, nil
}

func InitGraphQL() (graphql.Schema, error) {
	item.Hub = new(item.ItemHub)
	item.Hub.Register(app.NewUserDelegate())

	return item.Hub.BuildSchema()
}

func main() {
	// 定义Schema
	//schema := createSchema()

	schema, err := InitGraphQL()
	if err != nil {
		panic(err)
	}

	// 设置HTTP服务器
	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", httpHandler)
	http.ListenAndServe(":8080", nil)
}
