package service

// user周りのメソッド。DBとの通信周り

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/db"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct

// GetAll is get all User
func (s Service) GetUserAll() ([]User, error) {
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

// CreateModel is create User model
func (s Service) CreateUserModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetUserByID(id string) (User, error) {
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
func (s Service) UpdateUserByID(id string, c *gin.Context) (User, error) {
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
func (s Service) DeleteUserByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
