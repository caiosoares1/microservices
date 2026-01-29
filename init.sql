CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `payment`;
CREATE DATABASE IF NOT EXISTS `shipping`;

USE `order`;

CREATE TABLE IF NOT EXISTS stock_items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    product_code VARCHAR(191) UNIQUE,
    name VARCHAR(255),
    unit_price FLOAT,
    quantity INT,
    INDEX idx_stock_items_deleted_at (deleted_at),
    INDEX idx_stock_items_product_code (product_code)
);

INSERT INTO stock_items (created_at, updated_at, product_code, name, unit_price, quantity) VALUES
    (NOW(), NOW(), 'PROD001', 'Camiseta Básica', 49.90, 100),
    (NOW(), NOW(), 'PROD002', 'Calça Jeans', 129.90, 50),
    (NOW(), NOW(), 'PROD003', 'Tênis Esportivo', 199.90, 30),
    (NOW(), NOW(), 'PROD004', 'Boné', 39.90, 80),
    (NOW(), NOW(), 'PROD005', 'Mochila', 89.90, 40),
    (NOW(), NOW(), 'PROD006', 'Relógio Digital', 149.90, 25),
    (NOW(), NOW(), 'PROD007', 'Óculos de Sol', 79.90, 60),
    (NOW(), NOW(), 'PROD008', 'Carteira de Couro', 59.90, 45),
    (NOW(), NOW(), 'PROD009', 'Cinto', 34.90, 70),
    (NOW(), NOW(), 'PROD010', 'Meias (Pack 3)', 29.90, 120)
ON DUPLICATE KEY UPDATE updated_at = NOW();
