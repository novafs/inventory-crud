package repositories

import (
	"myapp/domain/item/models"

	"gorm.io/gorm"
)

type dbItem struct {
	Conn *gorm.DB
}


func (db *dbItem) Create(item models.Item) error {
	return db.Conn.Create(&item).Error
}

func (db *dbItem) Delete(itemID int) error {
	return db.Conn.Delete(&models.Item{ItemID: itemID}).Error
}

func (db *dbItem) GetAll() ([]models.Item, error) {
	var data []models.Item
	result := db.Conn.Find(&data)

	return data, result.Error
}

func (db *dbItem) GetById(itemID int) (models.Item, error) {
	var data models.Item

	result:= db.Conn.Where("item_id", itemID).First(&data)
	return data, result.Error
}

func (db *dbItem) Update(itemID int, item models.Item) error {
	return db.Conn.Where("item_id", itemID).Updates(item).Error
}

type ItemRepository interface{
	Create(item models.Item) error
	Update(itemID int, item models.Item) error
	Delete(itemID int) error
	GetById(itemID int) (models.Item, error)
	GetAll() ([]models.Item, error)
}

func NewItemRepository(Conn *gorm.DB) ItemRepository {
	return &dbItem{Conn: Conn}
}