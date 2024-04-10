package controllers

import (
	"CloudStorage.service/middleware"
	"CloudStorage.service/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.Engine) {
	router.POST("/user", services.CreateUser)
}

func GetUsers(router *gin.Engine) {
	router.GET("/users", middleware.AuthValidator, services.GetUsers)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user/:id", middleware.AuthValidator, services.UpdateUser)
}

func GetUser(router *gin.Engine) {
	router.GET("/user/:id", middleware.AuthValidator, services.GetUser)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", middleware.AuthValidator, services.DeleteUser)
}

func CreateObject(router *gin.Engine) {
	router.POST("/create-object", services.DeleteUser)
}
