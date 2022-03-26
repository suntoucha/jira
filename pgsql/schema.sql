drop table if exists issue_raw;
create table issue_raw (
	key varchar primary key,
	raw varchar not null
);

drop table if exists issue;
create table issue_flat (
	key varchar primary key,
	project varchar not null,
	description varchar not null,
	summary varchar not null,
	type varchar not null,
	type_name varchar not null,
	is_subtask boolean not null,
	status varchar not null,
	status_name varchar not null,
	assignee varchar not null default '',
	reporter varchar not null default '',
	dt_created timestamp not null,
	dt_updated timestamp null,
	dt_resolution timestamp null
);