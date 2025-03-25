package service

import (
	"context"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/entity"
	"github.com/SamuraiAkira/warehouse-management-service/internal/app/repository"
	"github.com/google/uuid"
)

type WarehouseService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseService(repo repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) Create(ctx context.Context, address string) (entity.Warehouse, error) {
	warehouse := entity.Warehouse{
		ID:      uuid.New(),
		Address: address,
	}

	if err := s.repo.Create(ctx, warehouse); err != nil {
		return entity.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *WarehouseService) List(ctx context.Context) ([]entity.Warehouse, error) {
	return s.repo.List(ctx)
}
