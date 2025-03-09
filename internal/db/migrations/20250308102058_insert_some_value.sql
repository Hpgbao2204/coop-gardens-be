-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, email, password) VALUES 
(2, 'user2@example.com', 'password'),
(3, 'user3@example.com', 'password'),
(4, 'user4@example.com', 'password'),
(5, 'user5@example.com', 'password'),
(6, 'user6@example.com', 'password');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email IN (
    'user2@example.com',
    'user3@example.com',
    'user4@example.com',
    'user5@example.com',
    'user6@example.com'
);
-- +goose StatementEnd
