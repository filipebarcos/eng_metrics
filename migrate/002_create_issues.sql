create table issues(
    id serial primary key,
    number integer not null,
    url text not null,
    title text not null,
    milestone text,
    labels_count integer not null default 0,
    labels text[] default '{}',
    assignees_count integer not null default 0,
    assignees text[] default '{}',
    repository_id integer not null,
    closed_at timestamp with time zone null,
    github_created_at timestamp with time zone not null,
    github_updated_at timestamp with time zone not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create unique index idx_issues_repository_id_number on issues(repository_id, number);
---- create above / drop below ----

drop table issues;
drop index idx_issues_repository_id_number;
