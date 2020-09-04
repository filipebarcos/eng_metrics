alter table issues rename column closed_at to github_closed_at;

---- create above / drop below ----

alter table issues rename column github_closed_at to closed_at;
