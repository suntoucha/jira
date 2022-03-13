module github.com/suntoucha/jira/pgexporter

go 1.17

replace github.com/suntoucha/jira => ../

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.4
	github.com/suntoucha/jira v0.0.0-20220201183802-99a1244276e7
)
