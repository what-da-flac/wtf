create table if not exists roles
(
    id          text not null primary key,
    name        text not null
        constraint roles_name_unique unique,
    description text not null
);

create table if not exists users
(
    email      text primary key,
    id         text                     not null unique,
    name       text                     not null,
    image_url  text,
    created    timestamp with time zone not null,
    last_login timestamp with time zone not null
);

create table if not exists user_role
(
    role_id text not null
        constraint fk_user_role_role
            references roles,
    user_id text not null
        constraint fk_user_role_user
            references users (id),
    constraint pk_user_role
        primary key (role_id, user_id)
);

