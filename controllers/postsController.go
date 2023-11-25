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

func PostsFindAll(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(http.StatusOK, gin.H{"message": "All posts", "posts": posts})
}

func PostFindOne(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")
	initializers.DB.First(&post, id)
	ctx.JSON(http.StatusOK, gin.H{"message": "post found", "post": post})
}

func PostUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	ctx.Bind(&input)

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return

	}

	if err := initializers.DB.Model(&post).Updates(&models.Post{Tittle: input.Title, Body: input.Body}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post updated", "post": post})
}

func PostDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if (initializers.DB.First(&models.Post{}, id).RowsAffected == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post not found"})
		return
	}
	initializers.DB.Delete(&models.Post{}, id)

	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})

}
