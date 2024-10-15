-- +migrate Up
ALTER TABLE
  participating_competitors
ADD
  COLUMN group_id text;

UPDATE
  participating_competitors
SET
  group_id = '94ddd19e-4bc6-4c77-a5f0-2ac7ae6e0026';

-- +migrate Down
ALTER TABLE
  participating_competitors DROP COLUMN group_id;