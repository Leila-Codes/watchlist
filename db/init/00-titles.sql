-- PostgreSQL Watchlist Tables
-- Titles Table

CREATE TYPE title_type AS ENUM ('movie', 'tv');

CREATE TABLE titles (
    "title_id" BIGSERIAL NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "type" title_type NOT NULL DEFAULT 'movie',
    "description" TEXT NULL,
    "running_time" BIGINT NULL,
    "seasons" INT NULL,
    "episodes" INT NULL,
    "is_verified" BOOLEAN NOT NULL DEFAULT FALSE,
    "submitted_by" BIGINT NULL
);