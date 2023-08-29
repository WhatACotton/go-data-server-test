package database

import (
	"github.com/pkg/errors"
)

type Top_Item struct {
	Item_Name string `json:"Item_Name"`
	Stock     int    `json:"Stock"`
}
type Top_Item_List []Top_Item

func Top() (Top_Item_List, error) {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Info_Name,Stock From Item_Info Where Top = 1")
	if err != nil {
		return nil, errors.Wrap(err, "error in getting Top_Item /Top_1")
	}
	var return_Item []Top_Item
	for rows.Next() {
		Top_Item := new(Top_Item)
		err := rows.Scan(&Top_Item.Item_Name, &Top_Item.Stock)
		if err != nil {
			return nil, errors.Wrap(err, "error in scanning Cart_ID /Top_2")
		}
		return_Item = append(return_Item, *Top_Item)
	}
	return return_Item, nil
}

type Item struct {
	Item_ID   string `json:"Item_ID"`
	Item_Name string `json:"Item_Name"`
	Stock     int    `json:"Stock"`
	Price     int    `json:"Price"`
	Status    string `json:"Status"`
}
type Item_List []Item

func Recommend() (Item_List, error) {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Item_Info.Info_Name,Item_Info.Stock,Item_Info.Price,Item_List.Status FROM Item_List JOIN Item_Info ON Item_List.Info_ID = Item_Info.Info_ID Where Item_List.Status = Available")
	if err != nil {
		return nil, errors.Wrap(err, "error in getting Top_Item /Recommend_1")
	}
	var return_Item []Item
	for rows.Next() {
		Item := new(Item)
		err := rows.Scan(&Item.Item_Name, &Item.Stock, &Item.Price, &Item.Status)
		if err != nil {
			return nil, errors.Wrap(err, "error in scanning Cart_ID /Recomend_2")
		}
		return_Item = append(return_Item, *Item)
	}
	return return_Item, nil
}
