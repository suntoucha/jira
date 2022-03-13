drop table if exists issue;

create table issue (
	key varchar primary key,
	project_key varchar not null,
	raw varchar not null
);