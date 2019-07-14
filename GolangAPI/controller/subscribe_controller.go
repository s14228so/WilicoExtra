package controller

//serverにレスポンスを返すところ

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/service"
)

// Controller is user controlller
type SubscribeController struct{}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	var s service.Service
	p, err := s.GetSubscribeAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateSubscribeModel(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetSubscribeByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdateSubscribeByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
