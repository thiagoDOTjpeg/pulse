package service

import (
	"context"
	"log/slog"
	"thiagodotjpeg/pulse/internal/monitor/model"
)

type MonitorService struct {
	Create  func()
	Update  func()
	Delete  func()
	GetById func()
	GetAll  func()
}

type MonitorServiceImpl struct {
	service MonitorService
	logger  *slog.Logger
}

func NewMonitorService(service MonitorService, logger *slog.Logger) *MonitorServiceImpl {
	return &MonitorServiceImpl{ service: service, logger: logger }
}

func (s *MonitorServiceImpl) Create(ctx context.Context, url string) (model.Url, error) {
	entity, err := model.NewUrl(url)
	if err != nil {
		return model.Url{}, err
	}

	
}