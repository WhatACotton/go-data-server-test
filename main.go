package main

import (
	"unify/internal/database"
	"unify/internal/handler"
	"unify/validation"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	validation.CORS(r)
	database.TestSQL()
	r.GET("/item/top", handler.Top)
	r.GET("/item/all", handler.ALL)
	r.GET("/item/details", handler.Item_Details)
	r.GET("/item/category/:category", handler.Category)
	r.Run(":8080") // 0.0.0.0:8080 でサーバーを立てます。
}
