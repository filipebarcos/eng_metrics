-- Write your migrate up statements here

alter table repositories add column url text not null default 'TO_BE_REPLACED';

---- create above / drop below ----

alter table repositories drop column url;
