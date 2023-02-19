CREATE TABLE "tweet" (
  "tweet_id" serial PRIMARY KEY,
  "status" text NOT NULL,
  "tweetBy" integer NOT NULL,
  "tweetImage" text NOT NULL,
  "postedOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  CONSTRAINT "fk_tweetBy" FOREIGN KEY ("tweetBy") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE
);
