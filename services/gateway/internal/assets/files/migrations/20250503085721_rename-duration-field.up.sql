ALTER TABLE audio_files RENAME duration to duration_seconds;
ALTER TABLE audio_files DROP COLUMN bit_depth;
ALTER TABLE audio_files DROP COLUMN status;