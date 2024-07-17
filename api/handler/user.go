package handler

import (
	_ "api-gateway/genproto/authentication"
	"fmt"

	// "api-gateway/genproto/user"
	pb "api-gateway/genproto/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GetUserProfile godoc
// @Summary Gets user profile
// @Description Retrieves user profile info from PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} user.GetUserProfileResponse
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while getting user profile"
// @Router /item-system/users/{user_id} [get]
func (h *Handler) GetUserProfile(c *gin.Context) {
	h.Logger.Info("GetUser method is starting")

	id := c.Param("user_id")

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.UserClient.GetUserProfile(ctx, &pb.UserID{UserId: id})
	if err != nil {
		h.Logger.Error("failed to get user profile", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile godoc
// @Summary Updates user profile
// @Description Updates user profile info in PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Param new_info body user.UpdateUserProfileRequest true "Update user info"
// @Success 200 {object} user.UpdateProfileResponse
// @Failure 400 {object} string "Invalid user ID or data"
// @Failure 500 {object} string "Server error while updating user profile"
// @Router /item-system/users/{user_id} [put]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	h.Logger.Info("UpdateUserProfile method is starting")

	id := c.Param("user_id")

	var userProfile pb.UpdateUserProfileRequest
	err := c.ShouldBind(&userProfile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind user profile data", "error", err)
		return
	}
	userProfile.UserId = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.UserClient.UpdateUserProfile(ctx, &userProfile)
	if err != nil {
		h.Logger.Error("failed to update user profile", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Deletes a user
// @Description Removes user info from PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} string "User deleted successfully"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while deleting user"
// @Router /item-system/users/{user_id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	h.Logger.Info("DeleteUser method is starting")

	id := c.Param("user_id")

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err := h.UserClient.DeleteUser(ctx, &pb.DeleteUserRequest{UserId: id})
	if err != nil {
		h.Logger.Error("failed to delete user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, "User deleted successfully")
}

// GetUsers godoc
// @Summary Gets list of users
// @Description Retrieves list of users from PostgreSQL
// @Tags user
// @Param new_info body user.GetUsersRequest true "filter user info"
// @Success 200 {object} []user.GetUsersResponse
// @Failure 500 {object} string "Server error while getting users"
// @Router /item-system/users [post]
func (h *Handler) GetUsers(c *gin.Context) {
	h.Logger.Info("GetUsers method is starting")
	var filter pb.GetUsersRequest
	err := c.ShouldBind(&filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind user profile data", "error", err)
		return
	}

	fmt.Println(&filter)
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	users, err := h.UserClient.GetUsers(ctx, &filter)
	if err != nil {
		h.Logger.Error("failed to get users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users.Users)
}

// GetEcoPoints godoc
// @Summary Gets eco points of a user
// @Description Retrieves eco points info from PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} user.GetEcoPointsResponse
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while getting eco points"
// @Router /item-system/users/{user_id}/eco-points [get]
func (h *Handler) GetEcoPoints(c *gin.Context) {
	h.Logger.Info("GetEcoPoints method is starting")

	id := c.Param("user_id")

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ecoPoints, err := h.UserClient.GetEcoPoints(ctx, &pb.GetEcoPointsRequest{UserId: id})
	if err != nil {
		h.Logger.Error("failed to get eco points", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get eco points"})
		return
	}

	c.JSON(http.StatusOK, ecoPoints)
}

// AddEcoPoints godoc
// @Summary Adds eco points to a user
// @Description Inserts eco points info into PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Param points body user.AddEcoPointsRequest true "Eco points info"
// @Success 200 {object} user.AddEcoPointsResponse
// @Failure 400 {object} string "Invalid user ID or data"
// @Failure 500 {object} string "Server error while adding eco points"
// @Router /item-system/users/{user_id}/eco-points [put]
func (h *Handler) AddEcoPoints(c *gin.Context) {
	h.Logger.Info("AddEcoPoints method is starting")

	id := c.Param("user_id")

	var ecoPoints pb.AddEcoPointsRequest

	err := c.ShouldBind(&ecoPoints)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind eco points data", "error", err)
		return
	}
	ecoPoints.UserId = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ecoPoint, err := h.UserClient.AddEcoPoints(ctx, &ecoPoints)
	if err != nil {
		h.Logger.Error("failed to add eco points", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add eco points"})
		return
	}

	c.JSON(http.StatusOK, ecoPoint)
}

// GetEcoPointsHistory godoc
// @Summary Gets eco points history of a user
// @Description Retrieves eco points history from PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Param points body user.GetEcoPointsHistoryRequest true "Eco points history info"
// @Success 200 {object} []user.GetEcoPointsHistoryResponse
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while getting eco points history"
// @Router /item-system/users/{user_id}/eco-points/history [post]
func (h *Handler) GetEcoPointsHistory(c *gin.Context) {
	h.Logger.Info("GetEcoPointsHistory method is starting")

	id := c.Param("user_id")

	var filter pb.GetEcoPointsHistoryRequest

	err := c.ShouldBind(&filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		h.Logger.Error("failed to bind eco points data", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	filter.UserId = id

	history, err := h.UserClient.GetEcoPointsHistory(ctx, &filter)
	if err != nil {
		h.Logger.Error("failed to get eco points history", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get eco points history"})
		return
	}

	c.JSON(http.StatusOK, history)
}
