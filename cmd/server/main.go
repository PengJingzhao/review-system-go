package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"review-system-go/internal/app/auth"
	"review-system-go/internal/app/user"
	"review-system-go/internal/db"
)

func main() {
	err := db.InitMysqlDB()
	if err != nil {
		log.Fatalf("failed to init mysql db: %v", err)
	}

	r := gin.Default()

	user.RegisterRoutes(r)
	auth.RegisterRoutes(r)

	if err := r.Run(":10201"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
