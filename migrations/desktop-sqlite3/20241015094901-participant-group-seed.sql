-- +migrate Up
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
    '1949e723-c162-45ee-b3ab-70614100a9ab',
    'Group 2',
    2,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
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
    'c3f0817e-fda4-4623-b3e7-93146fa54c30',
    'Group 3',
    3,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
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
    'a177c366-e24f-455c-b2fb-44808b7ba6c7',
    'Group 4',
    4,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
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
    '6f879516-2623-4949-a1d2-10fd0f9e8848',
    'Group 5',
    5,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
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
    '312d982b-145e-48af-893b-8aaa27203a82',
    'Group 6',
    6,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
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
    'd45f7c46-afd3-45a5-a3fe-937811de6df2',
    'Group 7',
    7,
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '2024-01-01 00:00:00'
  );

insert into
  participating_competitors (
    competition_id,
    competitor_id,
    group_id,
    deployment_number,
    starting_position
  )
values
  (
    'eebc1dc6-d3a8-460d-a23f-752692b88eca',
    '8318ffe4-9002-4fc4-a23e-804b9361078f',
    '1949e723-c162-45ee-b3ab-70614100a9ab',
    3,
    4
  );

-- +migrate Down