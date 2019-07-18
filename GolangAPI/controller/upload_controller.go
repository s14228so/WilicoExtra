package controller

//serverにレスポンスを返すところ

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/service"
	//planとしてserviceパッケージを読み込む
)

// Controller is user controlller
type UploaderController struct{}

// Index action: GET /users
// Create action: POST /users
func (pc UploaderController) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateImageModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}
