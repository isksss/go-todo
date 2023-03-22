CREATE TABLE todos (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  completed  BOOLEAN DEFAULT 0
);