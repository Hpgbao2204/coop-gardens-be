-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    quantity NUMERIC NOT NULL,
    unit VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'In Stock',
    created_by UUID NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_inventories_users FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO inventories (name, category, quantity, unit, created_by) VALUES
('Phân bón NPK', 'Phân bón', 500, 'Kg', (SELECT id FROM users WHERE email = 'admin@example.com')),
('Hạt giống lúa', 'Hạt giống', 10000, 'Cây', (SELECT id FROM users WHERE email = 'farmer@example.com')),

CREATE TABLE crop_inventories (
    id SERIAL PRIMARY KEY,
    crop_id INT NOT NULL,
    inventory_id INT NOT NULL,
    quantity NUMERIC NOT NULL,
    CONSTRAINT fk_cropinventories_crops FOREIGN KEY (crop_id) REFERENCES crops(id) ON DELETE CASCADE,
    CONSTRAINT fk_cropinventories_inventories FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE,
    CONSTRAINT uq_crop_inventory UNIQUE (crop_id, inventory_id)
);
-- Thêm dữ liệu mẫu vào bảng crop_inventories
INSERT INTO crop_inventories (crop_id, inventory_id, quantity) VALUES
(1, 1, 300),  -- Lúa sử dụng phân bón NPK
(2, 2, 5000),  -- Ngô sử dụng hạt giống lúa


CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    inventory_id INT NOT NULL,
    type VARCHAR(50) NOT NULL, -- "import" hoặc "export"
    quantity NUMERIC NOT NULL,
    performed_by UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transactions_inventory FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE,
    CONSTRAINT fk_transactions_users FOREIGN KEY (performed_by) REFERENCES users(id) ON DELETE CASCADE
);

-- Thêm dữ liệu mẫu vào bảng inventory_transactions
INSERT INTO inventory_transactions (inventory_id, type, quantity, performed_by) VALUES
(1, 'import', 200, (SELECT id FROM users WHERE email = 'admin@example.com')),
(2, 'export', 1000, (SELECT id FROM users WHERE email = 'farmer@example.com')),
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS inventory_transactions;
DROP TABLE IF EXISTS crop_inventories;
DROP TABLE IF EXISTS inventories;
-- +goose StatementEnd
