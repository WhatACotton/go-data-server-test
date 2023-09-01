package database

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Item_Register_Payload struct {
	Info_ID     string `json:"Info_ID"`
	Price       int    `json:"Price"`
	Name        string `json:"Name"`
	Stock       int    `json:"Stock"`
	Color       string `json:"Color"`
	Category    string `json:"Category"`
	Key_Words   string `json:"Key_Words"`
	Description string `json:"Description"`
	Top         int    `json:"Top"`
}

func (Item *Item_Register_Payload) Get_Info(c *gin.Context) {
	err := c.BindJSON(&Item)
	if err != nil {
		c.JSON(400, "{Bad Request}")
	}

}

func Register_Item(Item Item_Register_Payload) error {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	_, err := db.Exec(
		`INSERT INTO Item_Info (
			Info_ID,
			Name,
			Price,
			Stock,
			Color,
			Category,
			Key_Words,
			Description,
			Top
		) VALUES (?,?,?,?,?,?,?,?,?)`,
		Item.Info_ID,
		Item.Name,
		Item.Price,
		Item.Stock,
		Item.Color,
		Item.Category,
		Item.Key_Words,
		Item.Description,
		Item.Top)
	if err != nil {
		return errors.Wrap(err, "error in getting Top_Item /Get_Info_ID_1")
	}
	return nil
}
