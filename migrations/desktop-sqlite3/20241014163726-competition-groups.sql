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
  'state' text check(
    'state' in (
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
)
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

insert into
  match_states (id, 'state', created_at)
values
  (
    '2da77944-433c-4234-a3f6-5f35261b4450',
    'match_start',
    '2024-01-01 00:00:02'
  );

insert into
  match_states (id, 'state', created_at)
values
  (
    '1dc5cfef-d4fd-43c6-92de-88709b7d52be',
    'fight_start',
    '2024-01-01 00:00:02'
  );

insert into
  match_states (id, 'state', created_at)
values
  (
    'f8218de4-3c0d-4c8f-9644-53edc57f31a6',
    'fight_stop',
    '2024-01-01 00:00:20'
  );

insert into
  match_states (id, 'state', point_to, created_at)
values
  (
    '0fb65ed8-9e31-4693-90e7-bcfdedf77462',
    'point_added',
    'cf002140-487a-477b-8e65-7037e96abf9b',
    '2024-01-01 00:00:20'
  );

insert into
  match_states (id, 'state', created_at)
values
  (
    '2d4a6f55-8e7a-4c08-9fdb-92bb01fa309e',
    'fight_start',
    '2024-01-01 00:00:30'
  );

insert into
  match_states (id, 'state', created_at)
values
  (
    '9cb64fe6-dfd0-4963-b2ad-d7a28c060b21',
    'match_end',
    '2024-01-01 00:00:45'
  );

-- +migrate Down
drop table competition_groups;

drop table competition_matches;

drop table match_states;