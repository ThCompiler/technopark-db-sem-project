CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    id       bigserial not null primary key,
    nickname citext not null unique,
    fullname text   not null,
    about    text,
    email    citext not null unique
);

CREATE TABLE IF NOT EXISTS forum
(
    id bigserial primary key,
    title         text   not null,
    user_nickname bigint not null references users (id),
    slug          citext not null unique,
    posts         bigint default 0 not null,
    threads       integer    default 0 not null
);


CREATE TABLE IF NOT EXISTS threads
(
    id      bigserial not null primary key,
    title   text      not null,
    author  bigint    not null references users (id),
    forum   bigint    not null references forum (id),
    message text      not null,
    votes   integer     default 0 not null,
    slug    citext not null unique,
    created timestamptz default now()::timestamptz not null
);

create table posts
(
    id        bigserial  not null primary key,
    parent    integer                  default 0 not null,
    author    bigint not null references users(id),
    message   text   not null,
    is_edited boolean                  default false,
    forum     bigint not null references forum(id),
    thread    integer not null references threads(id),
    created timestamptz default now()::timestamptz not null,
    path      bigint[]                 default ARRAY []::integer[] not null
);

CREATE TABLE IF NOT EXISTS votes
(
    nickname  bigint not null references users (id),
    thread_id bigint    not null references threads (id),
    voice     int    not null,
    unique (nickname, thread_id)
);

create table if not exists users_to_forums
(
    nickname bigint not null references users (id),
    forum    bigint not null references forum (id),
    unique (nickname, forum)
);