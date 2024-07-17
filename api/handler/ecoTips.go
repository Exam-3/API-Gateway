package handler

import (
	pb "api-gateway/genproto/item"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// CreateEcoTip godoc
// @Summary Creates a new eco tip
// @Description Inserts new eco tip info into eco_tips table in PostgreSQL
// @Tags eco_tip
// @Param new_data body item.CreateEcoTipRequest true "New data"
// @Success 200 {object} item.CreateEcoTipResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while creating eco tip"
// @Router /item-system/eco-tips [post]
func (h *Handler) CreateEcoTip(c *gin.Context) {
	h.Logger.Info("CreateEcoTip method is starting")

	var req pb.CreateEcoTipRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind eco tip data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ecoTip, err := h.ItemClient.CreateEcoTip(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to create eco tip", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create eco tip"})
		return
	}

	c.JSON(http.StatusOK, ecoTip)
}

// GetEcoTips godoc
// @Summary Gets all eco tips
// @Description Retrieves all eco tips info from PostgreSQL
// @Tags eco_tip
// @Param new_data body item.GetEcoTipsRequest true "Request data"
// @Success 200 {object} item.GetEcoTipsResponse
// @Failure 500 {object} string "Server error while getting eco tips"
// @Router /item-system/eco-tips [get]
func (h *Handler) GetEcoTips(c *gin.Context) {
	h.Logger.Info("GetEcoTips method is starting")

	var req pb.GetEcoTipsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind eco tips request data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ecoTips, err := h.ItemClient.GetEcoTips(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to get eco tips", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get eco tips"})
		return
	}

	c.JSON(http.StatusOK, ecoTips)
}
