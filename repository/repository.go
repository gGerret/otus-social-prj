package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DbDriver = "mysql"
)

var database *sql.DB

type BaseRepository struct {
	db *sql.DB
}

func New() *BaseRepository {
	return &BaseRepository{db: database}
}

func NewDB(db *sql.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func InitDb(cfg *ConfigDb) *sql.DB {
	connectionString := fmt.Sprintf(
		"username=%s:%s@%s(%s:%d)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Net,
		cfg.Hostname,
		cfg.Port,
		cfg.Database,
	)

	db, err := sql.Open(DbDriver, connectionString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB with error: %s", err.Error()))
	}

	database = db
	return database
}

func (r *BaseRepository) GetDb() *sql.DB {
	return r.db
}

/*func BeginTX() (*sql.Tx, error) {
	return database.BeginTx()
}*/

//Only for tests
func SetMainDb(db *sql.DB) {
	database = db
}
