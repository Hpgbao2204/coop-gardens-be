-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    -- Kiểm tra bảng blogs đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'blogs'
    ) THEN
        CREATE TABLE blogs (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            author_id UUID NOT NULL REFERENCES users(id),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;

    -- Kiểm tra bảng comments đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'comments'
    ) THEN
        CREATE TABLE comments (
            id SERIAL PRIMARY KEY,
            blog_id INT REFERENCES blogs(id),
            author_id UUID REFERENCES users(id),
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;

    -- Kiểm tra bảng reviews đã tồn tại chưa
    IF NOT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'reviews'
    ) THEN
        CREATE TABLE reviews (
            id SERIAL PRIMARY KEY,
            inventory_id INT NOT NULL REFERENCES inventories(id),
            user_id UUID NOT NULL REFERENCES users(id),
            rating INT CHECK (rating >= 1 AND rating <= 5),
            comment TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS blogs CASCADE;
-- +goose StatementEnd