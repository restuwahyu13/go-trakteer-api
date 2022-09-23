CREATE TABLE IF NOT EXISTS customers (
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
  "role_id" int NOT NULL,
  "categorie_id" int NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now()),
  "deleted_at" timestamp NULL
);

ALTER TABLE "customers" ADD FOREIGN KEY (role_id) REFERENCES "roles" ON DELETE RESTRICT;
ALTER TABLE "customers" ADD FOREIGN KEY (categorie_id) REFERENCES "categories" ON DELETE RESTRICT;