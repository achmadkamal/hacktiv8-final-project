package controller

import (
	"hacktiv8_final_project/entity"
	"hacktiv8_final_project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaControllerImpl struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaController(repo repository.SocialMediaRepository) *SocialMediaControllerImpl {
	return &SocialMediaControllerImpl{repo}
}

func (ctrl *SocialMediaControllerImpl) GetAll(c *gin.Context) {
	socialMedias, err := ctrl.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, socialMedias)
}

func (ctrl *SocialMediaControllerImpl) GetOne(c *gin.Context) {
	id := c.Param("id")
	var socialMedia entity.SocialMedia
	if err := ctrl.repo.FindByID(id, &socialMedia); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}
	c.JSON(http.StatusOK, socialMedia)
}

func (ctrl *SocialMediaControllerImpl) CreateSocialMedia(c *gin.Context) {
	var socialMedia entity.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, _ := c.Get("user_id")
	userIDFloat, _ := userIDInterface.(float64)
	userID := uint(userIDFloat)

	socialMedia.UserId = userID

	createdSocialMedia, err := ctrl.repo.Create(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSocialMedia)
}

func (ctrl *SocialMediaControllerImpl) UpdateSocialMedia(c *gin.Context) {
	id := c.Param("id")
	var socialMedia entity.SocialMedia
	if err := ctrl.repo.FindByID(id, &socialMedia); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if socialMedia.UserId != socialMedia.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := ctrl.repo.Update(socialMedia); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (ctrl *SocialMediaControllerImpl) DeleteSocialMedia(c *gin.Context) {
	id := c.Param("id")
	var socialMedia entity.SocialMedia
	if err := ctrl.repo.FindByID(id, &socialMedia); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
