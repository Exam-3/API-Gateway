package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	pb "api-gateway/genproto/item"
)

// SendSwapRequest godoc
// @Summary Send swap request
// @Description Sends a swap request to the service
// @Tags swap
// @Param swap body item.SendSwapRequestRequest true "Swap request info"
// @Success 200 {object} item.SwapResponse
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Server error while sending swap request"
// @Router /item-system/swaps [post]
func (h *Handler) SendSwapRequest(c *gin.Context) {
	h.Logger.Info("SendSwapRequest method is starting")

	var req pb.SendSwapRequestRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind swap request data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.SendSwapRequest(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to send swap request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send swap request"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AcceptSwapRequest godoc
// @Summary Accept swap request
// @Description Accepts a swap request
// @Tags swap
// @Param swap body item.AcceptSwapRequestRequest true "Swap request info"
// @Success 200 {object} item.SwapResponse
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Server error while accepting swap request"
// @Router /item-system/swaps/accept [put]
func (h *Handler) AcceptSwapRequest(c *gin.Context) {
	h.Logger.Info("AcceptSwapRequest method is starting")

	req := pb.AcceptSwapRequestRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind swap request data", "error", err)
		return
	}
	fmt.Println(&req)


	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.AcceptSwapRequest(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to accept swap request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept swap request"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// RejectSwapRequest godoc
// @Summary Reject swap request
// @Description Rejects a swap request
// @Tags swap
// @Param swap body item.RejectSwapRequestRequest true "Swap request info"
// @Success 200 {object} item.SwapResponse
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Server error while rejecting swap request"
// @Router /item-system/swaps/reject [put]
func (h *Handler) RejectSwapRequest(c *gin.Context) {
	h.Logger.Info("RejectSwapRequest method is starting")


	var req pb.RejectSwapRequestRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind swap request data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.RejectSwapRequest(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to reject swap request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject swap request"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListSwapRequests godoc
// @Summary List swap requests
// @Description Lists all swap requests
// @Tags swap
// @Param filter body item.ListSwapRequestsRequest true "Swap request filter"
// @Success 200 {object} item.ListSwapRequestsResponse
// @Failure 400 {object} string "Invalid filter"
// @Failure 500 {object} string "Server error while listing swap requests"
// @Router /item-system/swaps/list [post]
func (h *Handler) ListSwapRequests(c *gin.Context) {
	h.Logger.Info("ListSwapRequests method is starting")

	var req pb.ListSwapRequestsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind swap request filter data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ItemClient.ListSwapRequests(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to list swap requests", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list swap requests"})
		return
	}

	c.JSON(http.StatusOK, res)
}
