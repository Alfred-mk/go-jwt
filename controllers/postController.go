package controllers

import (
	"net/http"

	"github.com/Alfred-mk/go-jwt/initializers"
	"github.com/Alfred-mk/go-jwt/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off request body
	var body struct {
		Title string
		Body  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	// create the Post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a post",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"success": "Your blog post has been created",
		"post":    post,
	})
}

func PostsIndex(c *gin.Context) {
	// get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with the array
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off of the URL
	id := c.Param("id")

	// get single post
	var post models.Post
	initializers.DB.First(&post, id)

	// respond with the array
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id off of the URL
	id := c.Param("id")

	// get data off request body
	var body struct {
		Title string
		Body  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	// get single post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the Post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with the array
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get id off of the URL
	id := c.Param("id")

	// Delete the Post
	initializers.DB.Delete(&models.Post{}, id)

	// respond with the array
	c.JSON(http.StatusOK, gin.H{
		"message": "The post has been deleted",
	})
}
