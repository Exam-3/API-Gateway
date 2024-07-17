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

// AddItem godoc
// @Summary Adds a new item
// @Description Inserts new item info into items table in PostgreSQL
// @Tags item
// @Param new_data body item.AddItemRequest true "New item data"
// @Success 200 {object} item.ItemResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while adding item"
// @Router /item-system/items/addItem [post]
func (h *Handler) AddItem(c *gin.Context) {
	h.Logger.Info("AddItem method is starting")

	var req pb.AddItemRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind item data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	item, err := h.ItemClient.AddItem(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to add item", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateItem godoc
// @Summary Updates an item
// @Description Updates item info in items table in PostgreSQL
// @Tags item
// @Param item_id path string true "Item ID"
// @Param update_data body item.UpdateItemRequest true "Updated item data"
// @Success 200 {object} item.ItemResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while updating item"
// @Router /item-system/items/{item_id} [put]
func (h *Handler) UpdateItem(c *gin.Context) {
	h.Logger.Info("UpdateItem method is starting")

	var req pb.UpdateItemRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind item data", "error", err)
		return
	}

	id := c.Param("item_id")
	
	req.ItemId = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	item, err := h.ItemClient.UpdateItem(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to update item", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteItem godoc
// @Summary Deletes an item
// @Description Deletes item from items table in PostgreSQL
// @Tags item
// @Param item_id path string true "Item ID"
// @Success 200 {object} item.DeleteItemResponse
// @Failure 400 {object} string "Invalid item ID"
// @Failure 500 {object} string "Server error while deleting item"
// @Router /item-system/items/{item_id} [delete]
func (h *Handler) DeleteItem(c *gin.Context) {
	h.Logger.Info("DeleteItem method is starting")

	id := c.Param("item_id")

	req := pb.DeleteItemRequest{ItemId: id}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	item, err := h.ItemClient.DeleteItem(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to delete item", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// ListItems godoc
// @Summary Lists all items
// @Description Retrieves all items info from items table in PostgreSQL
// @Tags item
// @Param update_data body item.ListItemsRequest true "list item data"
// @Success 200 {object} item.ListItemsResponse
// @Failure 500 {object} string "Server error while listing items"
// @Router /item-system/items [post]
func (h *Handler) ListItems(c *gin.Context) {
	h.Logger.Info("ListItems method is starting")

	var req pb.ListItemsRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind item data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	items, err := h.ItemClient.ListItems(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to list items", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetItem godoc
// @Summary Gets an item
// @Description Retrieves item info from items table in PostgreSQL
// @Tags item
// @Param item_id path string true "Item ID"
// @Success 200 {object} item.ItemResponse
// @Failure 400 {object} string "Invalid item ID"
// @Failure 500 {object} string "Server error while getting item"
// @Router /item-system/items/{item_id} [get]
func (h *Handler) GetItem(c *gin.Context) {
	h.Logger.Info("GetItem method is starting")

	id := c.Param("item_id")

	req := pb.GetItemRequest{ItemId: id}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	item, err := h.ItemClient.GetItem(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to get item", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// SearchItems godoc
// @Summary Searches for items
// @Description Searches items info in items table in PostgreSQL
// @Tags item
// @Param update_data body item.SearchItemsRequest true "list item data"
// @Success 200 {object} item.ListItemsResponse
// @Failure 500 {object} string "Server error while searching items"
// @Router /item-system/items/search [post]
func (h *Handler) SearchItems(c *gin.Context) {
	h.Logger.Info("SearchItems method is starting")

	var req pb.SearchItemsRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind item data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	items, err := h.ItemClient.SearchItems(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to search items", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search items"})
		return
	}

	c.JSON(http.StatusOK, items)
}