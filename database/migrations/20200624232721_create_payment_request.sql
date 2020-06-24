-- +goose Up
-- +goose StatementBegin
CREATE TABLE "payment_requests" (
  "id" serial primary key,
  "req" jsonb,
  "resp" jsonb,
  "status" integer default 200,
  "request_type" integer default 1,
  "payment_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE INDEX "payment_requests_payment_id_idx" ON  "payment_requests" ("payment_id");
CREATE INDEX "payment_requests_request_type_idx" ON  "payment_requests" ("request_type");
DROP TABLE "payment_requests";
-- +goose StatementEnd
