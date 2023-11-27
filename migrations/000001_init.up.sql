-- Creating the "Cinema" database
DO $$
    DECLARE
        current_table text;  -- Используйте более конкретное имя переменной, чтобы избежать неоднозначности
    BEGIN
        FOR current_table IN (SELECT table_name FROM information_schema.tables WHERE table_schema = 'public')
            LOOP
                EXECUTE 'DROP TABLE IF EXISTS "' || current_table || '" CASCADE';
            END LOOP;
    END $$;

select * from information_schema.tables;

CREATE EXTENSION IF NOT EXISTS hstore;

create table movie(
                      id serial primary key,
                      title varchar(255) not null,
                      description text,
                      duration int,
                      release_year int,
                      director int,
                      rating hstore
);

create table cinema(
                       id serial primary key,
                       name varchar(255) not null,
                       address varchar(255) not null
);

create table cinema_hall(
                            id serial primary key,
                            name varchar(255) not null,
                            capacity int,
                            class varchar(255),
                            cinema_id int references cinema(id)
);

create table session(
                        id serial primary key,
                        date date,
                        movie_id int references movie(id),
                        cinema_hall_id int references cinema_hall(id),
                        ticket_left int
);

create table "user"(
                       id serial primary key,
                       first_name varchar(255),
                       last_name varchar(255),
                       number varchar(255),
                       email varchar(255)
);

create table ticket(
                       id serial primary key,
                       class varchar(255),
                       cost float,
                       seat int,
                       session_id int references session(id),
                       user_id int references "user"(id)
);

create table friend_list(
                            user_id int references "user"(id),
                            friend_id int references "user"(id)
);

-- Fill "Cinema" Database
-- Заполнение таблицы cinema
INSERT INTO cinema (name, address) VALUES
                                       ('Cinema1', 'Address1'),
                                       ('Cinema2', 'Address2'),
                                       ('Cinema3', 'Address3');

-- Заполнение таблицы cinema_hall
INSERT INTO cinema_hall (name, capacity, class, cinema_id) VALUES
                                                               ('Hall1', 100, 'Standard', 1),
                                                               ('Hall2', 150, 'VIP', 2),
                                                               ('Hall3', 120, 'Standard', 3);

-- Заполнение таблицы movie
INSERT INTO movie (title, description, duration, release_year, director, rating) VALUES
                                                                                     ('Movie1', 'Description1', 120, 2020, 1, '"IMDb"=>"8.5", "RottenTomatoes"=>"90"'),
                                                                                     ('Movie2', 'Description2', 110, 2021, 2, '"IMDb"=>"7.8", "RottenTomatoes"=>"85"'),
                                                                                     ('Movie3', 'Description3', 130, 2019, 3, '"IMDb"=>"9.0", "RottenTomatoes"=>"95"');

-- Заполнение таблицы session
INSERT INTO session (date, movie_id, cinema_hall_id, ticket_left) VALUES
                                                                      ('2023-01-01', 1, 1, 50),
                                                                      ('2023-01-02', 2, 2, 75),
                                                                      ('2023-01-03', 3, 3, 60);

-- Заполнение таблицы user
INSERT INTO "user" (first_name, last_name, number, email) VALUES
                                                              ('John', 'Doe', '123456789', 'john.doe@example.com'),
                                                              ('Alice', 'Smith', '987654321', 'alice.smith@example.com');

-- Заполнение таблицы ticket
INSERT INTO ticket (class, cost, seat, session_id, user_id) VALUES
                                                                ('Standard', 10.0, 1, 1, 1),
                                                                ('VIP', 20.0, 2, 2, 2),
                                                                ('Standard', 12.0, 3, 3, 1);

-- Заполнение таблицы friend_list
INSERT INTO friend_list (user_id, friend_id) VALUES
                                                 (1, 2),
                                                 (2, 1);

