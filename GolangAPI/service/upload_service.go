package service

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/olahol/go-imageupload"
	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/entity"
)

// Service procides user's behavior

// User is alias of entity.User struct
type Image entity.Image

// CreateModel is create User model
func (s Service) CreateImageModel(c *gin.Context) (Image, error) {
	db := db.GetDB()
	var u Image
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		panic(err)
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		panic(err)
	}

	thumb.Write(c.Writer)
	fmt.Printf("%T %v", thumb, thumb)

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
