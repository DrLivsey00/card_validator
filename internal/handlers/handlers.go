package handlers

import "github.com/DrLivsey00/card_checker/pkg/service"

type Handlers struct {
	service *service.Service
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{
		services,
	}
}
