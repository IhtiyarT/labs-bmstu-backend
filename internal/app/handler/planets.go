package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"LABS-BMSTU-BACKEND/internal/app/ds"
)

func (h *Handler) GetOrders(ctx *gin.Context) {
	var orders []ds.Planets
	var err error

	searchQuery := ctx.Query("query") 
	if searchQuery == "" {
		orders, err = h.Repository.GetOrders()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		orders, err = h.Repository.GetOrdersByTitle(searchQuery) 
		if err != nil {
			logrus.Error(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"time":   time.Now().Format("15:04:05"),
		"orders": orders,
		"query":  searchQuery,
	})
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr) 
	if err != nil {
		logrus.Error(err)
	}

	order, err := h.Repository.GetOrder(id)
	if err != nil {
		logrus.Error(err)
	}
	
	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"order": order,
	})
}

func (h *Handler) GetTempRequestData(ctx *gin.Context) {
	temp_datas, err := h.Repository.GetPlanetsWithSystemData(1)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "temps_request.html", gin.H{
		"temp_datas": temp_datas,
	})
}

func (h *Handler) DeletePlanet(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Repository.DeletePlanet(id)
    ctx.Redirect(http.StatusFound, "/")
}

func (h *Handler) AddPlanetToSystem(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Error(err)
	}
	
	h.Repository.AddPlanetToSystem(uint(id), uint(1))
	ctx.Redirect(http.StatusFound, "/")
}
