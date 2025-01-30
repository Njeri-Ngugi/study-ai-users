-- Create "courses" table
CREATE TABLE "public"."courses" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "course_name" text NULL,
  "course_code" text NULL,
  "institution" text NOT NULL,
  "course_duration" bigint NOT NULL DEFAULT 4,
  PRIMARY KEY ("id")
);
-- Create index "idx_courses_deleted_at" to table: "courses"
CREATE INDEX "idx_courses_deleted_at" ON "public"."courses" ("deleted_at");
