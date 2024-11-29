package models

type Item struct{
	ItemID int `json:"item_id" gorm:"column:item_id;primaryKey;autoIncrement"`
	ItemName string `json:"item_name" gorm:"column:item_name"`
	Unit string `json:"unit" gorm:"column:unit"`
	Stock int `json:"stock" gorm:"column:stock"`
	Price float64 `json:"price" gorm:"column:price"`
}

func (Item) TableName() string {
	return "item" 
}