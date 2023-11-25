package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karanr1990/go-gin-gorm-app/initializers"
	"github.com/karanr1990/go-gin-gorm-app/models"
	"net/http"
)

func PostsCreate(ctx *gin.Context) {

	//create a post schema to validate input data
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	ctx.Bind(&body)
	post := models.Post{Tittle: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error Creating Post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}
