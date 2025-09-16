package handler

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "LABS-BMSTU-BACKEND/internal/app/ds"
)

func (h *Handler) GetTempRequestData(ctx *gin.Context) {
	system_id, err := h.Repository.GetDraftPlanetSystemID()
	if system_id == 0 && err == nil {
		logrus.Infof("Заявка не найдена")
	}

	temp_datas, err := h.Repository.GetPlanetsWithSystemData(uint(system_id))
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "temps_request.html", gin.H{
		"temp_datas": temp_datas,
	})
}