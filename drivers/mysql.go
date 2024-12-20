package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/1ets/lets"
	"github.com/1ets/lets/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MySQLConfig []types.IMySQL

type mysqlProvider struct {
	// debug bool
	// DSN   string
	Gorm   *gorm.DB
	Sql    *sql.DB
	Config types.IMySQL
}

func (m *mysqlProvider) Connect() {
	var logType logger.Interface = logger.Default.LogMode(logger.Warn)
	if m.Config.DebugMode() {
		logType = logger.Default.LogMode(logger.Info)
	}

	var err error
	m.Gorm, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Config.GetDsn(), // data source name
		DefaultStringSize:         256,               // default size for string fields
		DisableDatetimePrecision:  true,              // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,             // auto configure based on currently MySQL configs
	}), &gorm.Config{
		Logger:      logType,
		QueryFields: m.Config.GetQueryFields(),
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		},
		PrepareStmt:              false,
		DisableNestedTransaction: m.Config.GetDisableNestedTransaction(),
	})

	if err != nil {
		lets.LogE(err.Error())
		time.Sleep(time.Second * 3)
		m.Connect()

		return
	}

	m.Sql, err = m.Gorm.DB()
	if err != nil {
		time.Sleep(time.Second * 3)
		m.Connect()

		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	m.Sql.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	m.Sql.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	m.Sql.SetConnMaxLifetime(time.Minute * 3)
}

// Define MySQL service host and port
func MySQL() {
	if MySQLConfig == nil {
		return
	}

	lets.LogI("MySQL Client Starting ...")

	for _, config := range MySQLConfig {
		mySQL := mysqlProvider{
			Config: config,
		}
		mySQL.Connect()

		// Inject Gorm into repository
		for _, repository := range config.GetRepositories() {
			repository.SetDriver(mySQL.Gorm)
		}

		// Migration
		if config.Migration() {
			err := mySQL.Gorm.AutoMigrate(&migration{})
			if err != nil {
				lets.LogE("Unable to run migration %w", err)
				return
			}
			Migrate(mySQL.Gorm, mySQL.Sql)
		}
	}
}

type migration struct {
	ID        uint   `gorm:"primaryKey,column:id"`
	Migration string `gorm:"column:migration"`
	Batch     uint   `gorm:"column:batch"`
}

func Migrate(g *gorm.DB, db *sql.DB) {
	// Define batch number
	var batch uint = 1
	lastMigration := &migration{}
	result := g.Last(lastMigration)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		lets.LogE("Unable to run migration %w", result.Error)
		return
	}

	batch = lastMigration.Batch + 1

	// Get migration files
	files, err := os.ReadDir("migrations")
	if err != nil {
		lets.LogE(err.Error())
		time.Sleep(time.Second * 3)
		Migrate(g, db)
		return
	}

	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		// Search migration
		search := &migration{
			Migration: name,
		}

		result := g.Where("migration = ?", name).First(search)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			lets.LogE("Unable to run migration %w", result.Error)
			return
		}

		// Execute
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			lets.LogI("Migrating: %s", name)

			// Read file content
			filePath := fmt.Sprintf("migrations/%s", file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				lets.LogE("Unable to run migration: %s", err.Error())
				return
			}

			err = g.Transaction(func(tx *gorm.DB) error {
				for _, query := range strings.Split(string(content), ";") {
					query := strings.TrimSpace(query)
					if query == "" {
						continue
					}

					result = g.Exec(query)
					if result.Error != nil {
						return result.Error
					}
				}

				return nil
			})

			if err != nil {
				lets.LogE("Unable to run migration %w", err.Error())
				return
			}

			// Insert migration record
			m := &migration{
				Migration: name,
				Batch:     batch,
			}

			result = g.Create(m)
			if result.Error != nil {
				lets.LogE("Unable to run migration: %s", result.Error.Error())
				return
			}
		}

	}
}
