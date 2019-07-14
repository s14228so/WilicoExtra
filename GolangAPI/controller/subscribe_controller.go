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
func (pc SubscribeController) Index(c *gin.Context) {
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
func (pc SubscribeController) Create(c *gin.Context) {
	var s service.Service
	userid := c.Params.ByName("userid")
	planid := c.Params.ByName("planid")

	p, err := s.CreateSubscribeModel(userid, planid, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc SubscribeController) Show(c *gin.Context) {
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
func (pc SubscribeController) Update(c *gin.Context) {
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
