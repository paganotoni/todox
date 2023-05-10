package database

import _ "embed"

//go:embed schema.sql
var sql string

func Migrate() error {
	conn, err := connection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(sql)
	return err
}
