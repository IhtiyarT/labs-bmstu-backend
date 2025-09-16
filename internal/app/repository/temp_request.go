package repository

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
