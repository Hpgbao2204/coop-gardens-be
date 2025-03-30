-- +goose Up
-- +goose StatementBegin
CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(50) DEFAULT 'Planning',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE crops ADD COLUMN season_id INT;
ALTER TABLE crops ADD CONSTRAINT fk_crops_season FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE SET NULL;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
