package api

import (
	"api-gateway/api/handler"
	// "api-gateway/api/middleware"
	"api-gateway/config"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Item System
// @version 1.0
// @description API Gateway of User Item System
// @host localhost:8080
// BasePath: /
func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/item-system")
	// api.Use(middleware.Check)

	h := handler.NewHandler(cfg)

	u := api.Group("/users")
	{
		u.GET("/:user_id", h.GetUserProfile)
		u.PUT("/:user_id", h.UpdateUserProfile)
		u.DELETE("/:user_id", h.DeleteUser)
		u.POST("", h.GetUsers)
		u.GET("/:user_id/eco-points", h.GetEcoPoints)
		u.PUT("/:user_id/eco-points", h.AddEcoPoints)
		u.POST("/:user_id/eco-points/history", h.GetEcoPointsHistory)
	}

	category := api.Group("/category")
	{
		category.POST("/catogories", h.AddItemCategory)

	}
	

	item := api.Group("items")
	{
		item.POST("/addItem", h.AddItem)
		item.PUT("/:item_id", h.UpdateItem)
		item.DELETE("/:item_id", h.DeleteItem)
		item.POST("", h.ListItems)
		item.GET("/:item_id", h.GetItem)
		item.POST("/search", h.SearchItems)
	
	}

	ecoChannels := api.Group("ecosystem")
	{
		ecoChannels.POST("eco-challenge", h.CreateEcoChallenge)
        ecoChannels.POST("/participate", h.ParticipateEcoChallenge)
        ecoChannels.PUT("update", h.UpdateEcoChallengeProgress)
	}

	rating := api.Group("ratings")
	{
		rating.POST("add", h.AddRating)
		rating.POST("GetAll", h.GetRatings)
	
	}

	recycling := api.Group("recyclings")
	{
		recycling.POST("", h.AddRecyclingCenter)
		recycling.GET("search", h.SearchRecyclingCenters)
		recycling.GET("", h.SubmitItemsForRecycling)
	}

	statistics := api.Group("statistics")
	{
		statistics.POST("", h.Statistics)
	}
	

	swap := api.Group("swaps")
	{
		swap.POST("/", h.SendSwapRequest)
		swap.PUT("/accept", h.AcceptSwapRequest)
		swap.PUT("/:swap_id", h.ListSwapRequests)
		swap.PUT("/reject", h.RejectSwapRequest)
	}

	return router
}