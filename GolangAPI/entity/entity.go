package entity

import "github.com/jinzhu/gorm"

type Card struct {
	Brand string `json:"brand"`
	Last4 int    `json:"last4"`
	Month int    `json:"month"`
	Src   string `json:"src"`
	Year  int    `json:"year"`
}

// User is user models property
type User struct {
	gorm.Model
	Email        string `json:"email"`
	UserName     string `json:"username"`
	Introduction string `json:"introduction"`
	Age          string `json:"age"`
	Coach        bool   `json:"coach"`
	Card         Card   `json:"card"`
}

type Plan struct {
	gorm.Model
	Title   string `json:"title"`
	Price   int    `json:"price"`
	CoachID uint   `json:"coachid"`
	// Coach   Coach  `gorm:"association_autoupdate:false"`
	Users []User `gorm:"many2many:plan_users"`
}

type Coach struct {
	ID       uint
	Email    string `json:"email"`
	UserName string `json:"username"`
	Age      string `json:"age"`
	Coach    bool   `json:"coach"`
	Rating   int    `json:"rating"`
	Plans    []Plan `json:"plans"`
}

//plan作ってるのにコーチモデルが勝手に入る
//Coachオブゼクトを抜いたら勝手にはくなった

//https://blog.linkbal.co.jp/2704/
