package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/s14228so/WilicoExtra/GolangAPI/service"
)

type CardController struct{}

func (pc CardController) GET(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	p, err := s.GetCardByID(id)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, p)
	}
}

func (pc CardController) Create(c *gin.Context) {
	id := c.Params.ByName("userID")
	var s service.Service

	p, err := s.CreateCardModel(id, c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, p)
	}
}

func (pc CardController) Update(c *gin.Context) {

	id := c.Params.ByName("userID")
	var s service.Service

	p, err := s.UpdateCardByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc CardController) Delete(c *gin.Context) {
	id := c.Params.ByName("userID")
	var s service.Service

	if err := s.DeleteCardByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}

}
