-- Creating the "Movies" table
CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(255),
                        description TEXT,
                        director VARCHAR(255),
                        duration INT,
                        release_year INT
);

-- Creating the "Cinema Halls" table
CREATE TABLE cinema_halls (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(255),
                              seat_count INT
);

-- Creating the "Screenings" table
CREATE TABLE screenings (
                            id SERIAL PRIMARY KEY,
                            movie_id INT REFERENCES movies(id),
                            cinema_hall_id INT REFERENCES cinema_halls(id),
                            date_time TIMESTAMP
);

-- Creating the "Orders" table
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        screening_id INT REFERENCES screenings(id),
                        buyer_name VARCHAR(255),
                        ticket_count INT
);
