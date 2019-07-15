package service

// user周りのメソッド。DBとの通信周り

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
)

// Service procides user's behavior

// User is alias of entity.User struct

// GetAll is get all User
func (s Service) GetSubscribeAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	// for i, s := range u {
	// 	//前半だけなら一つのコーチに所属するプラン一覧を取得できる
	// 	//後半でcoachとplansをくっつけてる
	// 	if err := db.Model(&s).Related(&s.Card).Find(&u[i].Card).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// for _, v := range u {
	// 	if err := db.Model(&v).Related(&card).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	return u, nil
}

type Subscribe struct {
	UserID uint `json:"userid"`
	PlanID uint `json:"planid"`
}

// CreateModel is create User model
func (s Service) CreateSubscribeModel(userid string, planid string, c *gin.Context) (string, error) {
	db := db.GetDB()

	// var user User
	var plan Plan

	type Subscription struct {
		UserID uint
		PlanID uint
	}
	var subscription Subscription
	var user User
	log.Printf("planid: %v", planid)
	log.Printf("userid: %v", userid)
	u, _ := strconv.Atoi(userid)
	p, _ := strconv.Atoi(planid)

	user.ID = uint(u)
	plan.ID = uint(p)
	subscription.UserID = uint(u)
	subscription.PlanID = uint(p)
	log.Printf("plan: %v", plan)

	if err := db.Create(&subscription).Error; err != nil {
		return "error", err
	}

	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		panic("Could not find the user!")
	}

	if err := db.Where("id = ?", planid).First(&plan).Error; err != nil {
		return "error", err
	}

	db.Model(&user).Association("Plans").Append(&plan)

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
