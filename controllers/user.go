package controller

import (
	"cat-social/models/dto/request"
	service "cat-social/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{service}
}

func (uC *userController) SignUp(c *gin.Context) {
	var signUpRequest request.SignupRequest

	err := c.ShouldBindJSON(&signUpRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	user, err := uC.userService.Create(signUpRequest)

	if err != nil {
		var error string = err.Error()
		if error == "EMAIL ALREADY EXIST" {
			c.JSON(http.StatusConflict, gin.H{
				"errors": "Confilct: email already exist",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

func (uC *userController) SignIn(c *gin.Context) {
	var loginRequest request.SignInRequest

	err := c.ShouldBindJSON(&loginRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": "errorMessages",
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	user, err := uC.userService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged successfully",
		"data":    user,
	})
}
