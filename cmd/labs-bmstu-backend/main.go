package main

import (
	"log"

	"LABS-BMSTU-BACKEND/internal/api"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
}