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

-- +goose Down
-- +goose StatementBegin
ALTER TABLE crops DROP COLUMN IF EXISTS growth_stage;
-- +goose StatementEnd
