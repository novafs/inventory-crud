package controllers

import (
	"myapp/domain/item/models"
	"myapp/domain/item/services"
	"net/http"
	"strconv"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemController struct {
	itemService services.ItemService
	validate    vl.Validate
}

func (controller ItemController) Create(c echo.Context) error {
	type payload struct{
		ItemName string `json:"item_name" validate:"required"`
		Unit string `json:"unit" validate:"required"`
		Stock int `json:"stock" validate:"required"`
		Price float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err:=c.Bind(payloadValidator); err!=nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result:= controller.itemService.Create(models.Item{ItemName: payloadValidator.ItemName, Unit: payloadValidator.Unit, Stock: payloadValidator.Stock, Price: payloadValidator.Price})

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) Update(c echo.Context) error {
	type payload struct{
		ItemName  string  `json:"item_name" validate:"required"`
		Unit 	  string  `json:"unit" validate:"required"`
		Stock 	  int 	  `json:"stock" validate:"required"`
		Price 	  float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err:=c.Bind(payloadValidator); err!=nil {
		return err
	}

	itemID,_ := strconv.Atoi(c.Param("item_id"))
	result:= controller.itemService.Update(itemID, models.Item{ItemName: payloadValidator.ItemName, Unit: payloadValidator.Unit, Stock: payloadValidator.Stock, Price: payloadValidator.Price})

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) Delete(c echo.Context) error {
	itemID,_ := strconv.Atoi(c.Param("item_id"))
	result:= controller.itemService.Delete(itemID)

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) GetAll(c echo.Context) error {
	result:= controller.itemService.GetAll()

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) GetById(c echo.Context) error {
	itemID,_ := strconv.Atoi(c.QueryParam("item_id"))
	result:= controller.itemService.GetById(itemID)

	return c.JSON(http.StatusOK, result)
}

func NewItemController(db *gorm.DB) ItemController{
	service := services.NewItemService(db)

	controller := ItemController{
		itemService: service,
		validate:    *vl.New(),
	}
	return controller
}