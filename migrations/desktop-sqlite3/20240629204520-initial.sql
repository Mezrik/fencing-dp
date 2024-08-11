-- +migrate Up
create table competition_categories (
  id text not null constraint competition_categories_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  name text not null
);

create table weapons (
  id text not null constraint weapons_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  name text not null
);

create table competitions (
  id text not null constraint competitions_pk primary key,
  created_at datetime not null,
  updated_at datetime,
  category_id text not null,
  weapon_id text not null,
  name text not null,
  organizer_name text not null,
  federation_name text not null,
  competition_type text check(
    competition_type in ("national", "international")
  ) not null,
  gender text check(gender in ("male", "female", "mixed")) not null,
  date datetime not null,
  foreign key(weapon_id) references weapons(id),
  foreign key(category_id) references competition_categories(id)
);

-- +migrate Down
drop table competitions;

drop table competition_categories;

drop table weapons;