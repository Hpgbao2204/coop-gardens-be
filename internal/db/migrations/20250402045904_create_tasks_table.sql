-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    -- Kiểm tra bảng tasks đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'tasks'
    ) THEN
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
    END IF;

    -- Kiểm tra bảng task_crops đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'task_crops'
        ) THEN
            CREATE TABLE task_crops (
            task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
            crop_id INT REFERENCES crops(id) ON DELETE CASCADE,
            PRIMARY KEY (task_id, crop_id)
        );
    END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task_crops CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
-- +goose StatementEnd