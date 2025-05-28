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


-- Thêm dữ liệu mẫu vào bảng tasks
INSERT INTO tasks (title, description, status, assigned_to, season_id) VALUES
('Kiểm tra độ ẩm đất', 'Đo độ ẩm đất để đảm bảo đủ nước cho cây lúa', 'In Progress', NULL, 1),
('Bón phân cho ngô', 'Bón phân hữu cơ cho cây ngô vào giai đoạn sinh trưởng', 'Pending', NULL, 2),
('Thu hoạch cà phê', 'Thu hoạch cà phê sau khi chín', 'Pending', NULL, 4);

-- Gán cây trồng cho các nhiệm vụ
INSERT INTO task_crops (task_id, crop_id) VALUES
(1, 1),  -- Kiểm tra độ ẩm cho Lúa
(2, 2),  -- Bón phân cho Ngô
(3, 8);  -- Thu hoạch cà phê
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task_crops CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
-- +goose StatementEnd