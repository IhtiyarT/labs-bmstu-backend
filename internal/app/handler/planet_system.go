package handler

import (
	"net/http"
	"strconv"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "LABS-BMSTU-BACKEND/internal/app/ds"
)

func (h *Handler) AddPlanetToSystem(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Error(err)
	}
	
	h.Repository.AddPlanetToSystem(uint(id), uint(1))
	ctx.Redirect(http.StatusFound, "/")
}