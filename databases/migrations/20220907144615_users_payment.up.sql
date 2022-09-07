CREATE TABLE IF NOT EXISTS "users_payment" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL UNIQUE,
  "balance_id" int NOT NULL UNIQUE,
  "wallet_id" int NOT NULL UNIQUE
);

ALTER TABLE "users_payment" ADD FOREIGN KEY (user_id) REFERENCES "users" ON DELETE RESTRICT;
ALTER TABLE "users_payment" ADD FOREIGN KEY (balance_id) REFERENCES "balance" ON DELETE RESTRICT;
ALTER TABLE "users_payment" ADD FOREIGN KEY (wallet_id) REFERENCES "wallet" ON DELETE RESTRICT;