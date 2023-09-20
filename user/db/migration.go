package db

import (
	"embed"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/golang-migrate/migrate/v4/source/httpfs"
	"io"
	"net/http"
	"strings"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
	//go:embed seeds/*.sql
	seeds embed.FS
)

type driver struct {
	httpfs.PartialDriver
}

func (d *driver) Open(rawURL string) (source.Driver, error) {
	err := d.PartialDriver.Init(http.FS(migrations), "migrations")
	if err != nil {
		return nil, err
	}

	return d, nil
}
func getMigration() (*migrate.Migrate, error) {
	cfg := config.LoadConfigOrPanic()
	database := db.NewDatabase(cfg.DBConfig)
	sqldb, err := database.DB.DB()
	if err != nil {
		return nil, err
	}
	dbdriver, err := mysql.WithInstance(sqldb, &mysql.Config{})
	// log files in migrations folder
	files, err := migrations.ReadDir("migrations")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		println(file.Name())
	}

	source.Register("embed", &driver{})

	return migrate.NewWithDatabaseInstance("embed://", cfg.DBConfig.DataBase, dbdriver)
}

func GetAndRunSeeds() error {
	cfg := config.LoadConfigOrPanic()
	database := db.NewDatabase(cfg.DBConfig)
	sqldb, err := database.DB.DB()
	if err != nil {
		return err
	}

	// log files in migrations folder
	files, err := seeds.ReadDir("seeds")
	if err != nil {
		return err
	}
	for _, file := range files {
		println(file.Name())
	}

	source.Register("embed", &driver{})
	for _, file := range files {
		// open file and read contents
		f, err := seeds.Open("seeds/" + file.Name())
		if err != nil {
			return err
		}
		fileContents, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		// split file contents by new line
		// for each line, execute query
		lines := strings.Split(string(fileContents), "\n")
		for _, line := range lines {
			if line != "" {
				_, err = sqldb.Exec(line)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func MigrateUp() error {
	m, err := getMigration()
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func MigrateDown() error {
	m, err := getMigration()
	if err != nil {
		return err
	}

	return m.Down()
}
