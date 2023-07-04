DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    author_name TEXT REFERENCES authors(name) NOT NULL,
    created_at BIGINT NOT NULL,
    published_at BIGINT NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO posts (id, author_id, title, content, created_at) VALUES (0, 0, 'Статья', 'Содержание статьи', 0);