DROP TABLE IF EXISTS files;

CREATE TABLE audio_files (
    album TEXT NOT NULL,
    bit_depth INTEGER NOT NULL,
    compression_mode TEXT NOT NULL,
    content_type TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    duration BIGINT NOT NULL,
    file_extension TEXT NOT NULL,
    filename TEXT NOT NULL,
    format TEXT NOT NULL,
    genre TEXT NOT NULL,
    id TEXT PRIMARY KEY ,
    length BIGINT NOT NULL,
    performer TEXT NOT NULL,
    recorded_date INTEGER NOT NULL,
    sampling_rate INTEGER NOT NULL,
    status TEXT NOT NULL,
    title TEXT NOT NULL,
    total_track_count INTEGER NOT NULL,
    track_number INTEGER NOT NULL
);
