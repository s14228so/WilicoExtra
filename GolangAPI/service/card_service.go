package service

import (
	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

// Service procides user's behavior

// User is alias of entity.User struct
type Card entity.Card

func (s Service) CreateCardModel(userID string, c *gin.Context) (Card, error) {

	db := db.GetDB()
	var u Card

	if err := db.Where("userID = ?", userID).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) UpdateCardByID(userID string, c *gin.Context) (Card, error) {
	db := db.GetDB()
	var u Card

	if err := db.Where("userID = ?", userID).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User
func (s Service) DeleteCardByID(userID string) error {
	db := db.GetDB()
	var u Card

	if err := db.Where("userID = ?", userID).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
