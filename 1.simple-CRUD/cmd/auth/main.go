package main

import (
	"1.simple-CRUD/internal"
	"1.simple-CRUD/internal/auth/handler"
	"1.simple-CRUD/internal/auth/model"
	"1.simple-CRUD/internal/auth/repository"
	"1.simple-CRUD/internal/auth/router"
	"1.simple-CRUD/internal/auth/service"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := internal.InitDB()
	db.AutoMigrate(&model.User{})
	userRepo := repository.NewUserRepository(db)
	userRepo.Create(&model.User{Username: "arash", Password: "1234"})
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	r := router.SetupRouter(authHandler)
	port := os.Getenv("APP_PORT")
	r.Run(":" + port)
}
