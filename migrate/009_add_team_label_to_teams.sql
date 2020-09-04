-- Write your migrate up statements here

alter table teams add column team_label text;
---- create above / drop below ----

alter table teams drop column team_label;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
