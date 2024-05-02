package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"cat-social/utils"
	"net/http"
	"strings"
)

type LoginRequest struct {
	Email string `json:"email" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
		return
	}

	user, err := h.repository.FindByEmail(strings.ToLower(req.Email))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if !CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate access token: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged successfully",
		"data": LoginResponse{
			Email:    user.Email,
			Name:        user.Name,
			AccessToken: accessToken,
		},
	})
}
