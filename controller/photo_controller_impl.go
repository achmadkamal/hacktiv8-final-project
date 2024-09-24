package controller

import (
	"hacktiv8_final_project/entity"
	"hacktiv8_final_project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoControllerImpl struct {
	repo repository.PhotoRepository
}

func NewPhotoController(repo repository.PhotoRepository) *PhotoControllerImpl {
	return &PhotoControllerImpl{repo}
}

func (ctrl *PhotoControllerImpl) GetAll(c *gin.Context) {
	photos, err := ctrl.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, photos)
}

func (ctrl *PhotoControllerImpl) GetOne(c *gin.Context) {
	id := c.Param("id")
	var photo entity.Photo
	if err := ctrl.repo.FindByID(id, &photo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}
	c.JSON(http.StatusOK, photo)
}

func (ctrl *PhotoControllerImpl) CreatePhoto(c *gin.Context) {
	var photo entity.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, _ := c.Get("user_id")
	userIDFloat, _ := userIDInterface.(float64)
	photo.UserId = uint(userIDFloat)

	createdPhoto, err := ctrl.repo.Create(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPhoto)
}

func (ctrl *PhotoControllerImpl) UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	var photo entity.Photo
	if err := ctrl.repo.FindByID(id, &photo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, _ := c.Get("user_id")
	userIDFloat, _ := userIDInterface.(float64)
	userID := uint(userIDFloat)

	if userID != photo.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := ctrl.repo.Update(photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (ctrl *PhotoControllerImpl) DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	var photo entity.Photo
	if err := ctrl.repo.FindByID(id, &photo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	userIDInterface, _ := c.Get("user_id")
	userIDFloat, _ := userIDInterface.(float64)
	userID := uint(userIDFloat)

	if userID != photo.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
