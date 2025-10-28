CREATE TABLE IF NOT EXISTS products
(
    id            SERIAL PRIMARY KEY,
    product_name  VARCHAR NOT NULL,
    manufacturer  VARCHAR NOT NULL,
    product_count INT     NOT NULL,
    price         FLOAT   NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);