CREATE TABLE IF NOT EXISTS "withdraw" (
  "id" int PRIMARY KEY,
  "balance_id" int NOT NULL,
  "amount" bigint NOT NULL,
  "date_time" timestamp NOT NULL,
  "fee" int NOT NULL,
  "created_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "withdraw" ADD FOREIGN KEY (balance_id) REFERENCES "balance" ON DELETE CASCADE;