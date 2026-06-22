boards
```{=latex}    
\begin{tiny}    
```    
```md
| name          | type        | constraints                                       | description                     | notes                  |
|---------------|-------------|---------------------------------------------------|---------------------------------|------------------------|
| ID            | UUID        | PRIMARY KEY                                       |                                 |                        |
| NAME          | TEXT        | NOT NULL. CHECK(LENGTH(NAME) \< 1000)             |                                 |                        |
| DESCRIPTION   | TEXT        | CHECK(LENGTH(DESCRIPTION) \< 10000)               |                                 |                        |
| AUTHOR        | UUID        | NOT NULL. REFERENCES users(ID) ON DELETE RESTRICT |                                 |                        |
| CREATION      | TIMESTAMPTZ | NOT NULL. DEFAULT now()                           |                                 |                        |
| MOTHER_TICKET | UUID        | REFERENCES tickets(ID) ON DELETE CASCADE          | Ticket that contains this board | null = top-level board |

```    
```{=latex}    
\end{tiny}    
```  
columns
```{=latex}    
\begin{tiny}    
```    
```md
| name     | type    | constraints                                       | description | notes                             |
|----------|---------|---------------------------------------------------|-------------|-----------------------------------|
| ID       | UUID    | PRIMARY KEY                                       |             |                                   |
| BOARD_ID | UUID    | NOT NULL. REFERENCES boards(ID) ON DELETE CASCADE |             |                                   |
| NAME     | TEXT    | NOT NULL. CHECK(LENGTH(NAME) \< 1000)             |             |                                   |
| POSITION | INTEGER | NOT NULL. CHECK(POSITION \>= 0)                   |             | 0-based ordering within the board |

```    
```{=latex}    
\end{tiny}    
```  
comments
```{=latex}    
\begin{tiny}    
```    
```md
| name      | type        | constraints                                        | description                         | notes                    |
|-----------|-------------|----------------------------------------------------|-------------------------------------|--------------------------|
| ID        | UUID        | PRIMARY KEY                                        |                                     |                          |
| AUTHOR    | UUID        | NOT NULL. REFERENCES users(ID) ON DELETE RESTRICT  |                                     |                          |
| CREATION  | TIMESTAMPTZ | NOT NULL. DEFAULT now()                            |                                     |                          |
| CONTENT   | TEXT        | NOT NULL. CHECK(LENGTH(CONTENT) \< 10000)          |                                     |                          |
| TICKET_ID | UUID        | NOT NULL. REFERENCES tickets(ID) ON DELETE CASCADE |                                     |                          |
| PARENT_ID | UUID        | REFERENCES comments(ID) ON DELETE CASCADE          | Parent comment when this is a reply | null = top-level comment |

```    
```{=latex}    
\end{tiny}    
```  
tags
```{=latex}    
\begin{tiny}    
```    
```md
| name        | type    | constraints                                      | description | notes             |
|-------------|---------|--------------------------------------------------|-------------|-------------------|
| ID          | UUID    | PRIMARY KEY                                      |             |                   |
| NAME        | TEXT    | NOT NULL. CHECK(LENGTH(NAME) \< 1000)            |             |                   |
| DESCRIPTION | TEXT    | CHECK(LENGTH(DESCRIPTION) \< 10000)              |             |                   |
| COLOUR      | TEXT    | NOT NULL. CHECK(COLOUR ~ '^#\[0-9A-Fa-f\]{6}\$') |             | e.g. \#1D9E75     |
| SPECIAL     | BOOLEAN | NOT NULL. DEFAULT false                          |             |                   |
| BOARD_ID    | UUID    | REFERENCES boards(ID) ON DELETE CASCADE          |             | null = global tag |

```    
```{=latex}    
\end{tiny}    
```  
team_members
```{=latex}    
\begin{tiny}    
```    
```md
| name    | type | constraints                                      | description | notes                            |
|---------|------|--------------------------------------------------|-------------|----------------------------------|
| TEAM_ID | UUID | NOT NULL. REFERENCES teams(ID) ON DELETE CASCADE |             | Composite PK (TEAM_ID + USER_ID) |
| USER_ID | UUID | NOT NULL. REFERENCES users(ID) ON DELETE CASCADE |             | Composite PK (TEAM_ID + USER_ID) |
| ROLE    | TEXT | NOT NULL. DEFAULT 'member'                       |             | e.g. owner / admin / member      |

```    
```{=latex}    
\end{tiny}    
```  
teams
```{=latex}    
\begin{tiny}    
```    
```md
| name | type | constraints                           | description | notes |
|------|------|---------------------------------------|-------------|-------|
| ID   | UUID | PRIMARY KEY                           |             |       |
| NAME | TEXT | NOT NULL. CHECK(LENGTH(NAME) \< 1000) |             |       |

```    
```{=latex}    
\end{tiny}    
```  
ticket_assignee
```{=latex}    
\begin{tiny}    
```    
```md
| name      | type | constraints                                        | description | notes                              |
|-----------|------|----------------------------------------------------|-------------|------------------------------------|
| TICKET_ID | UUID | NOT NULL. REFERENCES tickets(ID) ON DELETE CASCADE |             | Composite PK (TICKET_ID + USER_ID) |
| USER_ID   | UUID | NOT NULL. REFERENCES users(ID) ON DELETE CASCADE   |             | Composite PK (TICKET_ID + USER_ID) |

```    
```{=latex}    
\end{tiny}    
```  
ticket_tags
```{=latex}    
\begin{tiny}    
```    
```md
| name      | type | constraints                                        | description | notes                             |
|-----------|------|----------------------------------------------------|-------------|-----------------------------------|
| TICKET_ID | UUID | NOT NULL. REFERENCES tickets(ID) ON DELETE CASCADE |             | Composite PK (TICKET_ID + TAG_ID) |
| TAG_ID    | UUID | NOT NULL. REFERENCES tags(ID) ON DELETE CASCADE    |             | Composite PK (TICKET_ID + TAG_ID) |

```    
```{=latex}    
\end{tiny}    
```  
tickets
```{=latex}    
\begin{tiny}    
```    
```md
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

```    
```{=latex}    
\end{tiny}    
```  
users
```{=latex}    
\begin{tiny}    
```    
```md
| name | type | constraints                           | description | notes |
|------|------|---------------------------------------|-------------|-------|
| ID   | UUID | PRIMARY KEY                           |             |       |
| NAME | TEXT | NOT NULL. CHECK(LENGTH(NAME) \< 1000) |             |       |

```    
```{=latex}    
\end{tiny}    
```  