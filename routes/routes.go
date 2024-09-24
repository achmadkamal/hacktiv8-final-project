package routes

import (
	"hacktiv8_final_project/controller"
	"hacktiv8_final_project/middleware"
	"hacktiv8_final_project/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userRepo repository.UserRepository,
	photoRepo repository.PhotoRepository,
	commentRepo repository.CommentRepository,
	socialMediaRepo repository.SocialMediaRepository,
) {
	userController := controller.NewUserController(userRepo)
	photoController := controller.NewPhotoController(photoRepo)
	commentController := controller.NewCommentController(commentRepo)
	socialMediaController := controller.NewSocialMediaController(socialMediaRepo)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	protected := router.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/photos", photoController.GetAll)
		protected.GET("/photos/:id", photoController.GetOne)
		protected.POST("/photos", photoController.CreatePhoto)
		protected.PUT("/photos/:id", photoController.UpdatePhoto)
		protected.DELETE("/photos/:id", photoController.DeletePhoto)

		protected.GET("/comments", commentController.GetAll)
		protected.GET("/comments/:id", commentController.GetOne)
		protected.POST("/comments", commentController.CreateComment)
		protected.PUT("/comments/:id", commentController.UpdateComment)
		protected.DELETE("/comments/:id", commentController.DeleteComment)

		protected.GET("/socialmedia", socialMediaController.GetAll)
		protected.GET("/socialmedia/:id", socialMediaController.GetOne)
		protected.POST("/socialmedia", socialMediaController.CreateSocialMedia)
		protected.PUT("/socialmedia/:id", socialMediaController.UpdateSocialMedia)
		protected.DELETE("/socialmedia/:id", socialMediaController.DeleteSocialMedia)
	}
}
