package models

type Item_Details struct {
	Item_ID     string `json:"Item_ID"`
	Item_Name   string `json:"Item_Name"`
	Stock       int    `json:"Stock"`
	Price       int    `json:"Price"`
	Description string `json:"Description"`
	Category    string `json:"category"`
	Key_Word    string `json:"Key_Word"`
	Status      string `json:"Status"`
}
