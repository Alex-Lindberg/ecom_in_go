-- A simple database for an ecommerce application
-- Written in PostgreSQL


CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO products ("name") VALUES ('Product 1');
INSERT INTO products ("name") VALUES ('Product 2');
INSERT INTO products ("name") VALUES ('Product 3');
INSERT INTO products ("name") VALUES ('Product 4');
INSERT INTO products ("name") VALUES ('Product 5');

CREATE TABLE variants (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO variants (product_id, "name") VALUES (1, 'Variant 1A');
INSERT INTO variants (product_id, "name") VALUES (1, 'Variant 1B');
INSERT INTO variants (product_id, "name") VALUES (2, 'Variant 2A');
INSERT INTO variants (product_id, "name") VALUES (3, 'Variant 3A');
INSERT INTO variants (product_id, "name") VALUES (3, 'Variant 3B');
INSERT INTO variants (product_id, "name") VALUES (3, 'Variant 3C');
INSERT INTO variants (product_id, "name") VALUES (4, 'Variant 4A');
INSERT INTO variants (product_id, "name") VALUES (5, 'Variant 5A');

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO customers ("name", email) VALUES ('Customer 1', 'customer1@example.com');
INSERT INTO customers ("name", email) VALUES ('Customer 2', 'customer2@example.com');
INSERT INTO customers ("name", email) VALUES ('Customer 3', 'customer3@example.com');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,
    customer_id INTEGER REFERENCES customers(id),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO orders (order_number, customer_id) VALUES ('ORDER123', 1);
INSERT INTO orders (order_number, customer_id) VALUES ('ORDER124', 1);
INSERT INTO orders (order_number, customer_id) VALUES ('ORDER125', 1);
INSERT INTO orders (order_number, customer_id) VALUES ('ORDER126', 2);
INSERT INTO orders (order_number, customer_id) VALUES ('ORDER127', 2);
INSERT INTO orders (order_number, customer_id) VALUES ('ORDER128', 3);

CREATE TABLE order_lines (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    variant_id INTEGER REFERENCES variants(id) ON DELETE CASCADE,
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (1, 1, 1, 2000, 2);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (1, 2, 3, 3000, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (2, 2, 2, 3000, 3);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (3, 4, 1, 3000, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (3, 5, 1, 3000, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (3, 3, 1, 3000, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (3, 3, 2, 3000, 1);


-- Costs breakdown table for order payment and shipping costs

--
-- Indexes
--

CREATE INDEX orders_customer_id_idx ON orders (customer_id);

CREATE INDEX order_lines_order_id_idx ON order_lines (order_id);

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

CREATE TRIGGER customers_updated_at
BEFORE UPDATE ON customers
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

--
-- Permissions
--

ALTER TABLE products ENABLE ROW LEVEL SECURITY;
ALTER TABLE variants ENABLE ROW LEVEL SECURITY;
ALTER TABLE customers ENABLE ROW LEVEL SECURITY;
ALTER TABLE orders ENABLE ROW LEVEL SECURITY;
ALTER TABLE order_lines ENABLE ROW LEVEL SECURITY;

