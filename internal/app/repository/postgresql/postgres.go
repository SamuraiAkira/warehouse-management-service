package postgresql

import (
	"context"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/entity"
	"github.com/SamuraiAkira/warehouse-management-service/internal/app/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type warehouseRepository struct {
	db *pgxpool.Pool
}

func NewWarehouseRepository(db *pgxpool.Pool) repository.WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) Create(ctx context.Context, warehouse entity.Warehouse) error {
	query := `INSERT INTO warehouses (id, address) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, warehouse.ID, warehouse.Address)
	return err
}

func (r *warehouseRepository) List(ctx context.Context) ([]entity.Warehouse, error) {
	query := `SELECT id, address FROM warehouses`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []entity.Warehouse
	for rows.Next() {
		var w entity.Warehouse
		if err := rows.Scan(&w.ID, &w.Address); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}
