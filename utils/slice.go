package utils

import (
	"GoWebServer/models"
	"errors"
)

func FindItemByID(itemID int, items []models.Item) (models.Item, error) {

	for i := range items {
		if items[i].Id == itemID {
			return items[i], nil
		}
	}
	emptyItem := models.Item{}
	return emptyItem, errors.New("Not item found")

}
