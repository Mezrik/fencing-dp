-- +migrate Up
create table if not exists clubs (
  id text not null constraint clubs_pk primary key,
  name text not NULL,
  created_at datetime not null,
  updated_at datetime
);

create table if not exists competitors (
  id text not null constraint competitors_pk primary key,
  surname text not NULL,
  firstname text not NULL,
  created_at datetime not null,
  updated_at datetime,
  gender text check(gender in ("male", "female", "mixed", "unset")) not null,
  club_id text constraint competitors_clubs_fk,
  license text,
  license_fie text,
  birthdate date
);

create table if not exists participating_competitors (
  competition_id text not null constraint participating_competitors_competitions_fk,
  competitor_id text not null constraint participating_competitors_competitors_fk,
  deployment_number int not null,
  points real default 0.0,
  starting_position int,
  primary key (competition_id, competitor_id)
);

insert into
  weapons (id, name, created_at)
values
  (
    '4b0dec55-2a2b-46f7-8fd6-c001761148da',
    'Weapon 1',
    '2022-01-01 00:00:00'
  );

insert into
  competition_categories (id, name, created_at)
values
  (
    '9ecf8eb0-4101-40fd-b46e-41d9cf39fc22',
    'Category 1',
    '2022-01-01 00:00:00'
  );

insert into
  competitions (
    id,
    created_at,
    category_id,
    weapon_id,
    name,
    organizer_name,
    federation_name,
    competition_type,
    gender,
    date
  )
values
  (
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2022-01-01 00:00:00',
    '9ecf8eb0-4101-40fd-b46e-41d9cf39fc22',
    '4b0dec55-2a2b-46f7-8fd6-c001761148da',
    'Competition 1',
    'Organizer 1',
    'Federation 1',
    'national',
    'mixed',
    '2022-01-01 00:00:00'
  );

insert into
  clubs (id, name, created_at)
values
  (
    '303235ea-0438-4a05-8bbc-557b730c83a9',
    'Club 1',
    '2022-01-01 00:00:00'
  );

insert into
  competitors (
    id,
    surname,
    firstname,
    created_at,
    gender,
    club_id,
    license,
    birthdate
  )
values
  (
    'cf002140-487a-477b-8e65-7037e96abf9b',
    'Doe',
    'John',
    '2022-01-01 00:00:00',
    'male',
    '303235ea-0438-4a05-8bbc-557b730c83a9',
    '123456789',
    '1990-01-01'
  );

insert into
  participating_competitors (
    competition_id,
    competitor_id,
    deployment_number,
    starting_position
  )
values
  (
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    'cf002140-487a-477b-8e65-7037e96abf9b',
    1,
    1
  );

-- +migrate Down
drop table competitors;

drop table clubs;

drop table participating_competitors;