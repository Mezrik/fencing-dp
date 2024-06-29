
-- +migrate Up
create table competition_categories (
  id         BLOB not null
    constraint competition_categories_pk
        primary key,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  name       text
);
create table competitions (
  id         BLOB not null
    constraint competitions_pk
        primary key,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  category_id BLOB,
  name       text,
  foreign key(category_id) references competition_categories(id)
);

-- +migrate Down
drop table competitions;
drop table competition_categories;
