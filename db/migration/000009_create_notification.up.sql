CREATE TYPE enum_notification_type AS ENUM ('like','comment','retweet','follow','message','mention');

CREATE TABLE "notification" (
  "ID" serial NOT NULL PRIMARY KEY,
  "notificationFor" integer NOT NULL,
  "notificationFrom" integer NOT NULL,
  "target" integer NOT NULL,
  "type" enum_notification_type NOT NULL,
  "notificationOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  "notificationCount" integer NOT NULL,
  "status" integer NOT NULL
);
