package entity

import (
	"time"
)

type MainColumn struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Card struct {
	Brand  string `json:"brand"`
	Last4  int    `json:"last4"`
	Month  int    `json:"month"`
	Src    string `json:"src"`
	Year   int    `json:"year"`
	UserID uint   `json:"userid"`
	User   User   `json:"user"`
}

// User is user models property
type User struct {
	MainColumn
	Email        string `json:"email"`
	UserName     string `json:"username"`
	Introduction string `json:"introduction"`
	Age          string `json:"age"`
	Coach        bool   `json:"coach"`
	Card         *Card  `json:"card"`
}

type Plan struct {
	MainColumn
	Title   string `json:"title"`
	Price   int    `json:"price"`
	CoachID uint   `json:"coachid"`
	Coach   Coach  `gorm:"association_autoupdate:false" json:"coach"`
	Users   []User `gorm:"many2many:plan_users"`
}

type Coach struct {
	MainColumn
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
