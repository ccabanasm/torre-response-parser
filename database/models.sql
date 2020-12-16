-- Data structure for Jobs
CREATE TABLE IF NOT EXISTS organization (
    id  int NOT NULL UNIQUE,
    name  text,
    picture  text,
    CONSTRAINT pk_organization PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS compensation (
    id serial NOT NULL UNIQUE,
    code text,
    currency text,
    min_amount REAL,
    max_amount REAL,
    periodicity text,
    CONSTRAINT pk_compensation PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS job (
    id serial NOT NULL UNIQUE,
    objective text,
    type text,
    organization INT,
    locations text[],
    remote BOOLEAN,
    external BOOLEAN,
    deadline timestamptz,
    status text,
    compensation int,
    CONSTRAINT pk_job PRIMARY KEY(id),
    CONSTRAINT fk_job_organization FOREIGN KEY(organization) REFERENCES organization(id),
    CONSTRAINT fk_job_compensation FOREIGN KEY(compensation) REFERENCES compensation(id)
);

CREATE TABLE IF NOT EXISTS skill_req (
    id  serial NOT NULL UNIQUE,
    name    text NOT NULL,
    experience  text,
    job_id int,
    CONSTRAINT pk_skill_req PRIMARY KEY(id),
    CONSTRAINT fk_skill_job FOREIGN KEY(job_id) REFERENCES job(id)
);

CREATE TABLE IF NOT EXISTS member (
    id serial NOT NULL UNIQUE,
    subject_id text,
    username text,
    prof_headline text,
    picture text,
    member BOOLEAN,
    manager BOOLEAN,
    poster BOOLEAN,
    weight BOOLEAN,
    CONSTRAINT pk_member PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS question (
    id text NOT NULL UNIQUE,
    text text NOT NULL,
    date timestamptz,
    job_id int,
    CONSTRAINT pk_question PRIMARY KEY(id),
    CONSTRAINT fk_question_job FOREIGN KEY(job_id) REFERENCES job(id)
);

CREATE TABLE IF NOT EXISTS job_to_member (
    id serial NOT NULL UNIQUE,
    job_id int NOT NULL,
    member_id int NOT NULL,
    CONSTRAINT pk_job_to_member PRIMARY KEY(id),
    CONSTRAINT fk_job_job_to_member FOREIGN KEY(job_id) REFERENCES job(id),
    CONSTRAINT fk_member_job_to_member FOREIGN KEY(member_id) REFERENCES member(id)
);

-- person data structure
CREATE TABLE IF NOT EXISTS compensation (
    id serial NOT NULL UNIQUE,
    amount real,
    currency text,
    periodicity text,
    CONSTRAINT pk_compensation PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS compensations (
    id serial NOT NULL UNIQUE,
    freelancer_id int,
    employee_id int,
    intern_id int,
    CONSTRAINT pk_compensations PRIMARY KEY(id),
    CONSTRAINT fk_free_comp FOREIGN KEY (freelancer_id) REFERENCES compensation(id),
    CONSTRAINT fk_empl_comp FOREIGN KEY (employee_id) REFERENCES compensation(id),
    CONSTRAINT fk_int_comp FOREIGN KEY (intern_id) REFERENCES compensation(id)
);

CREATE TABLE IF NOT EXISTS person (
    id serial NOT NULL UNIQUE,
    compensations_id int,
    location text,
    name text,
    picture text,
    prof_headline text,
    remoter boolean,
    subject_id text,
    username text,
    verified boolean,
    weight real,
    CONSTRAINT pk_person PRIMARY KEY (id),
    CONSTRAINT fk_compensation FOREIGN KEY (compensations_id) REFERENCES compensations(id)
);

CREATE TABLE IF NOT EXISTS open_to_options (
    id serial NOT NULL UNIQUE,
    name text,
    person_id int,
    CONSTRAINT pk_open_to_options PRIMARY KEY (id),
    CONSTRAINT fk_person_to_open FOREIGN KEY (person_id) REFERENCES person(id)
);

CREATE TABLE IF NOT EXISTS skills (
    id serial NOT NULL UNIQUE,
    name text,
    weight real,
    person_id int,
    CONSTRAINT pk_skills PRIMARY KEY (id),
    CONSTRAINT fk_person_to_skill FOREIGN KEY (person_id) REFERENCES person(id)
);

