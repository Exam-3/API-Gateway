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

// AddRecyclingCenter godoc
// @Summary Adds a new recycling center
// @Description Inserts new recycling center info into the database
// @Tags recycling_center
// @Param new_data body item.AddRecyclingCenterRequest true "New recycling center data"
// @Success 200 {object} item.RecyclingCenterResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while adding recycling center"
// @Router /item-system/recycling-centers [post]
func (h *Handler) AddRecyclingCenter(c *gin.Context) {
	h.Logger.Info("AddRecyclingCenter method is starting")

	var req pb.AddRecyclingCenterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind recycling center data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.AddRecyclingCenter(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to add recycling center", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add recycling center"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// SearchRecyclingCenters godoc
// @Summary Searches for recycling centers
// @Description Retrieves recycling centers based on search criteria
// @Tags recycling_center
// @Param search_criteria body item.SearchRecyclingCentersRequest true "Search criteria"
// @Success 200 {object} item.ListRecyclingCentersResponse
// @Failure 400 {object} string "Invalid search criteria"
// @Failure 500 {object} string "Server error while searching recycling centers"
// @Router /item-system/recycling-centers/search [post]
func (h *Handler) SearchRecyclingCenters(c *gin.Context) {
	h.Logger.Info("SearchRecyclingCenters method is starting")

	var req pb.SearchRecyclingCentersRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind search criteria data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.SearchRecyclingCenters(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to search recycling centers", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search recycling centers"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// SubmitItemsForRecycling godoc
// @Summary Submits items for recycling
// @Description Inserts recycling submission info into the database
// @Tags recycling
// @Param new_data body item.SubmitItemsForRecyclingRequest true "New recycling submission data"
// @Success 200 {object} item.RecyclingSubmissionResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while submitting items for recycling"
// @Router /item-system/recycling [post]
func (h *Handler) SubmitItemsForRecycling(c *gin.Context) {
	h.Logger.Info("SubmitItemsForRecycling method is starting")

	var req pb.SubmitItemsForRecyclingRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind recycling submission data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.SubmitItemsForRecycling(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to submit items for recycling", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit items for recycling"})
		return
	}

	c.JSON(http.StatusOK, res)
}
