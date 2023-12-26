-- create table to store the shortened urls
create table short_url (
                     id serial primary key,
                     url_code varchar(6) not null default '',
                     target_url text not null,
                     created_at timestamp not null default now()
);

-- add first ten rows so it is easier to know if it is encoding correctly
insert into short_url (url_code, target_url, created_at) values ('1', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('2', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('3', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('4', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('5', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('6', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('7', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('8', 'inserted by seed', now());
insert into short_url (url_code, target_url, created_at) values ('9', 'inserted by seed', now());
