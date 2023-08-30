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
	authorized := validation.Basic(r)
	r.GET("/item/top", handler.Top)
	r.GET("/item/all", handler.ALL)
	r.GET("/item/details", handler.Item_Details)
	r.GET("/item/category/:category", handler.Category)
	r.GET("/item/color/:color", handler.Color)

	authorized.POST("/register/item", handler.Register_Item)
	authorized.PATCH("/update/item", handler.Update_Item)
	authorized.POST("/update/detach", handler.Detach_Item)
	authorized.POST("/update/attach", handler.Attach_Item)
	authorized.POST("/update/Item_List", handler.Update_Item_List)
	r.Run(":8080") // 0.0.0.0:8080 でサーバーを立てます。
}
