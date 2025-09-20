package handler

import (
	"net/http"
	"strconv"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) AddPlanetToSystem(ctx *gin.Context) {
	planet_id, err1 := strconv.Atoi(ctx.Param("planet_id"))
	system_id, err2 := h.Repository.GetDraftPlanetSystemID()
	 if err1 != nil {
        logrus.Error(err1)
        ctx.Redirect(http.StatusFound, "/")
        return
    }
	if err2 != nil {
        logrus.Error(err2)
        ctx.Redirect(http.StatusFound, "/")
        return
    }
	
	if system_id == 0 {
		system_id, err2= h.Repository.CreateNewDraftPlanetSystem(uint(1))
		if err2 != nil {
			logrus.Error(err2)
		}
	}
	
	h.Repository.AddPlanetToSystem(uint(planet_id), uint(system_id))
	ctx.Redirect(http.StatusFound, "/")
}

func (h *Handler) DeletePlanetSystem(ctx *gin.Context) {
	system_id, err := h.Repository.GetDraftPlanetSystemID()
	if system_id == 0 && err == nil{
		logrus.Infof("Система для удаления не найдена")
	}
	h.Repository.DeletePlanetSystem(uint(system_id))
    ctx.Redirect(http.StatusFound, "/")
}

func (h *Handler) GetPlanetSystemByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr) 
	if err != nil {
		logrus.Error(err)
	}

	planet_system, err := h.Repository.GetPlanetSystemByID(uint(id))
	if err != nil {
		logrus.Error(err)
	}
	
	ctx.HTML(http.StatusOK, "planet.html", gin.H{
		"planet_system": planet_system,
	})
}
