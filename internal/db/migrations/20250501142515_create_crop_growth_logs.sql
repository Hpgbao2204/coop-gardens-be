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

ALTER TABLE crops ADD COLUMN growth_stage VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE crops DROP COLUMN IF EXISTS growth_stage;
DROP TABLE IF EXISTS crop_growth_logs;
-- +goose StatementEnd

