package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// UserHandler struct to hold handler methods
type UserHandler struct {
	// Add any dependencies if needed
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

// LoginUser handles user login
func (h *UserHandler) LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

// GetProfile handles getting user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User profile"})
}

// UpdateProfile handles updating user profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

// GetHandler returns the handler function based on function name
func (h *UserHandler) GetHandler(functionName string) gin.HandlerFunc {
	method := reflect.ValueOf(h).MethodByName(functionName)
	if !method.IsValid() {
		return func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Handler not found"})
		}
	}
	return func(c *gin.Context) {
		method.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
