package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/s14228so/WilicoExtra/GolangAPI/bookshelf"
)

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}

// CreateImageModel comment
func (s Service) CreateImageModel(c *gin.Context) (url string, err error) {
	fmt.Printf("c: %v", c)
	// f, err := r.MultipartForm()
	form, err := c.MultipartForm()
	post := c.PostForm("key")
	files := form.File[post]
	// fh, err := c.FormFile("image")

	if err == http.ErrMissingFile {
		fmt.Println("errorだよ")
		return "", nil
	}
	if err != nil {
		fmt.Println("なんかerrorだよ")
		return "", err
	}

	if bookshelf.StorageBucket == nil {
		fmt.Println("storageないよ！")
		return "", errors.New("storage bucket is missing - check config.go")
	}
	fmt.Printf("bookshelf: %v", bookshelf.StorageBucketName)

	var publicURL string
	var name string
	// random filename, retaining existing extension.

	for _, file := range files {
		name = uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)

		fmt.Printf("name: %v", name)
		// go routine などを呼び出す元の方でオブジェクトを生成する。context.Background() は空のコンテキストを生成する。
		ctx := context.Background()
		w := bookshelf.StorageBucket.Object(name).NewWriter(ctx)
		fmt.Printf("w: %v", w)
		// Warning: storage.AllUsers gives public read access to anyone.
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		w.ContentType = file.Header.Get("Content-Type")

		// Entries are immutable, be aggressive about caching (1 day).
		w.CacheControl = "public, max-age=86400"

		// if _, err := io.Copy(w, f); err != nil {
		// 	return "", err
		// }
		// if err := w.Close(); err != nil {
		// 	return "", err
		// }
		publicURL = "https://storage.googleapis.com/%s/%s"

		c.String(http.StatusOK, fmt.Sprintf(name))

	}

	return fmt.Sprintf(publicURL, bookshelf.StorageBucketName, name), nil

}

// func (s Service) CreateImageodel(c *gin.Context) (Image, error) {
// 	var u Image

// 	form, err := c.MultipartForm()
// 	post := c.PostForm("key")
// 	if err != nil {
// 		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
// 		return u, err
// 	}

// 	files := form.File[post]
// 	log.Printf("file: %v", files) //ここが空

// 	for _, file := range files {
// 		dstDir := "./public/unko.jpg"
// 		if err := c.SaveUploadedFile(file, dstDir); err != nil {
// 			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
// 			return u, err
// 		}
// 	}

// 	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files.", len(files)))

// 	// if err := c.BindJSON(&u); err != nil {
// 	// 	return u, err
// 	// }

// 	// if err := db.Create(&u).Error; err != nil {
// 	// 	return u, err
// 	// }

// 	return u, nil
// }
