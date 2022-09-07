CREATE TABLE IF NOT EXISTS "activities" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "user_info" jsonb NOT NULL,
  "first_login" timestamp NOT NULL,
  "last_login" timestamp NOT NULL,
  "created_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "activities" ADD FOREIGN KEY (user_id) REFERENCES "users" ON DELETE CASCADE;