package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/s14228so/WilicoExtra/GolangAPI/service"
)

type CoachController struct{}

func (pc CoachController) Index(c *gin.Context) {
	var s service.Service
	p, err := s.GetCoachAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}

}

// func (pc CoachController) Plans(c *gin.Context) {
// 	var s service.Service
// 	p, err := s.GetCoachPlans()

// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}

// }

func (pc CoachController) Create(c *gin.Context) {
	var s service.Service

	p, err := s.CreateCoachModel(c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, p)
	}
}

func (pc CoachController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetCoachByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc CoachController) Update(c *gin.Context) {

	id := c.Params.ByName("id")
	var s service.Service

	p, err := s.UpdateCoachByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc CoachController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteCoachByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}

}
