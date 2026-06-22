| name | type | constraints                           | description | notes |
|------|------|---------------------------------------|-------------|-------|
| ID   | UUID | PRIMARY KEY                           |             |       |
| NAME | TEXT | NOT NULL. CHECK(LENGTH(NAME) \< 1000) |             |       |
