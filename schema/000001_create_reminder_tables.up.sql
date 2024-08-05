CREATE TABLE reminds
(
    id serial not null unique,
    title varchar(255) not null,
    msg varchar(255) not null,
    remind_date TIMESTAMP not null
);

	