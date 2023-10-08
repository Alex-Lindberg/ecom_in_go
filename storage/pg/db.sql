-- A simple database for an ecommerce application
-- Written in PostgreSQL


CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE ("name")
);

INSERT INTO products ("name") VALUES ('Product A');
INSERT INTO products ("name") VALUES ('Product B');
INSERT INTO products ("name") VALUES ('Product C');


CREATE TABLE variants (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    sku VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (sku)
);
ALTER TABLE variants
ADD CONSTRAINT unique_productId_name UNIQUE (product_id, "name");


INSERT INTO variants (product_id, sku, "name") VALUES (1,'VA1', 'Variant A1');
INSERT INTO variants (product_id, sku, "name") VALUES (1,'VA2', 'Variant A2');
INSERT INTO variants (product_id, sku, "name") VALUES (2,'VB1', 'Variant B1');
INSERT INTO variants (product_id, sku, "name") VALUES (3,'VC1', 'Variant C1');


CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO customers ("name", email) VALUES ('Customer 1', 'customer1@example.com');
INSERT INTO customers ("name", email) VALUES ('Customer 2', 'customer2@example.com');
INSERT INTO customers ("name", email) VALUES ('Customer 3', 'customer3@example.com');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,             -- The id the customer sees
    order_reference VARCHAR(255) NOT NULL,          -- The id the business/platform uses
    customer_id INTEGER REFERENCES customers(id),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
ALTER TABLE orders
ADD CONSTRAINT unique_orderNumber UNIQUE (order_number);
ALTER TABLE orders
ADD CONSTRAINT unique_orderReference UNIQUE (order_reference);


INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('123', 'ORDER123', 1);
INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('124', 'ORDER124', 1);
INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('125', 'ORDER125', 1);
INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('126', 'ORDER126', 2);
INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('127', 'ORDER127', 2);
INSERT INTO orders (order_number, order_reference, customer_id) VALUES ('128', 'ORDER128', 3);

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
ALTER TABLE order_lines 
ADD CONSTRAINT unique_orderId_productId_variantId UNIQUE (order_id, product_id, variant_id);

-- Order 1
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (1, 1, 1, 2000, 2);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (1, 2, 3, 2500, 1);

-- Order 2
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (2, 1, 2, 2200, 3);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (2, 3, 4, 1800, 1);

-- Order 3
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (3, 2, 3, 2500, 2);

-- Order 4
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (4, 3, 4, 1800, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (4, 1, 1, 2000, 2);

-- Order 5
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (5, 1, 2, 2200, 1);
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (5, 2, 3, 2500, 1);

-- Order 6
INSERT INTO order_lines (order_id, product_id, variant_id, price, quantity) VALUES (6, 3, 4, 1800, 1);

-- Costs breakdown table for order payment and shipping costs

--
-- Indexes
--

CREATE INDEX variants_product_id_idx ON variants (product_id);

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

