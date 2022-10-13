package healthcheck

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	service "server/internal/adapters/api"
	"server/internal/domain/entity"
	"time"
)

type ServerHealthCheck struct {
	item   service.ItemUsecase
	logger *zap.Logger
}

func NewServerHealthCheck(item service.ItemUsecase, logger *zap.Logger) ServerHealthCheck {
	return ServerHealthCheck{
		item:   item,
		logger: logger,
	}
}

var ExpectedItem = entity.Item{
	Name:        "test_name",
	Data:        []byte("test_data"),
	Description: "test_description",
}

func (s ServerHealthCheck) HealthCheck(w http.ResponseWriter, r *http.Request) {

	timeout := time.Second * 5

	// // parse custom timeout

	if value := r.URL.Query().Get("timeout"); value != "" {
		var err error

		if timeout, err = time.ParseDuration(value); err != nil {
			WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	// // create context

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// // create item

	id, err := s.item.CreateItem(ctx, ExpectedItem.Name, ExpectedItem.Data, ExpectedItem.Description)
	if err != nil {
		s.logger.Error("HealthCheck: CreateItem", zap.Error(err))
		WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	// // first get: without cache

	_, err = s.item.GetItem(ctx, id)
	if err != nil {
		s.logger.Error("HealthCheck: GetItem", zap.Error(err))
		WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	// // second get: from cache

	_, err = s.item.GetItem(ctx, id)
	if err != nil {
		s.logger.Error("HealthCheck: GetItem", zap.Error(err))
		WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	// // delete item

	deleted, err := s.item.DeleteItem(ctx, id)
	if err != nil {
		s.logger.Error("HealthCheck: GetItem", zap.Error(err))
		WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		s.logger.Error("HealthCheck: DeleteItem", zap.Error(err))
		WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !deleted {
		s.logger.Error("HealthCheck: DeleteItem: Item not deleted")
		WriteJson(w, http.StatusInternalServerError, "DeleteItem: item not deleted")
		return
	}

	// // All Good

	WriteJson(w, http.StatusOK, "ok")
}
