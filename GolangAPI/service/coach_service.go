package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

type Coach entity.Coach

type Plans []entity.Plan

func (s Service) GetCoachAll() ([]Coach, error) {
	db := db.GetDB()
	var u []Coach

	//Coach型のレコードを全取得
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	for i, s := range u {
		//前半だけなら一つのコーチに所属するプラン一覧を取得できる
		//後半でcoachとplansをくっつけてる
		if err := db.Model(&s).Related(&s.Plans).Find(&u[i].Plans).Error; err != nil {
			log.Fatal(err)
		}
	}

	return u, nil
}

func (s Service) CreateCoachModel(c *gin.Context) (Coach, error) {
	db := db.GetDB()
	var u Coach
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetCoachByID(id string) (Coach, error) {
	db := db.GetDB()
	var u Coach

	if err := db.First(&u, id).Error; err != nil {
		log.Fatal(err)
	}
	if err := db.Model(&u).Related(&u.Plans).Error; err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func (s Service) UpdateCoachByID(id string, c *gin.Context) (Coach, error) {
	db := db.GetDB()
	var u Coach

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

func (s Service) DeleteCoachByID(id string) error {
	db := db.GetDB()

	var u Coach

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
