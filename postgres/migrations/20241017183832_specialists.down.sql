-- Drop unique index on specialists(nickname, deleted_at)
drop index if exists idx_specialists_nickname_deleted;

-- Drop services table
drop table if exists services;

-- Drop service_groups table
drop table if exists service_groups;

-- Drop specialists table
drop table if exists specialists;
