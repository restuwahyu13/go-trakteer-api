CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "username" varchar(50) NOT NULL UNIQUE,
  "name" varchar(50) NOT NULL UNIQUE,
  "email" varchar(50) NOT NULL UNIQUE,
  "password" varchar(255) NOT NULL,
  "active" boolean NOT NULL,
  "verified" boolean NULL,
  "social_link" jsonb NULL,
  "video_link" text NULL,
  "banner" text NULL,
  "photo" text NULL,
  "role_id" int(50) NOT NULL UNSIGNED,
  "categorie_id" int(50) NOT NULL UNSIGNED,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now()),
  "deleted_at" timestamp NULL
);

ALTER TABLE "users" ADD FOREIGN KEY (role_id) REFERENCES "roles" ON DELETE RESCRICT;
ALTER TABLE "users" ADD FOREIGN KEY (categorie_id) REFERENCES "categories" ON DELETE RESCRICT;