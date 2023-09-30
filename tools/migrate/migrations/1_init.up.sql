create table ftable (
    id serial primary key,
    query varchar(255) not null,
    freq int default 1,
    unique (query)
);