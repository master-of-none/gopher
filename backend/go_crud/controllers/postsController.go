package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/go_crud/initializers"
	"github.com/master-of-none/go_crud/models"
)

func PostsCreate(ctx *gin.Context) {
	// Get the data from the request body
	var body struct {
		Body  string
		Title string
	}
	ctx.Bind(&body)
	// Create a Post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}
	// Return It
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	// Get the data
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	// Ge ID URL
	id := ctx.Param("id")
	var post models.Post

	initializers.DB.First(&post, id)
	ctx.JSON(200, gin.H{
		"post": post,
	})
}
