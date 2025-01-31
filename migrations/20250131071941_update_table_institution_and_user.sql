-- Modify "institutions" table
ALTER TABLE "public"."institutions" ALTER COLUMN "institution_name" SET NOT NULL, DROP COLUMN "country_code", ADD COLUMN "country_of_study" text NOT NULL DEFAULT 'KE';
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "country_of_study" text NOT NULL DEFAULT 'KE';
