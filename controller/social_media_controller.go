package controller

import "github.com/gin-gonic/gin"

type SocialMediaController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreateSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
}
