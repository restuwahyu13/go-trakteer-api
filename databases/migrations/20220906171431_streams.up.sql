CREATE TABLE IF NOT EXISTS "streams" (
  "id" serial PRIMARY KEY,
  "notification" jsonb NULL,
  "leaderboard" jsonb NULL,
  "last_supporter" jsonb NULL,
  "target" jsonb NULL,
  "running_text" jsonb NULL,
  "subathon" jsonb NULL,
  "qrcode" jsonb NULL,
  "key" varchar(255) NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now())
)