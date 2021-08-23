CREATE TABLE "resources" (
  "id" bigserial PRIMARY KEY,
  "url" varchar UNIQUE NOT NULL,
  "language" char(2) NOT NULL,
  "difficulty" varchar,
  "title" varchar NOT NULL,
  "description" varchar,
  "media_type" varchar,
  "category" varchar NOT NULL,
  "thumbnail_url" varchar,
  "inserted_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "registration_date" timestamptz NOT NULL DEFAULT 'now()',
  "native_language" char(2) NOT NULL,
  "role" varchar
);

CREATE TABLE "learning" (
  "user_id" varchar NOT NULL,
  "language" char(2),
  PRIMARY KEY ("user_id", "language")
);

CREATE TABLE "user_post" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "resource_id" bigint UNIQUE NOT NULL,
  "post_time" timestamptz NOT NULL DEFAULT 'now()',
  "post_title" varchar NOT NULL,
  "post_description" varchar
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  "content" varchar NOT NULL,
  "comment_time" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "comments_likes" (
  "comment_id" bigint NOT NULL,
  "user_id" varchar NOT NULL,
  PRIMARY KEY ("comment_id", "user_id")
);

CREATE TABLE "comments_responses" (
  "source_comment_id" bigint NOT NULL,
  "response_comment_id" bigint NOT NULL,
  PRIMARY KEY ("source_comment_id", "response_comment_id")
);

CREATE TABLE "tags" (
  "post_id" varchar NOT NULL,
  "tag" varchar NOT NULL,
  PRIMARY KEY ("post_id", "tag")
);

CREATE TABLE "votes" (
  "user_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  "difficulty" varchar NOT NULL,
  "comment" varchar,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "likes" (
  "user_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "followers" (
  "follower_id" varchar NOT NULL,
  "followed_id" varchar NOT NULL,
  "creation_time" timestamptz NOT NULL DEFAULT 'now()',
  PRIMARY KEY ("follower_id", "followed_id")
);

CREATE TABLE "resource_discussions" (
  "id" bigserial PRIMARY KEY,
  "creator_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  "creation_time" timestamptz NOT NULL DEFAULT 'now()',
  "title" varchar NOT NULL,
  "description" varchar
);

CREATE TABLE "discussion_comments" (
  "id" bigserial PRIMARY KEY,
  "discussion_id" bigint NOT NULL,
  "user_id" varchar NOT NULL,
  "creation_time" timestamptz NOT NULL DEFAULT 'now()',
  "content" varchar
);

CREATE TABLE "study_lists" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "creation_time" timestamptz NOT NULL DEFAULT 'now()',
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "public" boolean NOT NULL DEFAULT 'true'
);

CREATE TABLE "study_list_resource" (
  "study_list_id" bigint NOT NULL,
  "resource_id" bigint NOT NULL,
  "time_added" timestamptz NOT NULL DEFAULT 'now()',
  PRIMARY KEY ("study_list_id", "resource_id")
);

ALTER TABLE "learning" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_post" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_post" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id");

ALTER TABLE "comments_likes" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id");

ALTER TABLE "comments_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments_responses" ADD FOREIGN KEY ("source_comment_id") REFERENCES "comments" ("id");

ALTER TABLE "comments_responses" ADD FOREIGN KEY ("response_comment_id") REFERENCES "comments" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("followed_id") REFERENCES "users" ("id");

ALTER TABLE "resource_discussions" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "resource_discussions" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id");

ALTER TABLE "discussion_comments" ADD FOREIGN KEY ("discussion_id") REFERENCES "resource_discussions" ("id");

ALTER TABLE "discussion_comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "study_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "study_list_resource" ADD FOREIGN KEY ("study_list_id") REFERENCES "study_lists" ("id");

ALTER TABLE "study_list_resource" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

CREATE INDEX ON "user_post" ("user_id", "resource_id");

CREATE INDEX ON "comments" ("post_id", "user_id");

CREATE INDEX ON "resource_discussions" ("post_id");

CREATE INDEX ON "resource_discussions" ("creator_id");

CREATE INDEX ON "discussion_comments" ("user_id");

CREATE INDEX ON "discussion_comments" ("discussion_id");

CREATE INDEX ON "study_lists" ("user_id");

COMMENT ON COLUMN "resources"."language" IS '2 chars language code';

COMMENT ON COLUMN "users"."id" IS 'use UUID';

COMMENT ON COLUMN "users"."native_language" IS '2 chars language code';

COMMENT ON COLUMN "user_post"."id" IS 'use UUID';