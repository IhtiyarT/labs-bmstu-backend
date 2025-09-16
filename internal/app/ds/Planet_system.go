package ds

import (
	"time"
)

type Planet_system struct {
	PlanetSystemID uint      `json:"id" gorm:"primaryKey"`
	DateCreated    time.Time `json:"date_created" gorm:"not null"`
	Status         string    `json:"status" gorm:"type:varchar(50);not null"`
	User           Users     `json:"UserID" gorm:"foreignKey:UserID;not null"`

	UserID  uint `json:"-"`
	ModerID uint `json:"-"`
	StarName          string `json:"star_name" gorm:"type:varchar(255)"`
	StarTemperature   uint   `json:"star_temperature"`
}
