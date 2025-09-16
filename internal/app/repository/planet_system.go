package repository

import (
	// "fmt"

	"LABS-BMSTU-BACKEND/internal/app/ds"
)

func (r *Repository) AddPlanetToSystem(planetID, systemID uint) error {
	tempRequest := &ds.Temperature_request{
		PlanetID:       planetID,
		PlanetSystemID: systemID,
	}
	
	return r.db.Create(tempRequest).Error
}