BEGIN;
ALTER TABLE IF EXISTS "users"
    ADD COLUMN bio VARCHAR(255) NULL,
    ADD COLUMN location VARCHAR(255) NULL,
    ADD COLUMN website VARCHAR(255) NULL,
    ADD COLUMN birth_date DATE NULL;
COMMIT;

