-- +goose Up
-- +goose StatementBegin
CREATE TABLE "ps_pets" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  tags VARCHAR(255) []
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "ps_pets";
-- +goose StatementEnd
