-- +goose Up

create table class_list (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  name text not null,
  section public.section_enum not null
);

create table subject_list (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  name text not null
);

create table nur_subjects_questions (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  subject_id uuid references subject_list (id),
  question text not null
);

create table sessions (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  name text not null
);

create table terms (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  session_id uuid references sessions (id),
  term_id public.term_enum not null,
  vacation_date timestamp default now(),
  next_resumption_date timestamp default now(),
  number_school_opened bigint,
  lock boolean not null
);

create table staff_profiles (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  first_name text not null,
  middle_name text not null,
  last_name text not null,
  passport text not null,
  date_of_birth date default now() not null,
  gender public.gender_enum not null,
  admission_no bigint not null,
  role public.roles_enum not null,
  term_id uuid references terms (id),
  joining_date timestamp default now() not null,
  active boolean not null,
  password_change boolean not null,
  phone_no text not null
);

create table classes (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  term_id uuid references terms (id),
  class_id uuid references class_list (id),
  class_teacher_id uuid references staff_profiles (id)
);

create table subjects (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  term_class_id uuid references classes (id),
  subject_id uuid references subject_list (id),
  teacher_id uuid references staff_profiles (id)
);

create table student_profiles (
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  first_name text not null,
  middle_name text not null,
  last_name text not null,
  passport text not null,
  date_of_birth date default now() not null,
  gender public.gender_enum not null,
  id uuid default uuid_generate_v4() primary key,
  admission_no bigint not null,
  joined_class_id uuid references classes (id),
  current_class_id uuid references classes (id),
  joining_date timestamp default now() not null,
  password_change boolean not null,
  active boolean not null
);

create table nursery_first_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table primary_third_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table nursery_second_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table primary_third_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references primary_third_term_results (id),
  subject_id uuid references subjects (id),
  test_2_score bigint not null,
  test_3_score bigint not null,
  test_4_score bigint not null,
  exam_score bigint not null,
  test_1_score bigint not null
);

create table primary_first_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table primary_first_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references primary_first_term_results (id),
  subject_id uuid references subjects (id),
  test_2_score bigint not null,
  test_3_score bigint not null,
  test_4_score bigint not null,
  exam_score bigint not null,
  test_1_score bigint not null
);

create table nursery_third_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table nursery_third_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references nursery_third_term_results (id),
  subject_id uuid references subjects (id),
  score public.nursery_subject_score_enum not null,
  question_id uuid references nur_subjects_questions (id)
);

create table nursery_first_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references nursery_first_term_results (id),
  subject_id uuid references subjects (id),
  score public.nursery_subject_score_enum not null,
  question_id uuid references nur_subjects_questions (id)
);

create table primary_second_term_results (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  student_id uuid references student_profiles (id),
  class_id uuid references classes (id)
);

create table primary_second_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references primary_second_term_results (id),
  subject_id uuid references subjects (id),
  test_2_score bigint not null,
  test_3_score bigint not null,
  test_4_score bigint not null,
  exam_score bigint not null,
  test_1_score bigint not null
);

create table nursery_second_term_subject_scores (
  id uuid default uuid_generate_v4() primary key,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null,
  result_id uuid references nursery_second_term_results (id),
  subject_id uuid references subjects (id),
  score public.nursery_subject_score_enum not null,
  question_id uuid references nur_subjects_questions (id)
);

-- +goose Down

drop table nursery_second_term_subject_scores;
drop table nursery_first_term_subject_scores;
drop table nursery_third_term_subject_scores;
drop table primary_first_term_subject_scores;
drop table primary_third_term_subject_scores;
drop table primary_second_term_subject_scores;
drop table nur_subjects_questions;
drop table primary_first_term_results;
drop table nursery_third_term_results;
drop table nursery_first_term_results;
drop table nursery_second_term_results;
drop table primary_third_term_results;
drop table primary_second_term_results;
drop table student_profiles;
drop table subjects;
drop table classes;
drop table staff_profiles;
drop table terms;
drop table class_list;
drop table subject_list;
drop table sessions;