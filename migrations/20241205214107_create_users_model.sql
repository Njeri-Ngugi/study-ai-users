-- Create "users" table
CREATE TABLE "public"."users" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "firstname" text NOT NULL,
  "lastname" text NOT NULL,
  "username" text NULL,
  "password" bytea NOT NULL,
  "email" text NULL,
  "institution" text NOT NULL,
  "year_of_study" bigint NULL DEFAULT 1,
  "completion_date" timestamptz NULL,
  "course_name" text NOT NULL,
  "gender" text NOT NULL,
  "date_of_birth" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
