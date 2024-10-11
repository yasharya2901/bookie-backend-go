package models

type AppwriteRespone struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ApprwriteUserID string  `json:"$id"`
}
