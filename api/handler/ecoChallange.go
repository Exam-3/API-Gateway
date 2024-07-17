package handler

import (
	pb "api-gateway/genproto/item"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// CreateEcoChallenge godoc
// @Summary Creates a new eco challenge
// @Description Inserts new eco challenge info into eco_challenges table in PostgreSQL
// @Tags eco_challenge
// @Param new_data body item.CreateEcoChallengeRequest true "New data"
// @Success 200 {object} item.CreateEcoChallengeResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while creating eco challenge"
// @Router /item-system/ecosystem/eco-challenge [post]
func (h *Handler) CreateEcoChallenge(c *gin.Context) {
	h.Logger.Info("CreateEcoChallenge method is starting")
	
	var req pb.CreateEcoChallengeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind eco challenge data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	fmt.Println(&req)
	ecoChallenge, err := h.ItemClient.CreateEcoChallenge(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to create eco challenge", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create eco challenge"})
		return
	}

	c.JSON(http.StatusOK, ecoChallenge)
}



// ParticipateEcoChallenge godoc
// @Summary Participates in an eco challenge
// @Description Inserts new participation info into challenge_participations table in PostgreSQL
// @Tags eco_challenge
// @Param new_data body item.ParticipateEcoChallengeRequest true "New data"
// @Success 200 {object} item.ParticipateEcoChallengeResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while participating in eco challenge"
// @Router /item-system/ecosystem/participate [post]
func (h *Handler) ParticipateEcoChallenge(c *gin.Context) {
	h.Logger.Info("ParticipateEcoChallenge method is starting")

	var req pb.ParticipateEcoChallengeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind participation data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	participation, err := h.ItemClient.ParticipateEcoChallenge(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to participate in eco challenge", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to participate in eco challenge"})
		return
	}

	c.JSON(http.StatusOK, participation)
}

// UpdateEcoChallengeProgress godoc
// @Summary Updates progress in an eco challenge
// @Description Updates progress info in challenge_participations table in PostgreSQL
// @Tags eco_challenge
// @Param new_data body item.UpdateEcoChallengeProgressRequest true "New data"
// @Success 200 {object} item.UpdateEcoChallengeProgressResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while updating eco challenge progress"
// @Router /item-system/ecosystem/update [put]
func (h *Handler) UpdateEcoChallengeProgress(c *gin.Context) {
	h.Logger.Info("UpdateEcoChallengeProgress method is starting")

	var req pb.UpdateEcoChallengeProgressRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})

		log.Println(err)
		h.Logger.Error("failed to bind progress data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	progress, err := h.ItemClient.UpdateEcoChallengeProgress(ctx, &req)
	if err != nil {
		h.Logger.Error("failed to update eco challenge progress", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update eco challenge progress"})
		return
	}

	c.JSON(http.StatusOK, progress)
}
