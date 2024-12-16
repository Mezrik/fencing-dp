-- +migrate Up
alter table
  competitors
add
  column rating int;

update
  competitors
set
  rating = 0;

-- +migrate Down
alter table
  competitors drop column rating;