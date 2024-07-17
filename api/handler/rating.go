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

// AddRating godoc
// @Summary Adds a new rating
// @Description Inserts new rating info into ratings table in PostgreSQL
// @Tags rating
// @Param new_data body item.AddRatingRequest true "New rating data"
// @Success 200 {object} item.Rating
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while adding rating"
// @Router /item-system/ratings/add [post]
func (h *Handler) AddRating(c *gin.Context) {
	h.Logger.Info("AddRating method is starting")

	var req pb.AddRatingRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind rating data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	rating, err := h.ItemClient.AddRating(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to add rating", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add rating"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

// GetRatings godoc
// @Summary Gets all ratings
// @Description Retrieves all ratings info from ratings table in PostgreSQL
// @Tags rating
// @Param new_data body item.GetRatingsRequest true "rating data"
// @Success 200 {object} item.GetRatingsResponse
// @Failure 500 {object} string "Server error while getting ratings"
// @Router /item-system/ratings/GetAll [post]
func (h *Handler) GetRatings(c *gin.Context) {
	h.Logger.Info("GetRatings method is starting")

	var req pb.GetRatingsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind rating data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ratings, err := h.ItemClient.GetRatings(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to get ratings", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

