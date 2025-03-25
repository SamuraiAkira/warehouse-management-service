package entity

import (
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID      uuid.UUID `json:"id"`
	Address string    `json:"address"`
}

type Product struct {
	ID              uuid.UUID         `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Characteristics map[string]string `json:"characteristics"`
	Weight          float64           `json:"weight"`
	Barcode         string            `json:"barcode"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

type Inventory struct {
	WarehouseID uuid.UUID `json:"warehouse_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Discount    float64   `json:"discount"`
}

type Sale struct {
	ID          uuid.UUID `json:"id"`
	WarehouseID uuid.UUID `json:"warehouse_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
	SoldAt      time.Time `json:"sold_at"`
}

type Analytics struct {
	WarehouseID uuid.UUID `json:"warehouse_id"`
	ProductID   uuid.UUID `json:"product_id"`
	TotalSold   int       `json:"total_sold"`
	TotalSum    float64   `json:"total_sum"`
}
