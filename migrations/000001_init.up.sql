-- Очищаем всю базу данных
DO $$
    DECLARE
        current_table text;
    BEGIN
        FOR current_table IN (SELECT table_name FROM information_schema.tables WHERE table_schema = 'public')
            LOOP
                EXECUTE 'DROP TABLE IF EXISTS "' || current_table || '" CASCADE';
            END LOOP;
    END $$;

select * from information_schema.tables;

-- Импортируем, если нет, тип данных hstore
CREATE EXTENSION IF NOT EXISTS hstore;

create table movie(
                      id serial primary key,
                      title varchar(255) not null ,
                      description text,
                      duration int,
                      release_year int,
                      director int,
                      rating hstore
);

create table cinema(
                       id serial primary key,
                       "name" varchar(255) not null,
                       address varchar(255) not null
);

create table cinema_hall(
                            id serial primary key,
                            "name" varchar(255) not null,
                            capacity int,
                            "class" varchar(255),
                            cinema_id int references cinema(id)
);

create table session(
                        id serial primary key,
                        "date" date,
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
                       cost float,
                       seat int,
                       session_id int references session(id),
                       user_id int references "user"(id)
);

create table friend_list(
                            user_id int references "user"(id),
                            friend_id int references "user"(id)
);

-- Вставка данных в таблицу cinema
INSERT INTO cinema (name, address) VALUES
                                       ('Кинотеатр "Зеркальный Экран"', 'ул. Центральная, д. 456, Город Г'),
                                       ('Арт-Кинотеатр "Творчество"', 'пр. Лесной, д. 789, Город Д'),
                                       ('Императорский Кинопалац', 'ул. Звездная, д. 101, Город Е'),
                                       ('Семейное КиноДейство', 'ул. Солнечная, д. 789, Город З'),
                                       ('Луна-Парк Кинотеатр', 'пр. Парковый, д. 101, Город И'),
                                       ('Гелиос Кинотеатр', 'ул. Луна, д. 456, Город К');

-- Вставка данных в таблицу movie
INSERT INTO movie (title, description, duration, release_year, director, rating) VALUES
                                                                                     ('Тайны Вселенной: Загадки Галактики', 'Фантастическое путешествие в неизведанные уголки космоса', 140, 2022, 3, '"IMDb"=>"9.0", "Metacritic"=>"88"'),
                                                                                     ('Следы Времени: Драма в Четырех Актах', 'Глубокая драма о течении времени и выборе', 110, 2021, 4, '"IMDb"=>"8.5", "Rotten Tomatoes"=>"87"'),
                                                                                     ('Ловушка Судьбы: Игра на Выживание', 'Захватывающий триллер с неожиданными сюжетными поворотами', 125, 2023, 2, '"IMDb"=>"8.8", "Rotten Tomatoes"=>"90"'),
                                                                                     ('По ту сторону горизонта', 'Романтическая история о поиске своего места в мире', 118, 2020, 5, '"IMDb"=>"8.7", "Metacritic"=>"85"'),
                                                                                     ('Тень Великого Взрыва', 'Научно-фантастический боевик с элементами детектива', 132, 2022, 1, '"IMDb"=>"8.9", "Rotten Tomatoes"=>"92"'),
                                                                                     ('Спасти Галактику: Последний Рубеж', 'Эпическое продолжение приключений в космосе', 150, 2023, 3, '"IMDb"=>"9.2", "Rotten Tomatoes"=>"94"');

-- Вставка данных в таблицу cinema_hall
INSERT INTO cinema_hall (name, capacity, class, cinema_id) VALUES
                                                               ('Зал 4', 80, 'Стандарт', 1),
                                                               ('Зал 5', 120, 'VIP', 1),
                                                               ('Зал 6', 90, 'Премиум', 1),
                                                               ('Зал 7', 100, 'Стандарт', 4),
                                                               ('Зал 8', 80, 'VIP', 5),
                                                               ('Зал 9', 120, 'Премиум', 6);

-- Вставка данных в таблицу session
INSERT INTO session (date, movie_id, cinema_hall_id, ticket_left) VALUES
                                                                      ('2023-01-04 19:15:00', 1, 1, 40),
                                                                      ('2023-01-05 21:00:00', 2, 2, 100),
                                                                      ('2023-01-06 18:30:00', 3, 3, 70),
                                                                      ('2023-01-07 20:30:00', 4, 1, 60),
                                                                      ('2023-01-08 22:15:00', 5, 2, 75),
                                                                      ('2023-01-09 17:45:00', 6, 1, 90);

-- Вставка данных в таблицу user
INSERT INTO "user" (first_name, last_name, number, email) VALUES
                                                              ('Ольга', 'Петрова', '456123789', 'olga.petrova@example.com'),
                                                              ('Алексей', 'Кузнецов', '321987654', 'aleksey.kuznetsov@example.com'),
                                                              ('Дмитрий', 'Иванов', '789456123', 'dmitriy.ivanov@example.com'),
                                                              ('Мария', 'Смирнова', '654321987', 'maria.smirnova@example.com');

-- Вставка данных в таблицу ticket
INSERT INTO ticket (cost, seat, session_id, user_id) VALUES
                                                         (11.2, 4, 1, 3),
                                                         (13.5, 5, 2, 4),
                                                         (10.0, 6, 3, 3),
                                                         (12.0, 7, 4, 1),
                                                         (14.5, 8, 5, 2),
                                                         (11.8, 9, 6, 2);

-- Вставка данных в таблицу friend_list
INSERT INTO friend_list (user_id, friend_id) VALUES
                                                 (3, 4),
                                                 (4, 3);
