-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'crops'
    ) THEN
        CREATE TABLE crops (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            type VARCHAR(100),
            status VARCHAR(50) DEFAULT 'Planted',
            planted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
    
    IF NOT EXISTS (
        SELECT FROM information_schema.columns 
        WHERE table_name = 'crops' AND column_name = 'growth_stage'
    ) THEN
        ALTER TABLE crops ADD COLUMN growth_stage VARCHAR(255);
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE crops ADD COLUMN season_id INT;
ALTER TABLE crops ADD CONSTRAINT fk_crops_season FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- Thêm dữ liệu mẫu vào bảng seasons
INSERT INTO seasons (name, start_date, end_date, status) VALUES
('Mùa Xuân 2025', '2025-02-01 00:00:00', '2025-04-30 23:59:59', 'Ongoing'),
('Mùa Hè 2025', '2025-05-01 00:00:00', '2025-07-31 23:59:59', 'Planning'),
('Mùa Thu 2025', '2025-08-01 00:00:00', '2025-10-31 23:59:59', 'Planning'),
('Mùa Đông 2025', '2025-11-01 00:00:00', '2025-01-31 23:59:59', 'Planned');
-- +goose Down
-- +goose StatementBegin
ALTER TABLE crops DROP COLUMN IF EXISTS growth_stage;
-- +goose StatementEnd
