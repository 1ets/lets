package types

type SqLite struct {
	Debug           bool
	DBPath          string
	Repositories    []IMySQLRepository
	EnableMigration bool

	QueryFields              bool
	DisableNestedTransaction bool

	dsn string
}
