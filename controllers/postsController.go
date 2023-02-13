package controllers

import (
	"github.com/gin-gonic/gin"
	"nothing.example/initializers"
	"nothing.example/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body string
		Name string
	}
	c.Bind(&body)
	//Create a POST
	post := models.Post{Name: body.Name, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsIndex(c *gin.Context) {
	//Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respnd with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}
func PostsShow(c *gin.Context) {
	id := c.Param("id")
	//Get posts
	var posts models.Post
	initializers.DB.Find(&posts, id)
	//Respnd with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}
func PostsUpdate(c *gin.Context) {
	//get id of url
	id := c.Param("id")
	//get dat of req body
	var body struct {
		Body string
		Name string
	}
	c.Bind(&body)
	//find post were updating
	var posts models.Post
	initializers.DB.Find(&posts, id)
	//update it
	initializers.DB.Model(&posts).Updates(models.Post{
		Name: body.Name,
		Body: body.Body,
	})
	//respond it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsDelete(c *gin.Context) {
	//Get id
	id := c.Param("id")
	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//Respond
	c.Status(200)
}
