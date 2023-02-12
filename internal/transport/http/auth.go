package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"imageUpload/internal/domain"
	"log"
	"net/http"
)

func (h *handler) SignUp(c *gin.Context) {
	var inputData domain.SignUpUser
	if err := c.BindJSON(&inputData); err != nil {
		log.Println("Cannot bind Data")
		return
	}
	if err := inputData.Validate(); err != nil {
		log.Println("Incorrect Data")
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	err := h.userService.SignUp(context.TODO(), inputData)
	if err != nil {
		log.Println("Cannot Create User")
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}

	c.JSON(http.StatusOK, inputData)
}

func (h *handler) SignIn(c *gin.Context) {
	var inputData domain.SignInUser
	if err := c.BindJSON(&inputData); err != nil {
		log.Println("Cannot bind Data")
		return
	}
	if err := inputData.Validate(); err != nil {
		log.Println("Incorrect Data")
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	token, err := h.userService.SignIn(context.TODO(), inputData)
	if err != nil {
		log.Println("Cannot auth")
		c.String(http.StatusUnauthorized, fmt.Sprintf("error: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
