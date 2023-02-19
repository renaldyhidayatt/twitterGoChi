CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "firstName" varchar(100) NOT NULL,
  "lastName" varchar(100) NOT NULL,
  "username" varchar(150) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "profileImage" varchar(255) NOT NULL,
  "profileCover" varchar(255) NOT NULL,
  "following" integer NOT NULL,
  "followers" integer NOT NULL,
  "bio" text NOT NULL,
  "country" varchar(255) NOT NULL,
  "website" varchar(255) NOT NULL
);
