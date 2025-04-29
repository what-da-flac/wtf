DROP TABLE IF EXISTS audio_files;

CREATE TABLE files (
    id TEXT PRIMARY KEY ,
    created TIMESTAMP NOT NULL,
    filename TEXT NOT NULL,
    length BIGINT NOT NULL,
    content_type TEXT NOT NULL,
    status TEXT NOT NULL
);