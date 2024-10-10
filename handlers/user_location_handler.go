package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasharya2901/bookie-backend-go/models"
	"github.com/yasharya2901/bookie-backend-go/services"
)

type UserLocationHandler struct {
	Service *services.UserLocationService
}

func NewUserLocationService(service *services.UserLocationService) *UserLocationHandler {
	return &UserLocationHandler{Service: service}
}

func (h *UserLocationHandler) CreateUserLocationHandler(c *gin.Context) {
	var userLocation models.UserLocation
	if err := c.ShouldBindJSON(&userLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	if err := h.Service.CreateUserLocation(&userLocation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user location"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User location created successfully"})
}

func (h *UserLocationHandler) GetUserLocationHandler(c *gin.Context) {
	id := c.Param("appwrite_user_id")

	userLocation, err := h.Service.GetUserLocationByAppwriteUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user location"})
		return
	}

	if userLocation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User location not found"})
		return
	}

	c.JSON(http.StatusOK, userLocation)
}
