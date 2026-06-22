| name        | type    | constraints                                      | description | notes             |
|-------------|---------|--------------------------------------------------|-------------|-------------------|
| ID          | UUID    | PRIMARY KEY                                      |             |                   |
| NAME        | TEXT    | NOT NULL. CHECK(LENGTH(NAME) \< 1000)            |             |                   |
| DESCRIPTION | TEXT    | CHECK(LENGTH(DESCRIPTION) \< 10000)              |             |                   |
| COLOUR      | TEXT    | NOT NULL. CHECK(COLOUR ~ '^#\[0-9A-Fa-f\]{6}\$') |             | e.g. \#1D9E75     |
| SPECIAL     | BOOLEAN | NOT NULL. DEFAULT false                          |             |                   |
| BOARD_ID    | UUID    | REFERENCES boards(ID) ON DELETE CASCADE          |             | null = global tag |
