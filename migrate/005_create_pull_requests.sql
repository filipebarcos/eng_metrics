create table pull_requests(
    id serial primary key,
    number integer not null,
    url text not null,
    title text not null,
    milestone text,
    draft boolean not null default false,
    state text,
    review_decision text,
    author text,
    merged_by text,
    labels_count integer not null default 0,
    labels text[] default '{}',
    repository_id integer not null,
    github_closed_at timestamp with time zone null,
    github_merged_at timestamp with time zone null,
    github_published_at timestamp with time zone null,
    github_created_at timestamp with time zone not null,
    github_updated_at timestamp with time zone not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create unique index idx_pull_requests_repository_id_number on pull_requests(repository_id, number);
---- create above / drop below ----

drop index idx_pull_requests_repository_id_number;
drop table pull_requests;
