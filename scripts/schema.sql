CREATE DATABASE IF NOT EXISTS orders_db;
USE orders_db;
CREATE TABLE orders (
                        id VARCHAR(255) PRIMARY KEY,
                        price DECIMAL(10, 2),
                        tax DECIMAL(10, 2),
                        final_price DECIMAL(10, 2)
);