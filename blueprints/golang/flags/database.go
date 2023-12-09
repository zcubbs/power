package flags

import (
	"fmt"
	"strings"
)

type Database string

const (
	MySql    Database = "mysql"
	Postgres Database = "postgres"
	Sqlite   Database = "sqlite"
	Mongo    Database = "mongo"
	None     Database = "none"
)

var AllowedDBDrivers = []string{string(MySql), string(Postgres), string(Sqlite), string(Mongo), string(None)}

func (f *Database) String() string {
	return string(*f)
}

func (f *Database) Type() string {
	return "Database"
}

func (f *Database) Set(value string) error {
	for _, database := range AllowedDBDrivers {
		if database == value {
			*f = Database(value)
			return nil
		}
	}

	return fmt.Errorf("database to use. Allowed values: %s", strings.Join(AllowedDBDrivers, ", "))
}
