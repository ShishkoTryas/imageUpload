package http

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"path/filepath"
)

func (h *handler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")

	if err != nil {
		fmt.Println("Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext != ".png" && ext != ".jpeg" {
		fmt.Println("Error: file format not supported")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File format not supported"})
		return
	}

	uploader := manager.NewUploader(h.client)

	result, uploderErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("my-aws-bucket-uploadimage"),
		Key:         aws.String(header.Filename),
		Body:        file,
		ACL:         "public-read",
		ContentType: aws.String(mime.TypeByExtension(ext)),
	}, func(u *manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10 MiB
	})

	if uploderErr != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to upload image",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"iamge": result.Location,
	})
}

func (h *handler) Fiels(c *gin.Context) {
	result, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("my-aws-bucket-uploadimage"),
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, object := range result.Contents {
		c.JSON(http.StatusOK, gin.H{
			"Size:":     object.Size,
			"Modified:": object.LastModified,
			"Key:":      *object.Key,
		})
	}
}

func (h *handler) Start(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
