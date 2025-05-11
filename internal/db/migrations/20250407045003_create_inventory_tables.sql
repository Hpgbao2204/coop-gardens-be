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

CREATE TABLE crop_inventories (
    id SERIAL PRIMARY KEY,
    crop_id INT NOT NULL,
    inventory_id INT NOT NULL,
    quantity NUMERIC NOT NULL,
    CONSTRAINT fk_cropinventories_crops FOREIGN KEY (crop_id) REFERENCES crops(id) ON DELETE CASCADE,
    CONSTRAINT fk_cropinventories_inventories FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE,
    CONSTRAINT uq_crop_inventory UNIQUE (crop_id, inventory_id)
);

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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS inventory_transactions;
DROP TABLE IF EXISTS crop_inventories;
DROP TABLE IF EXISTS inventories;
-- +goose StatementEnd
