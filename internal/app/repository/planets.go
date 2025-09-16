package repository

import (
	"fmt"

	"LABS-BMSTU-BACKEND/internal/app/ds"
)

type PlanetWithSystemData struct {
	PlanetID          uint    `json:"planet_id"`
	PlanetTitle       string  `json:"planet_title"`
	PlanetImage       string  `json:"planet_image"`
	PlanetDescription string  `json:"planet_description"`
	Albedo            float64 `json:"albedo"`
	PlanetSystemID    uint    `json:"planet_system_id"`
	PlanetTemperature uint    `json:"planet_temperature"`
	PlanetDistance    uint    `json:"planet_distance"`
	StarName          string  `json:"star_name"`
	StarTemperature   uint    `json:"star_temperature"`
}

func (r *Repository) GetOrders() ([]ds.Planets, error) {
	var orders []ds.Planets
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return orders, nil
}

func (r *Repository) GetOrder(id int) (ds.Planets, error) {
	order := ds.Planets{}
	err := r.db.Where("planet_id = ?", id).First(&order).Error
	if err != nil {
		return ds.Planets{}, err
	}
	return order, nil
}

func (r *Repository) GetOrdersByTitle(title string) ([]ds.Planets, error) {
	var orders []ds.Planets
	err := r.db.Where("planet_title ILIKE ?", "%"+title+"%").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *Repository) GetPlanetsWithSystemData(systemID uint) ([]PlanetWithSystemData, error) {
	var results []PlanetWithSystemData
	
	err := r.db.
		Table("planets p").
		Select("p.*, tr.planet_system_id, tr.planet_temperature, tr.planet_distance, ps.star_name, ps.star_temperature").
		Joins("INNER JOIN temperature_requests tr ON p.planet_id = tr.planet_id").
		Joins("INNER JOIN planet_systems ps ON tr.planet_system_id = ps.planet_system_id").
		Where("p.is_delete = ? AND tr.planet_system_id = ?", false, systemID).
		Scan(&results).Error
	
	if err != nil {
		return nil, err
	}
	
	return results, nil
}
