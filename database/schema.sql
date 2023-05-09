PRAGMA journal_mode = wal;
PRAGMA busy_timeout = 5000;
PRAGMA synchronous = NORMAL;

CREATE TABLE IF NOT EXISTS "todos" (
    "id" TEXT PRIMARY KEY,
    "content" TEXT NOT NULL,
    "completed" bool NOT NULL DEFAULT '1'
);