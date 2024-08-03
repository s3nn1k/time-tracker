CREATE TABLE users (
    id serial primary key unique not null,
    name varchar(255) unique not null,
    password_hash varchar(255) not null
);

CREATE TABLE timers (
    id serial primary key unique not null,
    user_id int references users(id) on delete cascade not null,
    name varchar(255) not null,
    start_time bigint not null,
    work_time bigint not null
)