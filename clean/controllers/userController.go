package controllers

import (
	"github.com/dipeshdulal/fxinit/clean/services"
	"github.com/gin-gonic/gin"
)

// UserController bundling
type UserController struct {
	Get  func(c *gin.Context)
	Post func(c *gin.Context)
}

// NewUserController creates a new controller
func NewUserController(userService services.UserService) UserController {
	return UserController{
		Get:  getUserController(userService),
		Post: postUserController(userService),
	}
}

func getUserController(userService services.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		userService.GetUser()
		c.JSON(200, "Get user called here")
	}
}

func postUserController(userService services.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		userService.CreateUser()
		c.JSON(200, "Create user called.")
	}
}
