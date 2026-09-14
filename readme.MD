# project structure : 

    http request
        |
    handler.go (request first comes here, understands http methods, URL, request path, request body, http status code)
        |
    service.go (understands business rules eg:A movie title cannot be empty.)
        |
    repository.go (talks/Interacts with DB)

- model.go - used to represent data
- go.mod is like requirements.txt in python which contains required go version and dependencies
- cmd is an convention in go which contains applications/executables (or kind of aplication entry point)

In MySQL : 
- create database movie_store_db;
- use movie_store_db;

DDL : 

```
CREATE DATABASE IF NOT EXISTS movie_store;

USE movie_store;


-- Movies
CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    release_date DATE NULL
);


-- Actors
CREATE TABLE actors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(500) NULL
);


-- Movie <-> Actor mapping
CREATE TABLE movie_actors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,

    CONSTRAINT fk_movie_actors_movie
        FOREIGN KEY (movie_id)
        REFERENCES movies(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_movie_actors_actor
        FOREIGN KEY (actor_id)
        REFERENCES actors(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_movie_actor
        UNIQUE (movie_id, actor_id)
);
```

DML:

```
INSERT INTO movies (name, release_date)
VALUES
    ('Inception', '2010-07-16'),
    ('The Dark Knight', '2008-07-18'),
    ('Interstellar', '2014-11-07'),
    ('The Matrix', '1999-03-31'),
    ('Gladiator', NULL);


INSERT INTO actors (name, image)
VALUES
    ('Leonardo DiCaprio', NULL),
    ('Christian Bale', NULL),
    ('Matthew McConaughey', NULL),
    ('Keanu Reeves', NULL),
    ('Russell Crowe', NULL);

INSERT INTO movie_actors (movie_id, actor_id)
VALUES
    -- Inception
    (1, 1),

    -- The Dark Knight
    (2, 2),

    -- Interstellar
    (3, 3),

    -- The Matrix
    (4, 4),

    -- Gladiator
    (5, 5),

    -- Additional relationships
    (1, 3),
    (2, 4),
    (3, 1),
    (4, 2),
    (5, 2);
```


- `go get github.com/go-sql-driver/mysql` (to get required mysql driver)

- why are we creating dependencies in cmd/api/main.go : 
    Main Job of main.go is create dependencies, connect dependencies and start application


