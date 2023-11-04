CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY NOT NULL ,
    firstname   VARCHAR(255)          NOT NULL ,
    middlename VARCHAR(255)   ,
    lastname    VARCHAR(255)          NOT NULL ,
    email        TEXT       UNIQUE   NOT NULL ,
    password     TEXT          NOT NULL ,
    created_at  TIMESTAMP             NOT NULL ,
    updated_at  TIMESTAMP             NOT NULL,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    otp         VARCHAR(255)          NOT NULL
);


CREATE INDEX email_index ON users (email);