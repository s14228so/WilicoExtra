package service

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

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

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return u, err
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return u, err
		}
	}
	//https://github.com/gin-gonic/examples/blob/master/upload-file/multiple/main.go
	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files.", len(files)))

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
