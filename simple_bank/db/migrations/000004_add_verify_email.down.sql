DROP TABLE IF EXISTS "verify_email" CASCADE;
ALTER TABLE "users" DROP COLUMN IF EXISTS "is_email_verified";
