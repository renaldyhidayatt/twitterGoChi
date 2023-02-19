CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "firstName" varchar(100) NOT NULL,
  "lastName" varchar(100) NOT NULL,
  "username" varchar(150) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "profileImage" varchar(255) NOT NULL,
  "profileCover" varchar(255) NOT NULL,
  "following" int(11) NOT NULL,
  "followers" int(11) NOT NULL,
  "bio" text NOT NULL,
  "country" varchar(255) NOT NULL,
  "website" varchar(255) NOT NULL,
  CONSTRAINT "fk_following" FOREIGN KEY ("following") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "fk_followers" FOREIGN KEY ("followers") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE "tweet" (
  "tweet_id" serial NOT NULL PRIMARY KEY,
  "status" text NOT NULL,
  "tweetBy" int(11) NOT NULL,
  "tweetImage" text NOT NULL,
  "postedOn" datetime NOT NULL DEFAULT current_timestamp(),
  CONSTRAINT fk_tweetby
    FOREIGN KEY ("tweetBy") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE "retweet" (
  "retweetID" serial PRIMARY KEY,
  "retweetBy" int(11) NOT NULL,
  "retweetFrom" int(11) NOT NULL,
  "status" text NOT NULL,
  "tweetOn" datetime NOT NULL DEFAULT current_timestamp(),
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


CREATE TABLE "likes" (
  "likeID" serial PRIMARY KEY,
  "likeOn" int(11) NOT NULL,
  "likeBy" int(11) NOT NULL,
  CONSTRAINT fk_likeon
    FOREIGN KEY ("likeOn") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_likeby
    FOREIGN KEY ("likeBy") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE "follow" (
  "followID" serial PRIMARY KEY NOT NULL,
  "sender" int(11) NOT NULL,
  "receiver" int(11) NOT NULL,
  "followStatus" enum('0','1') NOT NULL,
  "followOn" datetime NOT NULL DEFAULT current_timestamp(),
  CONSTRAINT "fk_sender" FOREIGN KEY ("sender") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "fk_receiver" FOREIGN KEY ("receiver") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "comment" (
  "commentID" serial NOT NULL PRIMARY KEY,
  "commentBy" int(11) NOT NULL,
  "commentOn" int(11) NOT NULL,
  "comment" text NOT NULL,
  "commentAt" datetime NOT NULL DEFAULT current_timestamp(),
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

CREATE TABLE "messages" (
  "messageID" serial NOT NULL PRIMARY KEY,
  "message" text NOT NULL,
  "messageTo" int(11) NOT NULL,
  "messageFrom" int(11) NOT NULL,
  "messageOn" datetime NOT NULL DEFAULT current_timestamp(),
  "status" int(11) NOT NULL,
  CONSTRAINT fk_messageto
    FOREIGN KEY ("messageTo") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_messagefrom
    FOREIGN KEY ("messageFrom") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE "notification" (
  "ID" serial NOT NULL PRIMARY KEY,
  "notificationFor" int(11) NOT NULL,
  "notificationFrom" int(11) NOT NULL,
  "target" int(11) NOT NULL,
  "type" enum('like','comment','retweet','follow','message','mention') NOT NULL,
  "notificationOn" datetime NOT NULL DEFAULT current_timestamp(),
  "notificationCount" int(11) NOT NULL,
  "status" int(11) NOT NULL
);


CREATE TABLE "trends" (
  "trendID" serial NOT NULL PRIMARY KEY,
  "hashtag" varchar(200) NOT NULL,
  "user_id" int(11) NOT NULL,
  "tweetId" int(11) NOT NULL,
  "createdOn" datetime NOT NULL DEFAULT current_timestamp(),
  CONSTRAINT fk_user
    FOREIGN KEY (user_id) 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_tweet
    FOREIGN KEY (tweetId) 
        REFERENCES "tweet"("tweet_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);




CREATE TABLE "verification" (
  "verify_id" serial NOT NULL PRIMARY KEY,
  "user_id" int(11) NOT NULL,
  "code" varchar(255) NOT NULL,
  "status" enum('0','1') NOT NULL,
  "createdAt" datetime NOT NULL DEFAULT current_timestamp(),
  CONSTRAINT fk_user
    FOREIGN KEY ("user_id") 
        REFERENCES "users"("user_id")
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

