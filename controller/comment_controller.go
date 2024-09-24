package controller

import "github.com/gin-gonic/gin"

type CommentController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}
