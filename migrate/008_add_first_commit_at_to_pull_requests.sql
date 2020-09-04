-- Write your migrate up statements here

alter table pull_requests add column first_commit_at timestamp with time zone;

---- create above / drop below ----

alter table pull_requests drop column first_commit_at;
