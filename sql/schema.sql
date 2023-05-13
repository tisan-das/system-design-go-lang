
create schema if not exists sample;
create table sample.grocery(
    id int unique not null,
    name text not null,
    quantity int,
    primary key(id)
);