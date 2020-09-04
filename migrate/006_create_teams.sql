-- Write your migrate up statements here

create table teams(
    id serial primary key,
    name text not null,
    members text[] default '{}',
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create unique index idx_teams_name on teams (name);
create index idx_teams_members on teams using GIN(members);

---- create above / drop below ----

drop index idx_teams_name;
drop index idx_teams_members;
drop table teams;
