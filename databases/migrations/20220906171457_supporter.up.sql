CREATE TABLE IF NOT EXISTS "supporter" (
  "id" serial PRIMARY KEY
  "name" int UNSIGNED NOT NULL
  "message" date NOT NULL
  "is_private" date NOT NULL
  "created_at" timestamp NULL DEFAULT (now())
)