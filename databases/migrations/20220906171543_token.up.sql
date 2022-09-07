CREATE TABLE IF EXISTS "token" (
  "id" serial PRIMARY KEY,
  "resource_id" varchar(255) NULL,
  "resource_type" varchar(255) NULL,
  "access_token" text NOT NULL,
  "refresh_token" text NOT NULL,
  "expired_at" datetime NULL,
  "created_at" timestamp NULL DEFAULT (now())
)