package service

// user周りのメソッド。DBとの通信周り

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
)

// Service procides user's behavior

// User is alias of entity.User struct

// GetAll is get all User

type Subscription struct {
	UserID uint `json:"userid"`
	PlanID uint `json:"planid"`
}

func (s Service) GetSubscribeAll() ([]Subscription, error) {
	db := db.GetDB()
	var u []Subscription

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil

	// db := db.GetDB()
	// var u []User

	// // if err := db.Find(&u).Error; err != nil {
	// // 	return nil, err
	// // }

	// // for i, s := range u {
	// // 	//前半だけなら一つのコーチに所属するプラン一覧を取得できる
	// // 	//後半でcoachとplansをくっつけてる
	// // 	if err := db.Model(&s).Related(&s.Card).Find(&u[i].Card).Error; err != nil {
	// // 		log.Fatal(err)
	// // 	}
	// // }

	// // for _, v := range u {
	// // 	if err := db.Model(&v).Related(&card).Error; err != nil {
	// // 		log.Fatal(err)
	// // 	}
	// // }

	// return u, nil
}

// CreateModel is create User model
func (s Service) CreateSubscribeModel(c *gin.Context) (string, error) {
	db := db.GetDB()

	var user User
	var plan Plan

	type Subscription struct {
		UserID uint
		PlanID uint
	}
	var subscription Subscription
	// var user User

	// 	func (s Service) CreatePlanModel(c *gin.Context) (Plan, error) {
	// 	db := db.GetDB()
	// 	var u Plan

	if err := c.BindJSON(&subscription); err != nil {
		return "ok", err
	}

	if err := db.Create(&subscription).Error; err != nil {
		return "ok", err
	}

	if err := db.Where("id = ?", subscription.UserID).First(&user).Error; err != nil {
		panic("Could not find the user!")
	}
	log.Printf("user: %v", user)

	log.Printf("subscription: %v", subscription.UserID)
	if err := db.Where("id = ?", subscription.PlanID).First(&plan).Error; err != nil {
		return "error", err
	}
	log.Printf("plan: %v", plan)

	db.Model(&user).Association("Plans").Append(&plan)
	db.Model(&plan).Association("Users").Append(&user)
	var plans []Plan
	db.First(&user, "id = ?", 1)
	db.Model(&user).Related(&plans, "Plans")

	return "OK!", nil
}

// GetByID is get a User
func (s Service) GetSubscribeByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	// var card Card

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := db.Model(&u).Related(&u.Card).Error; err != nil {
		log.Fatal(err)
	}

	return u, nil
}

// UpdateByID is update a User
func (s Service) UpdateSubscribeByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

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
func (s Service) DeleteSubscribeByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
