CREATE TABLE "messages" (
  "messageID" serial NOT NULL PRIMARY KEY,
  "message" text NOT NULL,
  "messageTo" integer NOT NULL,
  "messageFrom" integer NOT NULL,
  "messageOn" timestamp with time zone NOT NULL DEFAULT current_timestamp,
  "status" integer NOT NULL,
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