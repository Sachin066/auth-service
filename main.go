package main

import (
	"github.com/gin-gonic/gin"
	"auth-service/handlers"
)

func main() {
	r := gin.Default()
r.POST("/users", handlers.CreateUser)
r.GET("/users", handlers.GetUsers)
r.PUT("/users/:id", handlers.UpdateUser)
r.DELETE("/users/:id", handlers.DeleteUser)
r.Run()

}
