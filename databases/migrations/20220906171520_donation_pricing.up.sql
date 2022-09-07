CREATE TABLE IF NOT EXISTS "donation_pricing" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL UNSIGNED UNIQUE,
  "donation_type_id" int NOT NULL UNSIGNED,
  "amount" bigint NOT NULL UNSIGNED,
  "created_at" NULL DEFAULT (now()),
  "updated_at" NULL DEFAULT (now())
);

ALTER TABLE "donation_pricing" ADD FOREIGN KEY user_id REFERENCES "users" ON DELETE CASCADE;
ALTER TABLE "donation_pricing" ADD FOREIGN KEY donation_type_id REFERENCES "donation_type" ON DELETE CASCADE;