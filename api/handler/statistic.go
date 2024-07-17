package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	pb "api-gateway/genproto/item"
)

// Statistics godoc
// @Summary Gets statistics
// @Description Retrieves various statistics from the service
// @Tags statistics
// @Param filter body item.GetStatisticsRequest true "Statistics filter"
// @Success 200 {object} item.GetStatisticsResponse
// @Failure 400 {object} string "Invalid filter"
// @Failure 500 {object} string "Server error while getting statistics"
// @Router /item-system/statistics [post]
func (h *Handler) Statistics(c *gin.Context) {
	h.Logger.Info("Statistics method is starting")

	var req pb.GetStatisticsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind statistics filter data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.Statistics(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to get statistics", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get statistics"})
		return
	}

	c.JSON(http.StatusOK, res)
}
