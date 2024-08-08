package rest

import "github.com/olenka--91/reminder-app/internal/service"

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}
