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