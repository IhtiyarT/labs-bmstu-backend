package ds

type Planets struct {
	PlanetID          uint    `json:"planet_id" gorm:"primaryKey"`
	PlanetTitle       string  `json:"planet_title" gorm:"type:varchar(50)"`
	PlanetImage       string  `json:"planet_image" gorm:"type:varchar(100)"`
	PlanetDescription string  `json:"Planet_description" gorm:"type:text"`
	Albedo            float64 `json:"albedo" gorm:"not null"`
	IsDelete          bool    `json:"is_delete"`
}
