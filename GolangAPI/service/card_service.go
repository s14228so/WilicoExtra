package service

import (
	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

// Service procides user's behavior

// User is alias of entity.User struct
type Card entity.Card

type User entity.User

func (s Service) CreateCardModel(c *gin.Context) (Card, error) {

	db := db.GetDB()
	var u Card

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetCardByID(id string) (Card, error) {
	db := db.GetDB()
	var u Card

	// var user User

	if err := db.Where("user_id = ?", id).Find(&u).First(&u).Error; err != nil {
		return u, err
	}

	// if err := db.Model(&u).Related(&u.User).Error; err != nil {
	// 	return u, err
	// }

	return u, nil
}

func (s Service) UpdateCardByID(userID string, c *gin.Context) (Card, error) {
	db := db.GetDB()
	var u Card

	if err := db.Where("user_id = ?", userID).First(&u).Error; err != nil {
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

	if err := db.Where("user_id = ?", userID).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
