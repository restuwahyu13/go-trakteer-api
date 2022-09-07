CREATE TABLE IF NOT EXISTS "withdraw" (
  "id" int PRIMARY KEY
  "balance_id" int NOT NULL UNSIGNED
  "amount" bigint NOT NULL UNSIGNED
  "date_time" datetime NOT NULL
  "fee" int NOT NULL UNSIGNED
  "created_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "withdraw" ADD FOREIGN KEY balance_id REFERENCES "balance" ON DELETE CASCADE;