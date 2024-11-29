package services

import (
	"fmt"
	"myapp/domain/item/models"
	"myapp/domain/item/repositories"
	"myapp/helpers"

	"gorm.io/gorm"
)

type itemService struct {
	ItemRepo repositories.ItemRepository
}

// Create implements ItemService.
func (service *itemService) Create(item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.ItemRepo.Create(item); err!=nil{
		response.Status = 500
		response.Message = "Failed to create new item " + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success create new item"
	}
	return response
}

// Delete implements ItemService.
func (service *itemService) Delete(itemID int) helpers.Response {
	var response helpers.Response
	if err := service.ItemRepo.Delete(itemID); err!=nil{
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete item : ", itemID)
	} else {
		response.Status = 200
		response.Message = "Success delete item"
	}
	return response
}

// GetAll implements ItemService.
func (service *itemService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.ItemRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Message = "Failed to get all items"
	} else {
		response.Status = 200
		response.Message = "Success to get all items"
		response.Data = data
	}
	return response
}

// GetById implements ItemService.
func (service *itemService) GetById(itemID int) helpers.Response {
	var response helpers.Response
	data, err := service.ItemRepo.GetById(itemID)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to get item : ", itemID)
	} else {
		response.Status = 200
		response.Message = "Success to get item"
		response.Data = data
	}
	return response
}

// Update implements ItemService.
func (service *itemService) Update(itemID int, item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.ItemRepo.Update(itemID, item); err!=nil{
		response.Status = 500
		response.Message = fmt.Sprint("Failed to update item : ", itemID)
	} else {
		response.Status = 200
		response.Message = "Success update item"
	}
	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(itemID int, item models.Item) helpers.Response
	Delete(itemID int) helpers.Response
	GetById(itemID int) helpers.Response
	GetAll() helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{ItemRepo: repositories.NewItemRepository(db)}
}
