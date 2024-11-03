alter table torrent add status text not null;
alter table torrent add user_id text not null;

alter table torrent add constraint torrent_fk_users foreign key (user_id)
references users(id);