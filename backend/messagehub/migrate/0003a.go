package migrate

import (
	migrate "github.com/rubenv/sql-migrate"
)

func m0003a() *migrate.Migration {
	return &migrate.Migration{
		Id: "0003a",
		Up: []string{
			`
create table settings
(
	id text not null constraint settings_pk primary key,
	value text not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);
`,
		},
		Down: []string{},
	}
}
