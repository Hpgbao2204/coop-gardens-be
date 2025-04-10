-- +goose Up
-- +goose StatementBegin
SELECT COUNT(*) FROM crops;
SELECT COUNT(*) FROM seasons;
SELECT r.name AS role, COUNT(*) 
  FROM user_roles ur JOIN roles r ON ur.role_id = r.id 
  GROUP BY r.name;
SELECT COALESCE(AVG(rating), 0) FROM reviews;
SELECT COUNT(*) FROM orders;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
