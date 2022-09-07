CREATE TABLE IF NOT EXISTS "payments" (
  "id" serial PRIMARY KEY,
  "transaction_id" varchar(255) NULL,
  "external_id" varchar(255) NULL,
  "status" varchar(50) NOT NULL,
  "amount" bigint NOT NULL,
  "is_paid" boolean NULL DEFAULT (false),
  "type" varchar(50) NOT NULL,
  "goal_id" int NULL,
  "supporter_id" int NULL,
  "created_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "payments" ADD FOREIGN KEY (goal_id) REFERENCES "goal" ON DELETE CASCADE;
ALTER TABLE "payments" ADD FOREIGN KEY (supporter_id) REFERENCES "supporter" ON DELETE CASCADE;