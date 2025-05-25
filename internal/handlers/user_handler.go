package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bethojunior/go-blog/internal/models"
	"github.com/bethojunior/go-blog/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) ListUser(c *gin.Context) {
	users, err := h.service.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		return
	}

	userID := uint(0)
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
		return
	}

	user, err := h.service.GetUser(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
