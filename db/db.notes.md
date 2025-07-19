# DB Notes

## Requirements
1. CRUD of "Movies and TV Shows"
    - A complete collection of known movies/tv shows in the system.
    - Ability for users to "submit" new ones.
    - Ability to add a box art photo of each item.
    - "Verified" and "SubmittedBy" fields.
    - Verified is shown to ALL users, "SubmittedBy" only shown to the author (and subsequent lists) 

2. CRUD of "Watchlist" for Users.
    - Custom ranking of watchlist - i.e. a "Next Up" Queue.
    - Mark items as "watched", "in progress", "not yet watched"

## Tables
1. `titles`

| Column | Type | Description | Nullable |
| ------ | ---- | ----------- | -------- |
| title_id | BIGSERIAL | - | false |
| title  | TEXT | - | false |
| type | SMALLINT | Enum `0` = Movie, `1` = TV | false |
| description | TEXT |  | true |
| running_time | BIGINT | - | true |
| seasons | INT | - | true |
| episodes | INT | - | true |
| is_verified | BOOLEAN | - | false |
| submitted_by | BIGINT | - | true |

2. `watchlists`

| Column | Type | Description | Nullable |
| ------ | ---- | ----------- | -------- |
| list_id | VARCHAR(255) | Unique! | false |
| owner_id | INT | - | false |

3. `watchlist_users`

| Column | Type | Description | Nullable |
| ------ | ---- | ----------- | -------- |
| list_id | VARCHAR(255) | FK (`watchlists.list_id`) | false |
| user_id | BIGINT | - | false

4. `watchlist_content`

| Column | Type | Description | Nullable |
| ------ | ---- | ----------- | -------- |
| list_id | VARCHAR(255) | FK (`watchlists.list_id`) | false |
| title_id | BIGINT | FK (`titles.title_id`) | false |
| position | INT | User defined ordering number. | true |
| isWatched | BOOLEAN | Whether this show has been watched. | false