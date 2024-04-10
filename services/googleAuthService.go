package services

import (
	"net/http"

	"CloudStorage.service/models"
	"CloudStorage.service/repositories"
	"CloudStorage.service/tokens"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	type GoogleInfo struct {
		Credential string
		ClientId   string
	}
	var googleData GoogleInfo
	c.BindJSON(&googleData)
	payload := RetrieveUser(googleData.Credential, googleData.ClientId)
	if payload == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to authenticate"})
		return
	}
	isVerified := payload.Claims["email_verified"].(bool)
	if !isVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "your google account is not verified"})
		return
	}
	email := payload.Claims["email"].(string)
	firstName := payload.Claims["given_name"].(string)
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this account does not exist"})
	}
	refreshToken := tokens.GenerateRefreshToken(firstName, email)
	if refreshToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create refresh token",
		})
		return
	}
	user.RefreshToken = refreshToken
	repositories.SaveUserUpdate(user)
	accessToken := tokens.GenerateAccessToken(user.FirstName, user.Email, user.Id)
	if accessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create access token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"response":    "logged in successfully",
	})
}

func GoogleSignUp(c *gin.Context) {
	type GoogleInfo struct {
		Credential string
		ClientId   string
	}
	var googleData GoogleInfo
	c.BindJSON(&googleData)
	payload := RetrieveUser(googleData.Credential, googleData.ClientId)
	if payload == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to authenticate"})
		return
	}
	isVerified := payload.Claims["email_verified"].(bool)
	if !isVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "your google account is not verified"})
		return
	}
	email := payload.Claims["email"].(string)
	firstName := payload.Claims["given_name"].(string)
	refreshToken := tokens.GenerateRefreshToken(firstName, email)
	if refreshToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create refresh token",
		})
		return
	}
	newUser := models.User{
		Email:        email,
		FirstName:    firstName,
		IsActive:     true,
		RefreshToken: refreshToken,
	}
	isUserCreated := repositories.CreateUser(&newUser)
	if !isUserCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "this email already exists"})
		return
	}
	accessToken := tokens.GenerateAccessToken(newUser.FirstName, newUser.Email, newUser.Id)
	if accessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate access token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
