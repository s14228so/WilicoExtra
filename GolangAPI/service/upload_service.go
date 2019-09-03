package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/golang-samples/getting-started/bookshelf"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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
	StorageBucketName := "wilico-extra"
	StorageBucket, err := configureStorage(StorageBucketName)

	fmt.Printf("bucket: %v", StorageBucket)
	if StorageBucket == nil {
		fmt.Println("storageないよ！")
		return "", errors.New("storage bucket is missing - check config.go")
	}
	fmt.Printf("bookshelf: %v", StorageBucketName)

	var publicURL string
	var name string
	// random filename, retaining existing extension.

	for _, file := range files {
		name = uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)

		fmt.Printf("name: %v", name)
		// go routine などを呼び出す元の方でオブジェクトを生成する。context.Background() は空のコンテキストを生成する。
		ctx := context.Background()
		w := StorageBucket.Object(name).NewWriter(ctx)
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
