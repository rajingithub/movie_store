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



GOROUTINES : 

goroutine is a function that can run in background , so rest of the code can excute without waiting for it to finish. 


eg : 

1.
```
func sample():
    time.sleep(10)
    print("wait completed")
    return

func main():
    print("hello")
    sample()
    print("world")
```
output : 

```
hello
wait completed
world
```

- without using go routines, if we run main function, it will main function needs to wait till sample function finishes (which has 10secs delay)




2.

```
func sample():
    time.sleep(10)
    print("wait completed")
    return

func main():
    print("hello")
    go sample()
    print("world")
```

```
hello
world
wait completed
```

- if we use goroutines , then main function need not wait till sample function gets completed.


Concurrency : Dealing with multiple things  at same time
Parallelism : Running Multiple things at once. 

Go routines guarentees concurrency, whether they can be run in parallel depends upon number of CPU cores on the machine.

if we have cpu with 4 cores, go run 4 go routines truly in parallel, if we have 4 cpu cores and 100 goroutines
go switches those 100 go routines between 4 CPU cores.

Channels :


channel is a pipe that go routine uses to send/receive data with other go routine.
2 types : 

blocking channel : A blocking send/receive waits until the other side is ready
A non-blocking send/receive does not wait; it either succeeds or immediately skips.

