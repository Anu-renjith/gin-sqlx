package main

import (
	"log"
	"os"

	"github.com/Anu-renjith/gin-sqlx/handler"
	"github.com/Anu-renjith/gin-sqlx/repository"
	"github.com/Anu-renjith/gin-sqlx/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/godemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.MustExec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL)")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running at port %s", port)
	r.Run(":" + port)
}
