-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    -- Kiểm tra bảng seasons đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'seasons'
    ) THEN
        CREATE TABLE seasons (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            start_date TIMESTAMP NOT NULL,
            end_date TIMESTAMP NOT NULL,
            status VARCHAR(50) DEFAULT 'Planning',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;

    -- Thêm cột season_id nếu chưa tồn tại
    IF NOT EXISTS (
        SELECT FROM information_schema.columns 
        WHERE table_name = 'crops' AND column_name = 'season_id'
    ) THEN
        ALTER TABLE crops ADD COLUMN season_id INT;
    END IF;

    -- Thêm constraint nếu chưa tồn tại
    IF NOT EXISTS (
        SELECT FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_crops_season'
    ) THEN
        ALTER TABLE crops ADD CONSTRAINT fk_crops_season 
        FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE SET NULL;
    END IF;
END $$;
-- +goose StatementEnd

INSERT INTO seasons (name, start_date, end_date, status) VALUES
('Mùa Xuân 2025', '2025-02-01 00:00:00', '2025-04-30 23:59:59', 'Ongoing'),
('Mùa Hè 2025', '2025-05-01 00:00:00', '2025-07-31 23:59:59', 'Planning'),
('Mùa Thu 2025', '2025-08-01 00:00:00', '2025-10-31 23:59:59', 'Planning'),
('Mùa Đông 2025', '2025-11-01 00:00:00', '2025-01-31 23:59:59', 'Planned');
-- +goose Down
-- +goose StatementBegin
ALTER TABLE crops DROP CONSTRAINT IF EXISTS fk_crops_season;
ALTER TABLE crops DROP COLUMN IF EXISTS season_id;
DROP TABLE IF EXISTS seasons CASCADE;
-- +goose StatementEnd