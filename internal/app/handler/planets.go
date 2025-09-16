package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"LABS-BMSTU-BACKEND/internal/app/ds"
)

func (h *Handler) GetPlanets(ctx *gin.Context) {
	var planets []ds.Planets
	var err error

	searchQuery := ctx.Query("query") 
	if searchQuery == "" {
		planets, err = h.Repository.GetPlanets()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		planets, err = h.Repository.GetPlanetsByTitle(searchQuery) 
		if err != nil {
			logrus.Error(err)
		}
	}

	system_id, err1 := h.Repository.GetDraftPlanetSystemID()
	if err1 != nil {
		logrus.Error(err1)
	}
	
	cartCount, err := h.Repository.GetCountBySystemID(system_id)
    if err != nil {
        cartCount = 0
        logrus.Error("Error getting cart count:", err)
    }

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"time":   time.Now().Format("15:04:05"),
		"planets": planets,
		"query":  searchQuery,
		"cartCount": cartCount,
		"system_id": system_id,
	})
}

func (h *Handler) GetPlanet(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr) 
	if err != nil {
		logrus.Error(err)
	}

	planet, err := h.Repository.GetPlanet(id)
	if err != nil {
		logrus.Error(err)
	}
	
	ctx.HTML(http.StatusOK, "planet.html", gin.H{
		"planet": planet,
	})
}
