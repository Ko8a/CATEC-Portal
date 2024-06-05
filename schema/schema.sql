CREATE TABLE role
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null
);

INSERT INTO role (name) VALUES ('guest'), ('student'), ('parent'), ('teacher'), ('manager'), ('administrator');

CREATE TABLE groups
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null
);

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    surname varchar(255) not null,
    age int,
    email varchar(255) not null,
    password_hash varchar(255) not null,
    phone varchar(255),
    group_id int references groups(id) on delete cascade,
    time_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id int references role(id) on delete cascade
);

INSERT INTO users (name, surname, email, password_hash, role_id) VALUES ('Super', 'Admin', 'superadmin@mail.ru', '736a6e6361246b316b323370646a214b4a766b31326b6a316b3a960464d36c1b8bad183ed57ee79c0e39953cce', (SELECT id FROM role WHERE name = 'administrator'));

CREATE TABLE lesson_type
(
    id SERIAL PRIMARY KEY,
    name varchar(255)
);

CREATE TABLE lesson
(
    id SERIAL PRIMARY KEY,
    lesson_name varchar(255) not null,
    date DATE not null,
    teacher_id int references users(id) on delete cascade default null,
    group_id int references groups(id) on delete cascade default null,
    title varchar(255) not null,
    is_online boolean default false,
    start_time TIMESTAMP default null,
    end_time TIMESTAMP default null,
    type_id int references lesson_type(id) on delete cascade default null,
    classroom varchar(255)
);

CREATE TABLE student_performance
(
    id SERIAL PRIMARY KEY,
    user_id int references users(id) on delete cascade not null,
    lesson_id int references lesson(id) on delete cascade not null,
    mark int default null,
    is_came boolean default false
);

CREATE TABLE day_schedule
(
    id SERIAL PRIMARY KEY,
    lesson_1 int references lesson(id) on delete cascade,
    lesson_2 int references lesson(id) on delete cascade,
    lesson_3 int references lesson(id) on delete cascade,
    lesson_4 int references lesson(id) on delete cascade,
    lesson_5 int references lesson(id) on delete cascade,
    lesson_6 int references lesson(id) on delete cascade

);

CREATE TABLE schedule
(
    id SERIAL PRIMARY KEY,
    group_id int references groups(id) on delete cascade,
    monday int references day_schedule(id) on delete cascade,
    tuesday int references day_schedule(id) on delete cascade,
    wednesday int references day_schedule(id) on delete cascade,
    thursday int references day_schedule(id) on delete cascade,
    friday int references day_schedule(id) on delete cascade,
    saturday int references day_schedule(id) on delete cascade,
    sunday int references day_schedule(id) on delete cascade
);

CREATE TABLE weeks
(
    id SERIAL PRIMARY KEY,
    group_id int references groups(id) on delete cascade not null,
    start_date date not null,
    end_date date not null,
    type_id int references lesson_type(id) on delete cascade not null
);

-- Create table lesson_names
CREATE TABLE lesson_names (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(255) NOT NULL
);

-- Create table lesson_times
CREATE TABLE lesson_times (
                              id SERIAL PRIMARY KEY,
                              start_time TIMESTAMP NOT NULL,
                              end_time TIMESTAMP NOT NULL
);

-- Alter table lessons to add foreign key
-- ALTER TABLE lesson
--     ADD COLUMN lesson_name_id int references lesson_names(id) on delete cascade,
--     DROP COLUMN name;