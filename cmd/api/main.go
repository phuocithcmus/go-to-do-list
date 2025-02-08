package main

import (
	"log"

	"to_do_project/internal/application/services"
	postgres2 "to_do_project/internal/infras/db/postgres"
	"to_do_project/internal/interface/rest"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=to_do_list password=to_do_list dbname=to_do_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	port := ":8080"

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//gormDB.AutoMigrate()

	userRepo := postgres2.NewGormUserRepository(gormDB)

	userService := services.NewUserService(userRepo)

	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong"})
	})

	rest.NewUserController(route, userService)

	if err := route.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
