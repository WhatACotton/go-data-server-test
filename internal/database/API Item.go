package database

import (
	"github.com/pkg/errors"
)

type Top_Item struct {
	Item_Name string `json:"Item_Name"`
	Stock     int    `json:"Stock"`
	Order     int    `json:"Order"`
}
type Top_Item_List []Top_Item

func Get_Top() (Top_Item_List, error) {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Item_Info.Name,Item_Info.Stock,Item_List.Item_Order From Item_Info JOIN ON Item_List.Info_ID = Item_Info.Info_ID Where Item_Info.Top = 1 AND Item_List.Status = 'Available'")
	if err != nil {
		return nil, errors.Wrap(err, "error in getting Top_Item /Top_1")
	}
	var return_Item []Top_Item
	for rows.Next() {
		Top_Item := new(Top_Item)
		err := rows.Scan(&Top_Item.Item_Name, &Top_Item.Stock, &Top_Item.Order)
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
	Order     int    `json:"Order"`
}
type Item_List []Item

func Get_ALL() (Item_List, error) {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Item_List.Item_Order,Item_Info.Name,Item_Info.Stock,Item_Info.Price,Item_List.Status FROM Item_List JOIN Item_Info ON Item_List.Info_ID = Item_Info.Info_ID Where Item_List.Status = 'Available'")
	if err != nil {
		return nil, errors.Wrap(err, "error in getting Top_Item /Get_ALL_1")
	}
	var return_Item []Item
	for rows.Next() {
		Item := new(Item)
		err := rows.Scan(&Item.Order, &Item.Item_Name, &Item.Stock, &Item.Price, &Item.Status)
		if err != nil {
			return nil, errors.Wrap(err, "error in scanning Cart_ID /Get_ALL_2")
		}
		return_Item = append(return_Item, *Item)
	}
	return return_Item, nil
}

func Get_Item_Category(Category string) (Item_List, error) {
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Item_List.Item_Order,Item_Info.Name,Item_Info.Stock,Item_Info.Price,Item_List.Status FROM Item_List JOIN Item_Info ON Item_List.Info_ID = Item_Info.Info_ID Where Item_List.Status = 'Available' AND Category = ?", Category)
	if err != nil {
		return nil, errors.Wrap(err, "error in getting Top_Item /Get_Category_1")
	}
	var return_Item []Item
	for rows.Next() {
		Item := new(Item)
		err := rows.Scan(&Item.Order, &Item.Item_Name, &Item.Stock, &Item.Price, &Item.Status)
		if err != nil {
			return nil, errors.Wrap(err, "error in scanning Cart_ID /Get_Category_2")
		}
		return_Item = append(return_Item, *Item)
	}
	return return_Item, nil
}

func Get_Info_Id(Item_ID string) (Info_ID string, err error) {
	db := ConnectSQL()
	defer db.Close()
	Status := new(string)
	// SQLの実行
	rows, err := db.Query("SELECT Info_ID,Status FROM Item_List WHERE Item_ID = ?", Item_ID)
	if err != nil {
		return "error", errors.Wrap(err, "error in getting Top_Item /Recommend_1")
	}
	for rows.Next() {
		err := rows.Scan(&Info_ID, &Status)
		if err != nil {
			return "error", errors.Wrap(err, "error in scanning Cart_ID /Recomend_2")
		}
	}
	if *Status == `Available` {
		return Info_ID, nil
	} else {
		return "Couldn't get", nil
	}
}

type Item_Details struct {
	Item_ID     string `json:"Item_ID"`
	Item_Name   string `json:"Item_Name"`
	Stock       int    `json:"Stock"`
	Price       int    `json:"Price"`
	Color       string `json:"Color"`
	Description string `json:"Description"`
	Category    string `json:"category"`
	Key_Words   string `json:"Key_Words"`
	Status      string `json:"Status"`
}

func Get_Item_Details(Info_ID string) (Item_Details, error) {
	Item_Details := new(Item_Details)
	db := ConnectSQL()
	defer db.Close()
	// SQLの実行
	rows, err := db.Query("SELECT Info_ID,Name,Price,Stock,Color,Category,Key_Words,Description From Item_Info WHERE Info_ID = ?", Info_ID)
	if err != nil {
		return *Item_Details, errors.Wrap(err, "error in getting Top_Item /Recommend_1")
	}
	for rows.Next() {
		err := rows.Scan(&Item_Details.Item_Name, &Item_Details.Price, &Item_Details.Stock, &Item_Details.Color, &Item_Details.Category, &Item_Details.Key_Words, &Item_Details.Description)
		if err != nil {
			return *Item_Details, errors.Wrap(err, "error in scanning Cart_ID /Recomend_2")
		}
	}
	return *Item_Details, nil
}
