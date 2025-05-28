-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  price DOUBLE PRECISION NOT NULL,
  stock INTEGER NOT NULL,
  farmer_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_products_farmer FOREIGN KEY (farmer_id) REFERENCES users(id) ON DELETE CASCADE

);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  total DOUBLE PRECISION,
  status TEXT DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE

);

CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  price DOUBLE PRECISION,
  CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- Thêm dữ liệu mẫu cho bảng products
INSERT INTO products (name, description, price, stock, farmer_id) VALUES
('Gạo thơm ST25', 'Gạo dẻo, thơm, chất lượng cao', 15.0, 1000, (SELECT id FROM users WHERE email = 'farmer@example.com')),
('Khoai tây sạch', 'Khoai tây tươi, chất lượng cao', 10.0, 500, (SELECT id FROM users WHERE email = 'farmer@example.com')),
('Rau cải xanh', 'Rau cải xanh sạch, trồng hữu cơ', 5.0, 300, (SELECT id FROM users WHERE email = 'farmer@example.com'));

-- Thêm dữ liệu mẫu cho bảng orders
INSERT INTO orders (user_id, total, status) VALUES
((SELECT id FROM users WHERE email = 'customer@example.com'), 50.0, 'completed'),
((SELECT id FROM users WHERE email = 'customer@example.com'), 20.0, 'pending'),
((SELECT id FROM users WHERE email = 'farmer@example.com'), 30.0, 'completed');

-- Thêm dữ liệu mẫu cho bảng order_items
INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
(1, 1, 2, 30.0),
(1, 2, 1, 10.0),
(2, 3, 4, 20.0),
(3, 1, 1, 15.0),
(3, 2, 1, 10.0);
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
