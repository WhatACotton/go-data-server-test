package handler

import (
	"log"
	"unify/internal/database"

	"github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
	Top_Item_List, err := database.Top()
	if err != nil {
		log.Print(err)
	}
	c.JSON(200, Top_Item_List)
}
func Recommend(c *gin.Context) {
	Recommend_Item_List, err := database.Recommend()
	if err != nil {
		log.Print(err)
	}
	c.JSON(200, Recommend_Item_List)
}
func Item_Details(c *gin.Context) {
	// TODO
}
func Category(c *gin.Context) {
	// TODO
}
