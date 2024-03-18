

CREATE TABLE IF NOT EXISTS authors (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    author_id INTEGER  REFERENCES authors(id),
    title TEXT,
    content TEXT,
    created BIGINT NOT NULL DEFAULT extract(epoch from now())
);