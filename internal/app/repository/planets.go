package repository

import (
	"fmt"

	"LABS-BMSTU-BACKEND/internal/app/ds"
)

func (r *Repository) GetPlanets() ([]ds.Planets, error) {
	var planets []ds.Planets
	err := r.db.Find(&planets).Error
	if err != nil {
		return nil, err
	}
	if len(planets) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return planets, nil
}

func (r *Repository) GetPlanet(id int) (ds.Planets, error) {
	planet := ds.Planets{}
	err := r.db.Where("planet_id = ?", id).First(&planet).Error
	if err != nil {
		return ds.Planets{}, err
	}
	return planet, nil
}

func (r *Repository) GetPlanetsByTitle(title string) ([]ds.Planets, error) {
	var planets []ds.Planets
	err := r.db.Where("planet_title ILIKE ?", "%"+title+"%").Find(&planets).Error
	if err != nil {
		return nil, err
	}
	return planets, nil
}

func (r *Repository) GetCountBySystemID(system_id uint) (int64, error) {
    var count int64
    err := r.db.
        Model(&ds.Temperature_request{}).
        Where("planet_system_id = ?", system_id).
        Count(&count).Error
    
    return count, err
}
