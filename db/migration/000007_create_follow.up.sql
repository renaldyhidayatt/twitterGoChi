CREATE TABLE "follow" (
"followID" serial PRIMARY KEY,
"sender" integer NOT NULL,
"receiver" integer NOT NULL,
"followStatus" varchar(1) NOT NULL CHECK ("followStatus" IN ('0', '1')),
"followOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
CONSTRAINT "fk_sender" FOREIGN KEY ("sender") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE,
CONSTRAINT "fk_receiver" FOREIGN KEY ("receiver") REFERENCES "users" ("user_id") ON DELETE CASCADE ON UPDATE CASCADE
);