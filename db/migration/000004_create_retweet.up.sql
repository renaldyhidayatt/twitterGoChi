CREATE TABLE "retweet" (
  "retweetID" serial PRIMARY KEY,
  "retweetBy" integer NOT NULL,
  "retweetFrom" integer NOT NULL,
  "status" text NOT NULL,
  "tweetOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  CONSTRAINT fk_retweetby
    FOREIGN KEY ("retweetBy") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_retweetfrom
    FOREIGN KEY ("retweetFrom") 
        REFERENCES "tweet"("tweet_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
