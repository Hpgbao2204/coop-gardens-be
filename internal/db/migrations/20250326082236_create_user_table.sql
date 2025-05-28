-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    full_name TEXT NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    google_id TEXT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO users (email, password, full_name, is_verified, google_id) VALUES 
('admin@example.com', 'admmin', 'Nguyễn Văn A', TRUE, NULL),
('farmer@example.com', 'farmer', 'Lê Văn C', FALSE, NULL),
('customer@example.com', 'customer', 'Phạm Thị D', TRUE, 'google_1234'),

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL 
);

INSERT INTO roles (name) VALUES ('User'), ('Admin'), ('Farmer');
CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

INSERT INTO user_roles (user_id, role_id) 
SELECT u.id, r.id 
FROM users u, roles r
WHERE (u.email = 'admin@example.com' AND r.name = 'Admin')
   OR (u.email = 'farmer@example.com' AND r.name = 'Nông dân')
   OR (u.email = 'customer@example.com' AND r.name = 'Khách hàng')

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd