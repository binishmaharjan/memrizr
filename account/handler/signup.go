package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// signupReq is not exported, hence the lowercase name
// it is used for validation and json marshalling
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password "binding:"required,gte=6,lte=30"`
}

// Signup Handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}
