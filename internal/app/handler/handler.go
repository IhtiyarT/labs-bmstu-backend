package handler

import (
	"LABS-BMSTU-BACKEND/internal/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) GetPlanets(ctx *gin.Context) {
	var planets []repository.Planet
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

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"planets": planets,
		"query": searchQuery,
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

func (h *Handler) GetTempRequestData(ctx *gin.Context) {
	temp_datas, err := h.Repository.GetTempRequestDataById(1)
	if err != nil {
		logrus.Error(err)
	}

	ids := make([]int, len(temp_datas))
	for i, data := range temp_datas {
		ids[i] = data.Planet_id
	}

	planet_datas, err := h.Repository.GetPlanetsById(ids)
	if err != nil {
		logrus.Error(err)
	}

	new_temp_datas := make([]repository.Temps_data, len(temp_datas))
    for i := range temp_datas {
        new_temp_datas[i] = repository.Temps_data{
            Planet_image:       planet_datas[i].Planet_image,
            Planet_title:       planet_datas[i].Planet_title,
            Albedo:      planet_datas[i].Albedo,
            Planet_distance:    temp_datas[i].Planet_distance,
            Planet_temperature: temp_datas[i].Planet_temperature,
        }
    }

	ctx.HTML(http.StatusOK, "temps_request.html", gin.H{
		"temp_datas": new_temp_datas,
	})
}
