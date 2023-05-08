PRAGMA journal_mode = wal;

CREATE TABLE IF NOT EXISTS "todos" (
    "id" TEXT PRIMARY KEY,
    "content" TEXT NOT NULL,
    "completed" bool NOT NULL DEFAULT '1'
);