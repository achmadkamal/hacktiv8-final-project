package controller

import (
	"hacktiv8_final_project/entity"
	"hacktiv8_final_project/middleware"
	"hacktiv8_final_project/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserControllerImpl {
	return &UserControllerImpl{repo}
}

func (ctrl *UserControllerImpl) Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUsername, err := ctrl.repo.FindByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if storedUsername.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return
	}

	storedEmail, err := ctrl.repo.FindByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if storedEmail.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	createdUser, err := ctrl.repo.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (ctrl *UserControllerImpl) Login(c *gin.Context) {

	var user struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser, err := ctrl.repo.FindByUsername(user.Username)
	if err != nil || storedUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(storedUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
