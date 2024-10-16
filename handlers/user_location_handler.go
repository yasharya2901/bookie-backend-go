package handlers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func (h *UserLocationHandler) CreateUserFromAppwrite(c *gin.Context) {
	var appwriteResponse map[string]any

	// Get the raw data from the request
	payloadBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" { // The "X-Forwarded-Proto" header is set by Ngrok while forwarding the request to the local server
		scheme = "https"
	}

	fullURL := scheme + "://" + c.Request.Host + c.Request.URL.Path

	signature, err := generateSignature(fullURL, string(payloadBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify webhook signature"})
		return
	}

	headerSignature := c.GetHeader("X-Appwrite-Webhook-Signature")

	if signature != headerSignature {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed Authentication Check"})
		return
	}

	// Unmarshalling the payload body because c.GetRawData() empties the request body
	if err := json.Unmarshal(payloadBody, &appwriteResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the required fields are present
	if _, ok := appwriteResponse["latitude"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude is required"})
		return
	}
	if _, ok := appwriteResponse["longitude"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Longitude is required"})
		return
	}
	if _, ok := appwriteResponse["$id"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "$id is required"})
		return
	}

	// Create the user location
	userLocation := models.UserLocation{
		Latitude:       appwriteResponse["latitude"].(float64),
		Longitude:      appwriteResponse["longitude"].(float64),
		AppwriteUserID: appwriteResponse["$id"].(string),
	}

	// Database call to create the user location
	err = h.Service.CreateUserLocation(&userLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user location"})
		return
	}
	message := fmt.Sprintf("User %v location created successfully", userLocation.AppwriteUserID)
	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func generateSignature(url, payloadBody string) (string, error) {
	data := url + payloadBody

	hm := hmac.New(sha1.New, []byte(os.Getenv("APPWRITE_WEBHOOK_SECRET")))
	hm.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hm.Sum(nil)), nil
}
