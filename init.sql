create extension if not exists citext;

drop table if exists sessions;
drop table if exists users;

drop index if exists users_login;
drop index if exists sessions_token;
drop index if exists sessions_user;

create table users (
    login citext primary key,
    email varchar not null,
    phone varchar not null,
    password varchar not null
);

create index if not exists users_login on users (login);

create table sessions (
    token varchar primary key,
    user_id citext references users (login),
    expiration timestamptz not null default current_timestamp
);

create index if not exists sessions_token on sessions (token);
create index if not exists sessions_user on sessions (user_id);