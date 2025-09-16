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
		"orders": orders,
		"query":  searchQuery,
		"cartCount": cartCount,
		"system_id": system_id,
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
