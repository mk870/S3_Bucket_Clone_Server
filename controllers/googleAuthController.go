package controllers

import (
	"CloudStorage.service/services"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(router *gin.Engine) {
	router.POST("/google-login", services.GoogleLogin)
}

func GoogleSignUp(router *gin.Engine) {
	router.POST("/google-signup", services.GoogleSignUp)
}
