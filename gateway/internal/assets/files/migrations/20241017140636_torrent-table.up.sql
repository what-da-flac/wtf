create table torrent (
    id text primary key ,
    created timestamp with time zone not null,
    description text not null,
    magnet_link text not null,
    title text not null
);