package main

import (
	"github.com/shuishiyuanzhong/graphql-items/conf"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/shuishiyuanzhong/graphql-items/example/user-department/model"
	"github.com/shuishiyuanzhong/graphql-items/pkg/item"
)

func InitGraphQL() (*graphql.Schema, error) {
	item.HUB().Register(model.NewUserDelegate())
	item.HUB().Register(model.NewDepartmentDelegate())

	item.HUB().SetDB(conf.C().Mysql.GetDB())
	return item.HUB().BuildSchema()
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
