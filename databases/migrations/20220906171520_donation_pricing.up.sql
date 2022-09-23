CREATE TABLE IF NOT EXISTS "donation_pricing" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL UNIQUE,
  "donation_type_id" int NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "donation_pricing" ADD FOREIGN KEY (user_id) REFERENCES "customers" ON DELETE CASCADE;
ALTER TABLE "donation_pricing" ADD FOREIGN KEY (donation_type_id) REFERENCES "donation_type" ON DELETE CASCADE;