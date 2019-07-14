package db

import (
	"github.com/jinzhu/gorm"
	// Use PostgreSQL in gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Plan{})
	db.AutoMigrate(&entity.Coach{})
	db.AutoMigrate(&entity.Card{})
}
