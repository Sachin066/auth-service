package handlers

import (
	"auth-service/db"
	"auth-service/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&user)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := db.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.ShouldBindJSON(&user)
	db.DB.Save(&user)
	c.JSON(200, user)
}
func DeleteUser(c *gin.Context) {
	db.DB.Delete(&models.User{}, c.Param("id"))
	c.JSON(200, gin.H{"message": "User deleted"})
}
