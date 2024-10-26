alter table torrent add filename text null;
alter table torrent add metadata_info text null;
alter table torrent add metadata_files text null;
alter table torrent add last_error text null;
alter table torrent add updated timestamp without time zone null;
