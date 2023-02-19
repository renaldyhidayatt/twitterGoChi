
CREATE TABLE "likes" (
  "likeID" serial PRIMARY KEY,
  "likeOn" integer NOT NULL,
  "likeBy" integer NOT NULL,
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