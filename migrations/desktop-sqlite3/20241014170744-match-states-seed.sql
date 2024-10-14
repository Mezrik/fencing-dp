-- +migrate Up
insert into
  match_states (id, change, created_at)
values
  (
    '2da77944-433c-4234-a3f6-5f35261b4450',
    'match_start',
    '2024-01-01 00:00:02'
  );

insert into
  match_states (id, change, created_at)
values
  (
    '1dc5cfef-d4fd-43c6-92de-88709b7d52be',
    'fight_start',
    '2024-01-01 00:00:02'
  );

insert into
  match_states (id, change, created_at)
values
  (
    'f8218de4-3c0d-4c8f-9644-53edc57f31a6',
    'fight_stop',
    '2024-01-01 00:00:20'
  );

insert into
  match_states (id, change, point_to, created_at)
values
  (
    '0fb65ed8-9e31-4693-90e7-bcfdedf77462',
    'point_added',
    'cf002140-487a-477b-8e65-7037e96abf9b',
    '2024-01-01 00:00:20'
  );

insert into
  match_states (id, change, created_at)
values
  (
    '2d4a6f55-8e7a-4c08-9fdb-92bb01fa309e',
    'fight_start',
    '2024-01-01 00:00:30'
  );

insert into
  match_states (id, change, created_at)
values
  (
    '9cb64fe6-dfd0-4963-b2ad-d7a28c060b21',
    'match_end',
    '2024-01-01 00:00:45'
  );

-- +migrate Down