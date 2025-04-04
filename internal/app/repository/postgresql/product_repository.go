package postgresql

import (
	"context"
	"errors"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Объявляем общие ошибки репозитория
var (
	ErrNotFound = errors.New("record not found")
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

func (r *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (entity.Product, error) {
	var product entity.Product
	err := r.db.QueryRow(ctx, `
		SELECT id, name, description, characteristics, weight, barcode, created_at, updated_at
		FROM products WHERE id = $1`, id).Scan(
		&product.ID, &product.Name, &product.Description,
		&product.Characteristics, &product.Weight, &product.Barcode,
		&product.CreatedAt, &product.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Product{}, ErrNotFound
	}

	return product, err
}
