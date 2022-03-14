drop table if exists issue;

create table issue (
	key varchar primary key,
	project_key varchar not null,
	description varchar not null default '',
	summary varchar not null default '',
	type_id varchar not null,
	type_name varchar not null,
	is_subtask bool not null,
	status_id varchar not null,
	status_name varchar not null,
	assignee_email varchar not null default '',
	reporter_email varchar not null default '',
	dt_created timestamp not null,
	dt_updated timestamp null,
	dt_resolution timestamp not null,
	raw varchar not null
);