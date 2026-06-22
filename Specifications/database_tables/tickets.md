| name        | type        | constraints                                         | description                | notes                                             |
|-------------|-------------|-----------------------------------------------------|----------------------------|---------------------------------------------------|
| ID          | UUID        | PRIMARY KEY                                         |                            |                                                   |
| NAME        | TEXT        | NOT NULL. CHECK(LENGTH(NAME) \< 1000)               |                            |                                                   |
| DESCRIPTION | TEXT        | CHECK(LENGTH(DESCRIPTION) \< 10000)                 |                            |                                                   |
| AUTHOR      | UUID        | NOT NULL. REFERENCES users(ID) ON DELETE RESTRICT   |                            |                                                   |
| CREATION    | TIMESTAMPTZ | NOT NULL. DEFAULT now()                             |                            |                                                   |
| DUE_DATE    | TIMESTAMPTZ |                                                     |                            |                                                   |
| BOARD_ID    | UUID        | NOT NULL. REFERENCES boards(ID) ON DELETE CASCADE   | Board the ticket lives on  |                                                   |
| COLUMN_ID   | UUID        | NOT NULL. REFERENCES columns(ID) ON DELETE RESTRICT |                            | Added: was missing from the struct                |
| SUB_BOARD   | UUID        | REFERENCES boards(ID) ON DELETE SET NULL            | Board this ticket contains | Add UNIQUE if one board has a single owner ticket |
