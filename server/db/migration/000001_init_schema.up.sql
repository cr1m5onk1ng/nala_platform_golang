CREATE TABLE "resources" (
  "id" bigserial PRIMARY KEY,
  "url" varchar UNIQUE NOT NULL,
  "language" varchar(2) NOT NULL,
  "difficulty" varchar NOT NULL,
  "media_type" varchar NOT NULL,
  "category" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "registration_date" timestamptz NOT NULL DEFAULT 'now()',
  "native_language" varchar(2) NOT NULL,
  "role" varchar
);

CREATE TABLE "learning" (
  "user_id" varchar,
  "language" varchar(2),
  PRIMARY KEY ("user_id", "language")
);

CREATE TABLE "community" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "language" varchar(2) NOT NULL,
  "title" varchar NOT NULL,
  "thumbnail_url" varchar NOT NULL,
  "metadata" jsonb
);

CREATE TABLE "community_users" (
  "community_id" bigserial NOT NULL,
  "user_id" varchar NOT NULL,
  PRIMARY KEY ("community_id", "user_id")
);

CREATE TABLE "user_post" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "resource_id" bigint UNIQUE NOT NULL,
  "post_time" timestamptz NOT NULL DEFAULT 'now()',
  "post_title" varchar NOT NULL,
  "post_description" varchar
);

CREATE TABLE "tags" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "tag" varchar NOT NULL
);

CREATE TABLE "post_tags" (
  "tag_id" bigint NOT NULL,
  "post_id" varchar NOT NULL,
  PRIMARY KEY ("tag_id", "post_id")
);

CREATE TABLE "topics" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "topic" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "post_topics" (
  "post_id" varchar NOT NULL,
  "topic_id" bigint NOT NULL,
  PRIMARY KEY ("post_id", "topic_id")
);

CREATE TABLE "votes" (
  "user_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  "difficulty" varchar NOT NULL,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "likes" (
  "user_id" varchar NOT NULL,
  "post_id" varchar NOT NULL,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "user_is_followed" (
  "followed_user_id" varchar NOT NULL,
  "follower_user_id" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  PRIMARY KEY ("followed_user_id", "follower_user_id")
);

CREATE TABLE "user_follows" (
  "follower_user_id" varchar NOT NULL,
  "followed_user_id" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  PRIMARY KEY ("follower_user_id", "followed_user_id")
);

CREATE TABLE "post_discussions" (
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
  "parent_comment_id" bigint NOT NULL DEFAULT '-1',
  "user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "content" varchar NOT NULL
);

CREATE TABLE "comments_likes" (
  "comment_id" bigint NOT NULL,
  "user_id" varchar NOT NULL,
  PRIMARY KEY ("comment_id", "user_id")
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

ALTER TABLE "learning" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "community_users" ADD FOREIGN KEY ("community_id") REFERENCES "community" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "community_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_post" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_post" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_tags" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_topics" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_topics" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "votes" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_is_followed" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_is_followed" ADD FOREIGN KEY ("follower_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_follows" ADD FOREIGN KEY ("follower_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_discussions" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "post_discussions" ADD FOREIGN KEY ("post_id") REFERENCES "user_post" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "discussion_comments" ADD FOREIGN KEY ("discussion_id") REFERENCES "post_discussions" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "discussion_comments" ADD FOREIGN KEY ("parent_comment_id") REFERENCES "discussion_comments" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "discussion_comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "comments_likes" ADD FOREIGN KEY ("comment_id") REFERENCES "discussion_comments" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "comments_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "study_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "study_list_resource" ADD FOREIGN KEY ("study_list_id") REFERENCES "study_lists" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "study_list_resource" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

CREATE INDEX ON "community" ("language");

CREATE INDEX ON "user_post" ("user_id");

CREATE INDEX ON "user_post" ("post_title");

CREATE INDEX ON "tags" ("tag");

CREATE INDEX ON "post_discussions" ("post_id");

CREATE INDEX ON "post_discussions" ("creator_id");

CREATE INDEX ON "discussion_comments" ("user_id");

CREATE INDEX ON "discussion_comments" ("discussion_id");

CREATE INDEX ON "study_lists" ("user_id");

COMMENT ON COLUMN "resources"."language" IS '2 chars language code';

COMMENT ON COLUMN "users"."id" IS 'use UUID';

COMMENT ON COLUMN "users"."native_language" IS '2 chars language code';

COMMENT ON COLUMN "user_post"."id" IS 'use UUID';