CREATE TABLE IF NOT EXISTS "assets" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "type" varchar(50) NOT NULL,
  "url" text NOT NULL,
  "created_at" timestamp NULL DEFAULT (now())
)