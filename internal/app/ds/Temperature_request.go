package ds

type Temperature_request struct {
	PlanetID          uint `json:"planet_id" gorm:"primaryKey;auto_increment:false"`
	PlanetSystemID    uint `json:"planet_system_id" gorm:"primaryKey;auto_increment:false"`
	PlanetTemperature uint `json:"planet_temperature"`
	PlanetDistance    uint `json:"planet_distance"`
}
