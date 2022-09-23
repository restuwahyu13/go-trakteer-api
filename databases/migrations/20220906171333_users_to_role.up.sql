CREATE TABLE IF NOT EXISTS "users_to_roles" (
  "id" serial PRIMARY KEY,
  "role_id" int NOT NULL,
  "user_id" int NOT NULL,
  "created_at" timestamp NULL DEFAULT(now())
);

ALTER TABLE "users_to_roles" ADD FOREIGN KEY (role_id) REFERENCES "roles" ON DELETE RESTRICT;
ALTER TABLE "users_to_roles" ADD FOREIGN KEY (user_id) REFERENCES "users" ON DELETE RESTRICT;
ALTER TABLE "users_to_roles" ADD FOREIGN KEY (user_id) REFERENCES "customers" ON DELETE RESTRICT;