package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karanr1990/go-gin-gorm-app/controllers"
	"github.com/karanr1990/go-gin-gorm-app/initializers"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)

	router.Run()

}
