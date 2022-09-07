CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY
  "username" varchar(25) NOT NULL UNIQUE,
  "email" varchar(25) NOT NULL UNIQUE,
  "password" varchar(255) NOT NULL,
  "role_id" int(25) NOT NULL UNSIGNED,
  "categorie_id" int(25) NOT NULL UNSIGNED
);

ALTER TABLE "users" ADD FOREIGN KEY (role_id) REFERENCES "roles" ON DELETE CASCADE;
ALTER TABLE "users" ADD FOREIGN KEY (categorie_id) REFERENCES "categories" ON DELETE CASCADE;