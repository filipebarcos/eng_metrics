create table repositories(
    id serial primary key,
    name text not null,
    owner text not null,
    is_fork boolean not null default false,
    is_private boolean not null default true,
    is_archived boolean not null default false,
    is_template boolean not null default false,
    github_created_at timestamp with time zone not null default now(),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create unique index idx_repositories_owner_name on repositories(owner, name);
---- create above / drop below ----

drop table repositories;
drop index idx_repositories_owner_name;
