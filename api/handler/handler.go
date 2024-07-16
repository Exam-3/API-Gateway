package handler

import (
	"api-gateway/config"
	"api-gateway/genproto/item"
	"api-gateway/genproto/user"
	"api-gateway/pkg"
	"api-gateway/pkg/logger"
	"log/slog"
)

type Handler struct {
	UserClient        user.UserClient
	RestaurantClient  item.RestaurantClient
	Logger            *slog.Logger
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		UserClient: pkg.NewUserClient(cfg),
		ItemClient: pkg.NewItemClient(cfg),
		Logger:     logger.NewLogger(),
	}
}
