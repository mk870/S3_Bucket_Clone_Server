package main

import (
	"CloudStorage.service/controllers"
	"CloudStorage.service/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	db.Connect()
	controllers.GoogleSignUp(router)
	controllers.GoogleLogin(router)
	controllers.CreateUser(router)
	controllers.GetUsers(router)
	controllers.UpdateUser(router)
	controllers.GetUser(router)
	controllers.DeleteUser(router)
	controllers.Login(router)
	controllers.LoginOut(router)
	controllers.GetVerificationToken(router)
	controllers.VerificationTokenValidation(router)
	router.Run()
}
