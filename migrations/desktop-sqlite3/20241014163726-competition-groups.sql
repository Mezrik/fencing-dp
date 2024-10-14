-- +migrate Up
create table if not exists competition_groups (
  id text not null constraint competition_groups_pk primary key,
  name text not NULL,
  piste_number int not null,
  competition_id text not null constraint competition_groups_competitions_fk,
  created_at datetime not null,
  updated_at datetime
);

create table if not exists competition_matches (
  id text not null constraint competition_matches_pk primary key,
  competition_group_id text not null constraint competition_matches_competition_groups_fk,
  participant_one_id text not null constraint competition_matches_competitors_fk,
  participant_two_id text not null constraint competition_matches_competitors_fk,
  created_at datetime not null,
  updated_at datetime
);

create table if not exists match_states (
  id text not null constraint match_states_pk primary key,
  change text check(
    change in (
      'match_start',
      'fight_start',
      'fight_stop',
      'point_added',
      'point_subtracted',
      'match_end'
    )
  ) not null,
  point_to int,
  created_at datetime not null,
  updated_at datetime
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
    '024fe93c-467a-4b24-a256-d5e412bb0cac',
    2,
    2
  );

insert into
  competition_groups (
    id,
    name,
    piste_number,
    competition_id,
    created_at
  )
values
  (
    '94ddd19e-4bc6-4c77-a5f0-2ac7ae6e0026',
    'Group 1',
    1,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
  );

insert into
  competition_matches (
    id,
    competition_group_id,
    participant_one_id,
    participant_two_id,
    created_at
  )
values
  (
    '5515499d-959d-460f-a6f3-0dbcd877e5e9',
    '94ddd19e-4bc6-4c77-a5f0-2ac7ae6e0026',
    'cf002140-487a-477b-8e65-7037e96abf9b',
    '024fe93c-467a-4b24-a256-d5e412bb0cac',
    '2024-01-01 00:00:01'
  );

-- +migrate Down
drop table competition_groups;

drop table competition_matches;

drop table match_states;