
CREATE TABLE "trends" (
  "trendID" serial NOT NULL PRIMARY KEY,
  "hashtag" varchar(200) NOT NULL,
  "user_id" integer NOT NULL,
  "tweetId" integer NOT NULL,
  "createdOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "fk_tweet" FOREIGN KEY ("tweetId") REFERENCES "tweet" ("tweet_id") ON DELETE CASCADE ON UPDATE CASCADE
);


