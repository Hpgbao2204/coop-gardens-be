-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'Pending',
    assigned_to UUID REFERENCES users(id) ON DELETE SET NULL,
    season_id INT REFERENCES seasons(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE task_crops (
    task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
    crop_id INT REFERENCES crops(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, crop_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
