CREATE TABLE "comment" (
  "commentID" serial NOT NULL PRIMARY KEY,
  "commentBy" integer NOT NULL,
  "commentOn" integer NOT NULL,
  "comment" text NOT NULL,
  "commentAt" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  CONSTRAINT fk_user
    FOREIGN KEY ("commentBy") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_tweet
    FOREIGN KEY ("commentOn") 
        REFERENCES "tweet"("tweet_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);