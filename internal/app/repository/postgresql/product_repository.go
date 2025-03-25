package postgresql

import (
	"context"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product entity.Product) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO products 
        (id, name, description, characteristics, weight, barcode) 
        VALUES ($1, $2, $3, $4, $5, $6)`,
		product.ID, product.Name, product.Description,
		product.Characteristics, product.Weight, product.Barcode)
	return err
}
