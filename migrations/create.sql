CREATE TABLE actors (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        gender VARCHAR(10) NOT NULL CHECK (gender IN ('Male', 'Female')),
                        birth_date DATE NOT NULL
);

CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(150) NOT NULL,
                        description VARCHAR(1000),
                        release_date DATE NOT NULL,
                        rating DECIMAL(3, 1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE actor_movie (
                             actor_id INTEGER REFERENCES actors(id),
                             movie_id INTEGER REFERENCES movies(id),
                             PRIMARY KEY (actor_id, movie_id)
);