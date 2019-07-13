package service

import (
	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

// Service procides user's behavior

// User is alias of entity.User struct
type Plan entity.Plan

// GetAll is get all User
func (s Service) GetPlanAll() ([]Plan, error) {
	db := db.GetDB()
	var u []Plan

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}
	for i, s := range u {
		db.Where("id = ?", s.CoachID).Find(&u[i].Coach)
	}

	return u, nil
}

// CreateModel is create User model
func (s Service) CreatePlanModel(c *gin.Context) (Plan, error) {
	db := db.GetDB()
	var u Plan

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetPlanByID(id string) (Plan, error) {
	db := db.GetDB()
	var u Plan

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := db.Model(&u).Related(&u.Coach).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (s Service) UpdatePlanByID(id string, c *gin.Context) (Plan, error) {
	db := db.GetDB()
	var u Plan

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User
func (s Service) DeletePlanByID(id string) error {
	db := db.GetDB()
	var u Plan

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
