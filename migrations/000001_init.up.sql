CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE warehouses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    address TEXT NOT NULL
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    characteristics JSONB,
    weight FLOAT,
    barcode TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE inventory (
    warehouse_id UUID REFERENCES warehouses(id),
    product_id UUID REFERENCES products(id),
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    price FLOAT NOT NULL CHECK (price > 0),
    discount FLOAT DEFAULT 0 CHECK (discount >= 0 AND discount <= 1),
    PRIMARY KEY (warehouse_id, product_id)
);

CREATE TABLE sales (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID REFERENCES warehouses(id),
    product_id UUID REFERENCES products(id),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    total_price FLOAT NOT NULL CHECK (total_price > 0),
    sold_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE analytics (
    warehouse_id UUID REFERENCES warehouses(id),
    product_id UUID REFERENCES products(id),
    total_sold INTEGER NOT NULL DEFAULT 0,
    total_sum FLOAT NOT NULL DEFAULT 0,
    PRIMARY KEY (warehouse_id, product_id)
);