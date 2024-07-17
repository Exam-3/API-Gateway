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

// AddItemCategory godoc
// @Summary Adds an item category
// @Description Inserts new item category info into item_categories table in PostgreSQL
// @Tags item
// @Param new_data body item.AddItemCategoryRequest true "New data"
// @Success 200 {object} item.AddItemCategoryResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while adding item category"
// @Router /item-system/category/catogories [post]
func (h *Handler) AddItemCategory(c *gin.Context) {
	h.Logger.Info("AddItemCategory method is starting")

	var req pb.AddItemCategoryRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind item category data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	itemCategory, err := h.ItemClient.AddItemCategory(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to add item category", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item category"})
		return
	}

	c.JSON(http.StatusOK, itemCategory)
}
