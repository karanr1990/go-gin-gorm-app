package main

import (
	"github.com/karanr1990/go-gin-gorm-app/initializers"
	"github.com/karanr1990/go-gin-gorm-app/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
