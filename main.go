package main

import (
	"hacktiv8_final_project/database"
	"hacktiv8_final_project/repository"
	"hacktiv8_final_project/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.NewDB()
	userRepo := repository.NewUserRepository(db)
	photoRepo := repository.NewPhotoRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	socialMediaRepo := repository.NewSocialMediaRepository(db)

	router := gin.Default()
	routes.SetupRoutes(router, userRepo, photoRepo, commentRepo, socialMediaRepo)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
