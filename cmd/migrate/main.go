package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"LABS-BMSTU-BACKEND/internal/app/ds"
	"LABS-BMSTU-BACKEND/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&ds.Planets{},
		&ds.Planet_system{},
		&ds.Temperature_request{},
		&ds.Users{},
	)
	if err != nil {
		panic("cant migrate db" + err.Error())
	}
}