package handlers

import "github.com/thechervonyashiy/todoapp/pkg/services"

type Handler struct {
	services *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{services: service}
}
