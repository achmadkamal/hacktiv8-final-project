package controller

import "github.com/gin-gonic/gin"

type PhotoController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreatePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
}
