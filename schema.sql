DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);

INSERT INTO authors (name) VALUES ('Дмитрий');
INSERT INTO authors (name) VALUES ('Semen');
INSERT INTO posts (author_id, title, content) VALUES (1, 'Статья', 'Содержание статьи');
INSERT INTO posts (author_id, title, content) VALUES (2, 'Статья2', 'Содержание статьи 2');
