-- Write your migrate up statements here

alter table pull_requests add column deployed_at timestamp with time zone;

---- create above / drop below ----

alter table pull_requests drop column deployed_at;
