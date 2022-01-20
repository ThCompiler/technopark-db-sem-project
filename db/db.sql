CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    nickname citext not null unique primary key,
    fullname text   not null,
    about    text,
    email    citext not null unique
);

CREATE TABLE IF NOT EXISTS forums
(
    title         text              not null,
    user_nickname citext            not null references users (nickname),
    slug          citext            not null primary key,
    posts         bigint  default 0 not null,
    threads       integer default 0 not null
);


CREATE TABLE IF NOT EXISTS threads
(
    id      bigserial                              not null primary key,
    title   text                                   not null,
    author  citext                                 not null references users (nickname),
    forum   citext                                 not null references forums (slug),
    message text                                   not null,
    votes   integer     default 0                  not null,
    slug    citext                                 unique,
    created timestamptz default now()::timestamptz not null
);

CREATE TABLE IF NOT EXISTS posts
(
    id        bigserial                               not null primary key,
    parent    integer     default 0                   not null,
    author    citext                                  not null references users (nickname),
    message   text                                    not null,
    is_edited boolean     default false               not null,
    forum     citext                                  not null references forums (slug),
    thread    bigint                                  not null references threads (id),
    created   timestamptz default now()::timestamptz  not null,
    path      bigint[]    default ARRAY []::integer[] not null
);

CREATE TABLE IF NOT EXISTS votes
(
    nickname  citext not null references users (nickname),
    thread_id bigint not null references threads (id),
    voice     int    not null,
    unique (nickname, thread_id)
);

CREATE TABLE IF NOT EXISTS users_to_forums
(
    nickname citext not null references users (nickname),
    forum    citext not null references forums (slug),
    unique (nickname, forum)
);

-- forums trigger --

CREATE OR REPLACE FUNCTION add_post_to_forum()
RETURNS TRIGGER AS
$$
BEGIN
    UPDATE forums
    SET posts = forums.posts + 1
    WHERE slug = NEW.forum;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_post_to_forum_trg
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE add_post_to_forum();

CREATE OR REPLACE FUNCTION add_thread_to_forum()
RETURNS TRIGGER AS
$$
BEGIN
    UPDATE forums
    SET threads = forums.threads + 1
    WHERE slug = NEW.forum;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_thread_to_forum_trg
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE add_thread_to_forum();


-- user_to_forums trigger --


CREATE OR REPLACE FUNCTION add_user_to_forum()
RETURNS TRIGGER AS
$$
BEGIN
    INSERT INTO users_to_forums (nickname, forum)
    VALUES (NEW.author, NEW.forum)
    ON CONFLICT do nothing;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_user_forum
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE add_user_to_forum();

CREATE TRIGGER update_users_forum
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE add_user_to_forum();


-- votes trigger --


CREATE OR REPLACE FUNCTION add_votes_into_threads()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE threads
    SET votes = votes + NEW.voice
    WHERE id = NEW.thread_id;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER insert_votes
    AFTER INSERT
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE add_votes_into_threads();

CREATE OR REPLACE FUNCTION update_votes_in_threads()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE threads
    SET votes = votes + NEW.voice - OLD.voice
    WHERE id = NEW.thread_id;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_votes
    AFTER UPDATE
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE update_votes_in_threads();


-- post trigger --


CREATE OR REPLACE FUNCTION update_post_past() RETURNS TRIGGER AS
$$
BEGIN
    new.path = (SELECT path FROM posts WHERE id = new.parent) || new.id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_post_path
    BEFORE INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE update_post_past();


