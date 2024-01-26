package main

import (
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

func main() {
	// 定义Schema
	//schema := createSchema()

	schema, err := item.InitGraphQL()
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
