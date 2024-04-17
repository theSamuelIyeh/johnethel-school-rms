-- +goose Up

CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TYPE nursery_subject_score_enum AS ENUM ('A', 'B', 'C', 'D');
CREATE TYPE roles_enum AS ENUM ('teacher', 'accountant', 'admin', 'student');
CREATE TYPE section_enum AS ENUM ('Primary Section', 'Nursery Section');
CREATE TYPE term_enum AS ENUM ('First Term', 'Second Term', 'Third Term');

-- +goose Down
DROP TYPE gender_enum;
DROP TYPE nursery_subject_score_enum;
DROP TYPE roles_enum;
DROP TYPE section_enum;
DROP TYPE term_enum;