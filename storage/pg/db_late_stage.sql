-- A simple database for an ecommerce application
-- Written in PostgreSQL

--
-- Enums
--

CREATE TYPE payment_status AS ENUM (
    'pending',
    'paid',
    'refunded',
    'cancelled'
);
CREATE TYPE shipping_status AS ENUM (
    'pending',
    'in_transit',
    'delivered',
    'cancelled',
    'automatic_return'
);
CREATE TYPE shipping_service AS ENUM (
    'postnord',
    'bring',
    'ups',
    'dhl'
);
CREATE TYPE shipping_method_type AS ENUM (
    'pickup',
    'delivery'
);
CREATE TYPE discount_type AS ENUM (
    'automatic',
    'code',
    'manual'
);
CREATE TYPE discount_application_type AS ENUM (
    'order',
    'order_line'
);
CREATE TYPE discount_amount_type AS ENUM (
    'fixed',
    'percentage'
);

--
-- Tables
--

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    image_url TEXT[],
    "description" TEXT,
    meta JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE variants (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    sku VARCHAR(255) NOT NULL,
    ean VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    "size" VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    original_price INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    meta JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    parent_id INTEGER REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    meta JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
    "address" VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,
    customer_id INTEGER REFERENCES customers(id),
    shipping_address_id INTEGER REFERENCES addresses(id),
    billing_address_id INTEGER REFERENCES addresses(id),
    "status" VARCHAR(255) NOT NULL,
    meta JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_lines (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    variant_id INTEGER REFERENCES variants(id) ON DELETE CASCADE,
    price INTEGER NOT NULL,
    original_price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE costs_breakdown (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    total_amount INTEGER NOT NULL,
    tax_amount INTEGER NOT NULL,
    sub_total INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    "status" payment_status NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE fullfillments (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    shipping_address_id INTEGER REFERENCES addresses(id),
    shipping_cost INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    free_shipping BOOLEAN NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE fulfillment_lines (
    id SERIAL PRIMARY KEY,
    fulfillment_id INTEGER REFERENCES fullfillments(id) ON DELETE CASCADE,
    order_line_id INTEGER REFERENCES order_lines(id) ON DELETE CASCADE,
    shipping_method shipping_method_type NOT NULL,
    "status" shipping_status NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE discounts (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    order_line_id INTEGER REFERENCES order_lines(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    "type" discount_type NOT NULL,
    application_type discount_application_type NOT NULL,
    amount_type discount_amount_type NOT NULL,
    amount INTEGER NOT NULL,
    uses INTEGER NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

---
--- Indexes
---

CREATE INDEX categories_product_id_idx ON categories (product_id);
CREATE INDEX categories_name_idx ON categories (name);

CREATE INDEX tags_product_id_idx ON tags (product_id);
CREATE INDEX tags_name_idx ON tags (name);

CREATE INDEX orders_customer_id_idx ON orders (customer_id);

CREATE INDEX order_lines_order_id_idx ON order_lines (order_id);

CREATE INDEX transactions_order_id_idx ON transactions (order_id);

CREATE INDEX fullfillments_order_id_idx ON fullfillments (order_id);

CREATE INDEX fulfillment_lines_fulfillment_id_idx ON fulfillment_lines (fulfillment_id);

CREATE INDEX discounts_order_id_idx ON discounts (order_id);

CREATE INDEX discounts_order_line_id_idx ON discounts (order_line_id);

--
-- Functions
--

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

--
-- Triggers
--

CREATE TRIGGER products_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER variants_updated_at
BEFORE UPDATE ON variants
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER categories_updated_at
BEFORE UPDATE ON categories
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER tags_updated_at
BEFORE UPDATE ON tags
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER customers_updated_at
BEFORE UPDATE ON customers
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER addresses_updated_at
BEFORE UPDATE ON addresses
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER orders_updated_at
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER order_lines_updated_at
BEFORE UPDATE ON order_lines
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER costs_breakdown_updated_at
BEFORE UPDATE ON costs_breakdown
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER transactions_updated_at
BEFORE UPDATE ON transactions
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER fullfillments_updated_at
BEFORE UPDATE ON fullfillments
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER fulfillment_lines_updated_at
BEFORE UPDATE ON fulfillment_lines
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER discounts_updated_at
BEFORE UPDATE ON discounts
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

--
-- Permissions
--

ALTER TABLE products ENABLE ROW LEVEL SECURITY;
ALTER TABLE variants ENABLE ROW LEVEL SECURITY;
ALTER TABLE categories ENABLE ROW LEVEL SECURITY;
ALTER TABLE tags ENABLE ROW LEVEL SECURITY;
ALTER TABLE customers ENABLE ROW LEVEL SECURITY;
ALTER TABLE addresses ENABLE ROW LEVEL SECURITY;
ALTER TABLE orders ENABLE ROW LEVEL SECURITY;
ALTER TABLE order_lines ENABLE ROW LEVEL SECURITY;
ALTER TABLE costs_breakdown ENABLE ROW LEVEL SECURITY;
ALTER TABLE transactions ENABLE ROW LEVEL SECURITY;
ALTER TABLE fullfillments ENABLE ROW LEVEL SECURITY;
ALTER TABLE fulfillment_lines ENABLE ROW LEVEL SECURITY;
ALTER TABLE discounts ENABLE ROW LEVEL SECURITY;

