CREATE TABLE IF NOT EXISTS "users_payment" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL UNIQUE UNSIGNED,
  "balance_id" int NOT NULL UNIQUE UNSIGNED,
  "wallet_id" int NOT NULL UNIQUE UNSIGNED
)

ALTER TABLE "users_payment" ADD FOREIGN KEY user_id REFERENCES "users" ON DELETE RESCRICT;
ALTER TABLE "users_payment" ADD FOREIGN KEY balance_id REFERENCES "balance" ON DELETE RESCRICT;
ALTER TABLE "users_payment" ADD FOREIGN KEY wallet_id REFERENCES "wallet" ON DELETE RESCRICT;