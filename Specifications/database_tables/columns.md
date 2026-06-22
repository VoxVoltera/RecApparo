| name     | type    | constraints                                       | description | notes                             |
|----------|---------|---------------------------------------------------|-------------|-----------------------------------|
| ID       | UUID    | PRIMARY KEY                                       |             |                                   |
| BOARD_ID | UUID    | NOT NULL. REFERENCES boards(ID) ON DELETE CASCADE |             |                                   |
| NAME     | TEXT    | NOT NULL. CHECK(LENGTH(NAME) \< 1000)             |             |                                   |
| POSITION | INTEGER | NOT NULL. CHECK(POSITION \>= 0)                   |             | 0-based ordering within the board |
