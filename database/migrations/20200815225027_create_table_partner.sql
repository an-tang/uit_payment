-- +goose Up
-- +goose StatementBegin
CREATE TABLE "partners" (
  "id" serial primary key,
  "name" text,
  "key" text,
  "callback_url" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "payments" 
ADD COLUMN "partner_id" integer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE INDEX "partners_id_idx" ON  "partners" ("id");
CREATE INDEX "partners_key_idx" ON  "partners" ("key");
DROP TABLE "partners";

ALTER TABLE "payments"
REMOVE COLUMN "partner_id";
-- +goose StatementEnd
