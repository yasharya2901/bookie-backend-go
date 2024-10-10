package services

import (
	"github.com/yasharya2901/bookie-backend-go/models"
	"gorm.io/gorm"
)

type UserLocationService struct {
	DB *gorm.DB
}

func NewUserLocationService(db *gorm.DB) *UserLocationService {
	return &UserLocationService{DB: db}
}

func (s *UserLocationService) CreateUserLocation(userLocation *models.UserLocation) error {
	return s.DB.Create(userLocation).Error
}

func (s *UserLocationService) GetUserLocationByAppwriteUserID(appwriteUserId string) (*models.UserLocation, error) {
	var userLocation models.UserLocation
	if err := s.DB.Where("appwrite_user_id = ?", appwriteUserId).First(&userLocation).Error; err != nil {
		return nil, err
	}
	return &userLocation, nil
}
