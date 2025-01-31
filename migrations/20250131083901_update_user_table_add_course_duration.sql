-- Rename a column from "course_duration" to "course_duration_in_years"
ALTER TABLE "public"."courses" RENAME COLUMN "course_duration" TO "course_duration_in_years";
-- Modify "courses" table
ALTER TABLE "public"."courses" ALTER COLUMN "course_name" SET NOT NULL, ALTER COLUMN "course_code" SET NOT NULL;
-- Create index "idx_courses_course_name" to table: "courses"
CREATE INDEX "idx_courses_course_name" ON "public"."courses" ("course_name");
-- Create index "idx_courses_institution_id" to table: "courses"
CREATE INDEX "idx_courses_institution_id" ON "public"."courses" ("institution_id");
-- Create index "idx_institutions_institution_name" to table: "institutions"
CREATE INDEX "idx_institutions_institution_name" ON "public"."institutions" ("institution_name");
-- Modify "units" table
ALTER TABLE "public"."units" ALTER COLUMN "unit_name" SET NOT NULL, ALTER COLUMN "course_id" SET NOT NULL;
-- Create index "idx_units_course_id" to table: "units"
CREATE INDEX "idx_units_course_id" ON "public"."units" ("course_id");
-- Create index "idx_units_unit_name" to table: "units"
CREATE INDEX "idx_units_unit_name" ON "public"."units" ("unit_name");
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "username" SET NOT NULL, ADD COLUMN "course_duration_in_years" bigint NOT NULL DEFAULT 4;
-- Create index "idx_users_institution_id" to table: "users"
CREATE INDEX "idx_users_institution_id" ON "public"."users" ("institution_id");
-- Create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "public"."users" ("username");
