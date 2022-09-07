CREATE TABLE IF NOT EXISTS "supporter" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "message" text NOT NULL,
  "is_private" boolean NULL DEFAULT (false),
  "created_at" timestamp NULL DEFAULT (now())
)