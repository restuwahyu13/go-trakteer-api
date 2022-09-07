CREATE TABLE IF NOT EXISTS "donation_type" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "icon" text NOT NULL,
  "created_at" NULL DEFAULT (now()),
  "updated_at" NULL DEFAULT (now())
)