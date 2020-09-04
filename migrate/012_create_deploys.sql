create table deploys(
    id serial primary key,
    created_at timestamp with time zone not null default now()
);

---- create above / drop below ----

drop table deploys;
