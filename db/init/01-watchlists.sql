-- PostgreSQL Watchlist Tables
-- Watchlist Tables

-- watchlists table
-- defines watchlists created by users and who owns them.
CREATE TABLE watchlists (
    "list_id" VARCHAR(255) NOT NULL PRIMARY KEY,
    "owner_id" BIGINT NOT NULL
);

-- watchlist_users table
-- defines which "other" users have access to a watchlist. (enabling others to edit a watchlist)
CREATE TABLE watchlist_users (
    "list_id" VARCHAR(255) NOT NULL,
    "user_id" BIGINT NOT NULL,
    CONSTRAINT fk_watchlist_list FOREIGN KEY (list_id) REFERENCES watchlists(list_id)
);

-- watchlist_content
-- defines contents of watchlist, i.e. which titles are in the list,
-- their position and whether or not a title has been watched.
CREATE TABLE watchlist_content (
    "list_id" VARCHAR(255) NOT NULL,
    "title_id" BIGINT NOT NULL,
    "position" INT NOT NULL,
    "is_watched" BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_watchlist_list FOREIGN KEY (list_id) REFERENCES watchlists(list_id),
    CONSTRAINT fk_watchlist_title FOREIGN KEY (title_id) REFERENCES titles(title_id)
);