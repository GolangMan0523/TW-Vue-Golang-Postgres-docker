BEGIN;
CREATE TABLE IF NOT EXISTS "password_resets" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "email" varchar(255) NULL,
    "token" varchar(255) NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
ALTER TABLE "password_resets"
    ADD PRIMARY KEY ("id");
COMMIT;
