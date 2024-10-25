create table users (
    id text primary key ,
    name text not null,
    email text not null unique,
    created timestamp with time zone not null
);

