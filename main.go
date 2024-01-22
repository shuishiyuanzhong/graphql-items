package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/shuishiyuanzhong/graphql-items/app"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

func InitGraphQL() (*graphql.Schema, error) {
	item.Hub = new(item.ItemHub)
	item.Hub.Register(app.NewUserDelegate())
	item.Hub.Register(app.NewDepartmentDelegate())

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
		Schema: schema,
		Pretty: true,
	})
	http.Handle("/graphql", httpHandler)
	http.ListenAndServe(":8080", nil)
}
