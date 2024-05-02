package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"cat-social/entities"
	"cat-social/utils"
	"net/http"
	"strings"
)

type RegisterRequest struct {
	Email string `json:"email" validate:"required,min=5,max=15,email"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type RegisterResponse struct {
	Email    string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	fmt.Println("HELLO!")
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request data: %s", err.Error())})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
		return
	}

	hashedPw, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser := &entities.User{
		Model:    entities.Model{ID: uuid.New()},
		Email: strings.ToLower(req.Email),
		Name:     strings.ToLower(req.Name),
		Password: hashedPw,
	}

	if err := h.repository.Create(newUser); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create user: %s", err.Error())})
		return
	}

	accessToken, err := utils.GenerateToken(newUser.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate access token: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": RegisterResponse{
			Email:    newUser.Email,
			Name:        newUser.Name,
			AccessToken: accessToken,
		},
	})
}
