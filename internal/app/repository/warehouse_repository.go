package repository

import (
	"context"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/entity"
)

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse entity.Warehouse) error
	List(ctx context.Context) ([]entity.Warehouse, error)
}
