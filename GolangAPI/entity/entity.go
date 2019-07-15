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

// User is user models property
type User struct {
	MainColumn
	Email        string     `json:"email"`
	UserName     string     `json:"username"`
	Introduction string     `json:"introduction"`
	Age          string     `json:"age"`
	Coach        bool       `json:"coach"`
	Card         Card       `json:"card"`
	Plans        []*Plan    `gorm:"many2many:subscriptions" json:"plans"`
	Favorites    []Favorite `json:"favorites"`
}

type Card struct {
	MainColumn
	Brand  string `json:"brand"`
	Last4  int    `json:"last4"`
	Month  int    `json:"month"`
	Src    string `json:"src"`
	Year   int    `json:"year"`
	UserID uint   `json:"userid"`
}

type Plan struct {
	MainColumn
	Title   string  `json:"title"`
	Price   int     `json:"price"`
	CoachID uint    `json:"coachid"`
	Coach   Coach   `gorm:"association_autoupdate:false" json:"coach"`
	Users   []*User `gorm:"many2many:subscriptions" json:"users"`
}

//plan.usersを作れば勝手にsubscribes作られる説
//でもuserを作らないといけないから違うか
type Coach struct {
	MainColumn
	Email    string `json:"email"`
	UserName string `json:"username"`
	Age      string `json:"age"`
	Coach    bool   `json:"coach"`
	Rating   int    `json:"rating"`
	Plans    []Plan `json:"plans"` //DBには入ってないもの
}

type Album struct {
	MainColumn
	Image     string     `json:"image"`
	UserID    uint       `json:"userid"`
	User      User       `json:"user"`
	Body      string     `json:"body"`
	Favorites []Favorite `json:"favorites"`
	Comments  []Comment  `json:"comments"`
}

type Comment struct {
	MainColumn
	AlbumID uint   `json:"albumid"`
	Album   Album  `json:"album"`
	Body    string `json:"body"`
	UserID  uint   `json:"userid"`
	User    User   `json:"user"`
}

type Favorite struct {
	MainColumn
	AlbumID uint `json:"albumid"`
	UserID  uint `json:"userid"`
	User    User `json:"user"`
}

//plan作ってるのにコーチモデルが勝手に入る
//Coachオブゼクトを抜いたら勝手にはくなった

//https://blog.linkbal.co.jp/2704/
