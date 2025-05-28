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
INSERT INTO blogs (title, content, author_id) VALUES
('Cách chăm sóc lúa mùa hè', 'Bài viết về cách chăm sóc lúa trong mùa hè để đạt năng suất cao.', (SELECT id FROM users WHERE email = 'admin@example.com')),
('Sử dụng phân bón NPK hiệu quả', 'Hướng dẫn cách sử dụng phân bón NPK để tăng trưởng cây trồng.', (SELECT id FROM users WHERE email = 'farmer@example.com')),
('Kiểm soát sâu bệnh mùa vụ', 'Những biện pháp kiểm soát sâu bệnh hiệu quả cho cây trồng.', (SELECT id FROM users WHERE email = 'farmer@example.com'));
-- Thêm dữ liệu mẫu cho bảng comments
INSERT INTO comments (blog_id, author_id, content) VALUES
(1, (SELECT id FROM users WHERE email = 'customer@example.com'), 'Bài viết rất hữu ích, cảm ơn!'),
(2, (SELECT id FROM users WHERE email = 'admin@example.com'), 'Đúng rồi, phân bón NPK rất quan trọng.'),
(3, (SELECT id FROM users WHERE email = 'farmer@example.com'), 'Mình cũng hay dùng thuốc trừ sâu theo cách này.');

-- Thêm dữ liệu mẫu cho bảng reviews
INSERT INTO reviews (inventory_id, user_id, rating, comment) VALUES
(1, (SELECT id FROM users WHERE email = 'farmer@example.com'), 5, 'Chất lượng tốt, rất hài lòng.'),
(2, (SELECT id FROM users WHERE email = 'farmer@example.com'), 4, 'Hạt giống tốt, tỷ lệ nảy mầm cao.'),
(3, (SELECT id FROM users WHERE email = 'admin@example.com'), 3, 'Cần cải thiện chất lượng thuốc.');
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS blogs CASCADE;
-- +goose StatementEnd