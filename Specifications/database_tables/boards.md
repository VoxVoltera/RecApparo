| name          | type        | constraints                                       | description                     | notes                  |
|---------------|-------------|---------------------------------------------------|---------------------------------|------------------------|
| ID            | UUID        | PRIMARY KEY                                       |                                 |                        |
| NAME          | TEXT        | NOT NULL. CHECK(LENGTH(NAME) \< 1000)             |                                 |                        |
| DESCRIPTION   | TEXT        | CHECK(LENGTH(DESCRIPTION) \< 10000)               |                                 |                        |
| AUTHOR        | UUID        | NOT NULL. REFERENCES users(ID) ON DELETE RESTRICT |                                 |                        |
| CREATION      | TIMESTAMPTZ | NOT NULL. DEFAULT now()                           |                                 |                        |
| MOTHER_TICKET | UUID        | REFERENCES tickets(ID) ON DELETE CASCADE          | Ticket that contains this board | null = top-level board |
