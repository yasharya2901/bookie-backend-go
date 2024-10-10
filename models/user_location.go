package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLocation struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	AppwriteUserID string    `gorm:"varchar(255); not null"`
	Latitude       float64   `gorm:"type:float;not null"`
	Longitude      float64   `gorm:"type:float;not null"`
	Location       string    `gorm:"type:geography(POINT, 4326);not null"`
}

func (u *UserLocation) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.Location = fmt.Sprintf("POINT(%f %f)", u.Longitude, u.Latitude)
	return nil
}
