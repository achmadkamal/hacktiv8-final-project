package controller

import (
	"hacktiv8_final_project/entity"
	"hacktiv8_final_project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentControllerImpl struct {
	repo repository.CommentRepository
}

func NewCommentController(repo repository.CommentRepository) *CommentControllerImpl {
	return &CommentControllerImpl{repo}
}

func (ctrl *CommentControllerImpl) GetAll(c *gin.Context) {
	comments, err := ctrl.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (ctrl *CommentControllerImpl) GetOne(c *gin.Context) {
	id := c.Param("id")
	var comment entity.Comment
	if err := ctrl.repo.FindByID(id, &comment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (ctrl *CommentControllerImpl) CreateComment(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, _ := c.Get("user_id")
	userIDFloat, _ := userIDInterface.(float64)
	userID := uint(userIDFloat)

	comment.UserId = userID

	createdComment, err := ctrl.repo.Create(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdComment)
}

func (ctrl *CommentControllerImpl) UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment entity.Comment
	if err := ctrl.repo.FindByID(id, &comment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if comment.UserId != comment.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := ctrl.repo.Update(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (ctrl *CommentControllerImpl) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment entity.Comment
	if err := ctrl.repo.FindByID(id, &comment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
