# ShareCode
An simple platform for sharing code

## Setup database
```sql
-- Create a new database.
CREATE DATABASE sharecode;

-- Switvh to using 'sharecode' database.
USE sharecode;

-- Create 'snippets' table.
CREATE TABLE snippets (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);

-- Add an index on 'created' column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records.
INSERT INTO snippets (title, content, created, expires) VALUES (
  'An old silent pond',
  'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'Over the wintry forest',
  'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'First autumn morning',
  'First autumn morning\nthe mirror I stare into\nshows my father''s face.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

## Create database user
```sql
CREATE USER 'sharecode'@'localhost';
GRANT SELECT, INSERT ON sharecode.* TO 'sharecode'@'localhost';

-- Replace 'pass' with your password.
ALTER USER 'sharecode'@'localhost' IDENTIFIED BY 'pass';
```
