-- +migrate Up
create table if not exists competition_parameters (
  id text not null constraint competition_parameters_pk primary key,
  expected_participants integer not null,
  deployment_type text check(deploymentType in ("deployment", "points")) not null,
  group_hits int not null,
  elimination_hits int not null,
  qualification_based_on_rounds int not null,
  created_at datetime not null,
  updated_at datetime,
);

alter table
  competitions
add
  column competition_parameters_id text constraint competitions_competition_parameters_fk;

create table if not exists competition_group_rounds (
  id text not null constraint competition_group_rounds_pk primary key,
  competition_id text not null constraint competition_group_rounds_competitions_fk,
  number int not null,
  participants_starting_count int not null,
  number_of_groups int not null,
  participants_in_groups text not null,
  shift_criteria text not null,
  number_of_advancers int not null,
  created_at datetime not null,
  updated_at datetime,
);

-- +migrate Down
drop table if exists competition_parameters;

drop table if exists competition_group_rounds;

alter table
  competitions drop column competition_parameters_id;