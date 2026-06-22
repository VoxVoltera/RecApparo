| name      | type        | constraints                                        | description                         | notes                    |
|-----------|-------------|----------------------------------------------------|-------------------------------------|--------------------------|
| ID        | UUID        | PRIMARY KEY                                        |                                     |                          |
| AUTHOR    | UUID        | NOT NULL. REFERENCES users(ID) ON DELETE RESTRICT  |                                     |                          |
| CREATION  | TIMESTAMPTZ | NOT NULL. DEFAULT now()                            |                                     |                          |
| CONTENT   | TEXT        | NOT NULL. CHECK(LENGTH(CONTENT) \< 10000)          |                                     |                          |
| TICKET_ID | UUID        | NOT NULL. REFERENCES tickets(ID) ON DELETE CASCADE |                                     |                          |
| PARENT_ID | UUID        | REFERENCES comments(ID) ON DELETE CASCADE          | Parent comment when this is a reply | null = top-level comment |
