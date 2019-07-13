package entity

type Card struct {
	Brand string `json:"brand"`
	Last4 int    `json:"last4"`
	Month int    `json:"month"`
	Src   string `json:"src"`
	Year  int    `json:"year"`
}

// User is user models property
type User struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	UserName     string `json:"username"`
	Introduction string `json:"introduction"`
	Age          string `json:"age"`
	Coach        bool   `json:"coach"`
	Card         Card   `json:"card"`
}

type Plan struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Price   int    `json:"price"`
	CoachId int    `json:"coachid"`
}

type Coach struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Age      string `json:"age"`
	Coach    bool   `json:"coach"`
	Rating   int    `json:"rating"`
}
