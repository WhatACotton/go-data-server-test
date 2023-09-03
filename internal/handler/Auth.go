package handler

import (
	"unify/internal/database"

	"github.com/gin-gonic/gin"
)

func Register_Item(c *gin.Context) {
	Item := new(database.Item_Register_Payload)
	Item.Get_Info(c)
	database.Register_Item(*Item)
	//TODO
}
func Update_Item(c *gin.Context) {
	//TODO
}
func Detach_Item(c *gin.Context) {
	//TODO
}
func Attach_Item(c *gin.Context) {
	//TODO
}
func Update_Item_List(c *gin.Context) {
	//TODO
}
