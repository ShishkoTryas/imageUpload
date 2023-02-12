package http

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"imageUpload/internal/domain"
)

type UserService interface {
	SignUp(ctx context.Context, inp domain.SignUpUser) error
	SignIn(ctx context.Context, inp domain.SignInUser) (string, error)
	ParseToken(ctx context.Context, token string) (int64, error)
}

type handler struct {
	userService UserService
	client      *s3.Client
}

func NewHandler(user UserService, client *s3.Client) *handler {
	return &handler{
		userService: user,
		client:      client,
	}
}

func (h *handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 10 << 20

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}

	file := router.Group("/")
	{
		file.GET("/upload", h.Start)
		file.POST("/upload", h.Upload)
		file.GET("/files", h.Fiels)
	}

	return router
}
