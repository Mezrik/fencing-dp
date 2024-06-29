-- +migrate Up
create table competition_categories (
  id BLOB not null constraint competition_categories_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  name text not null
);

create table weapons (
  id BLOB not null constraint weapons_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  name text not null
);

create table competitions (
  id BLOB not null constraint competitions_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  category_id BLOB not null,
  weapon_id BLOB not null,
  name text not null,
  organizer_name text not null,
  federation_name text not null,
  competitionType integer check(competitionType in (0, 1)) not null,
  gender integer check(gender in (0, 1, 2)) not null,
  date datetime not null,
  foreign key(weapon_id) references weapons(id),
  foreign key(category_id) references competition_categories(id)
);

-- +migrate Down
drop table competitions;

drop table competition_categories;

drop table weapons;