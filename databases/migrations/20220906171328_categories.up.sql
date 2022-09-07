CREATE TABLE IF NOT EXISTS "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar(25) UNIQUE NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now()),
  "deleted_at" timestamp NULL
  )