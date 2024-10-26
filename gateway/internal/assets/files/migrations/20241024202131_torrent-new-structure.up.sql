truncate table torrent;

alter table torrent drop column description;
alter table torrent drop column title;
alter table torrent drop column metadata_info;
alter table torrent drop column filename;

alter table torrent add hash text not null;
alter table torrent add name text not null;
alter table torrent add piece_count integer not null;
alter table torrent add piece_size text not null;
alter table torrent add privacy text not null;
alter table torrent add total_size text not null;
alter table torrent add filename text not null;

create table torrent_file (
    id text primary key,
    torrent_id text not null,
    file_name text not null,
    file_size text not null
);

alter table torrent_file add constraint fk_torrent_file_torrent
foreign key (torrent_id) references torrent(id);