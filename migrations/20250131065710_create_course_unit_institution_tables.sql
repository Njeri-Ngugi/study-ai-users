-- Rename a column from "institution" to "institution_id"
ALTER TABLE "public"."courses" RENAME COLUMN "institution" TO "institution_id";
-- Rename a column from "institution" to "institution_id"
ALTER TABLE "public"."users" RENAME COLUMN "institution" TO "institution_id";
-- Rename a column from "course_name" to "course_id"
ALTER TABLE "public"."users" RENAME COLUMN "course_name" TO "course_id";
-- Create "institutions" table
CREATE TABLE "public"."institutions" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "institution_name" text NULL,
  "country_code" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_institutions_deleted_at" to table: "institutions"
CREATE INDEX "idx_institutions_deleted_at" ON "public"."institutions" ("deleted_at");
-- Create "units" table
CREATE TABLE "public"."units" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "unit_name" text NULL,
  "course_id" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_units_deleted_at" to table: "units"
CREATE INDEX "idx_units_deleted_at" ON "public"."units" ("deleted_at");
