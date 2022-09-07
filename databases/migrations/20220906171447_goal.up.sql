CREATE TABLE IF NOT EXISTS "goal" (
  "id" serial PRIMARY KEY,
  "stream_id" int NOT NULL UNSIGNED,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "target_balance" bigint NOT NULL,
  "progress" varchar(5) NULL DEFAULT ("0"),
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now())
);

ALTER TABLE "goal" ADD FOREIGN KEY (goal_id) REFERENCES "streams" ON DELETE CASCADE;