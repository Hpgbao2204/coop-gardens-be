-- +goose Up
-- +goose StatementBegin
CREATE TABLE crop_growth_logs (
    id SERIAL PRIMARY KEY,
    crop_id INT NOT NULL,
    log_date TIMESTAMPTZ NOT NULL,
    growth_stage VARCHAR(255),
    height DOUBLE PRECISION,
    health_status VARCHAR(255),
    notes TEXT,
    image_url VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_crop
        FOREIGN KEY(crop_id)
        REFERENCES crops(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- ALTER TABLE crops ADD COLUMN growth_stage VARCHAR(255);
-- +goose StatementEnd
INSERT INTO crop_growth_logs (crop_id, log_date, growth_stage, height, health_status, notes, image_url) VALUES
(1, NOW() - INTERVAL '10 days', 'Gieo mạ', 10.5, 'Tốt', 'Cây phát triển bình thường', NULL),
(2, NOW() - INTERVAL '5 days', 'Nảy mầm', 15.3, 'Trung bình', 'Có dấu hiệu thiếu nước', NULL),
(3, NOW() - INTERVAL '2 days', 'Sắp thu hoạch', 80.0, 'Rất tốt', 'Chuẩn bị thu hoạch', NULL);

-- Cập nhật growth_stage cho bảng crops
UPDATE crops SET growth_stage = 'Gieo mạ' WHERE id = 1;
UPDATE crops SET growth_stage = 'Nảy mầm' WHERE id = 2;
UPDATE crops SET growth_stage = 'Sắp thu hoạch' WHERE id = 3;
-- +goose Down
-- +goose StatementBegin
ALTER TABLE crops DROP COLUMN IF EXISTS growth_stage;
DROP TABLE IF EXISTS crop_growth_logs;
-- +goose StatementEnd

