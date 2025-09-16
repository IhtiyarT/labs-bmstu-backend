package handler

import (
	"net/http"
	// "strconv"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "LABS-BMSTU-BACKEND/internal/app/ds"
)

func (h *Handler) GetTempRequestData(ctx *gin.Context) {
	temp_datas, err := h.Repository.GetPlanetsWithSystemData(1)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "temps_request.html", gin.H{
		"temp_datas": temp_datas,
	})
}