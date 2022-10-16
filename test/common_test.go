package test

import (
	"database/sql"
	"fmt"
	"github.com/gGerret/otus-social-prj/repository"
	_ "github.com/go-sql-driver/mysql"
)

func InitDbTest() *sql.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?parseTime=true",
		"social_svc",
		"social_sql_passw0rd",
		"tcp",
		"localhost",
		13306,
		"social",
	)

	db, err := sql.Open(repository.DbDriver, connectionString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Test DB with error: %s", err.Error()))
	}

	return db
}
