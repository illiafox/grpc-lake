package http

import (
	"net/http"
	service "server/internal/adapters/api"
)

type ServerHealthCheck struct {
	item service.ItemUsecase
}

func NewServerHealthCheck(item service.ItemUsecase) ServerHealthCheck {
	return ServerHealthCheck{
		item: item,
	}
}

func (s ServerHealthCheck) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return s.item.HealthCheck()
}
