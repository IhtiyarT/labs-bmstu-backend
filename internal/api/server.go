package api

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"LABS-BMSTU-BACKEND/internal/app/handler"
	"LABS-BMSTU-BACKEND/internal/app/repository"
)

func StartServer() {
	log.Println("Server start up")

	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Error("Ошибка инициализации репозитория")
	}

	handler := handler.NewHandler(repo)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./resources")

	r.GET("/", handler.GetPlanets)
	r.GET("/planet/:id", handler.GetPlanet)
	r.GET("/temps-request", handler.GetTempRequestData)

	r.Run()
	log.Println("Server down")

}