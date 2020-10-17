-- +goose Up
-- +goose StatementBegin
CREATE TABLE "payments" (
  "id" serial primary key,
  "uid" varchar(40) not null,
  "transaction_id" text not null,
  "payment_method" integer not null,
  "amount" decimal not null,
  "currency" varchar(3),
  "store_id" varchar(10),
  "qr_code" text,
  "payment_tx" text,
  "status" integer default 1,
  "paid_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp
);
-- +goose StatementEnd
CREATE UNIQUE INDEX "payments_uid_unique_indx" ON "payments" ("uid");
CREATE INDEX "payment_transaction_id_idx" ON "payments" ("transaction_id");
CREATE INDEX "payment_payment_method_idx" ON "payments" ("payment_method");
CREATE INDEX "payments_payment_tx_idx" ON  "payments" ("payment_tx");

-- +goose Down
-- +goose StatementBegin
DROP TABLE "payments";
-- +goose StatementEnd
