package cmd

import (
	"api-gateway/internal/app"
	"api-gateway/internal/config"
	"api-gateway/log"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	libmigrate "github.com/golang-migrate/migrate"
	libmigratemysql "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"
)

var (
	steps           int
	migrationsPath  string
	migrationsTable string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  "Run database migrations",
	Run:   migrate,
}

func init() {
	migrateCmd.Flags().StringVarP(
		&migrationsPath,
		"migrations-path",
		"m",
		"./migrations",
		"path to migrations directory",
	)

	migrateCmd.Flags().StringVarP(
		&migrationsTable,
		"migrations-table",
		"t",
		"migrations",
		"database table holding migrations",
	)

	migrateCmd.Flags().IntVarP(
		&steps,
		"steps",
		"n",
		0,
		"number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.",
	)
}

func migrate(_ *cobra.Command, _ []string) {
	if migrationsPath == "" {
		log.Fatal("migration path is required")
	}
	if !(strings.HasPrefix(migrationsPath, "/")) {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}
		migrationsPath, err = filepath.Abs(filepath.Join(path, migrationsPath))
		if err != nil {
			log.Fatal("cannot resolve full migration path")
		}
	}
	app.A.Ctx = context.Background()
	app.WithDatabase()
	cfg := &libmigratemysql.Config{
		MigrationsTable: migrationsTable,
	}
	driver, err := libmigratemysql.WithInstance(app.A.DB.DB, cfg)
	if err != nil {
		log.Fatal("error in instantiating the migration driver")
	}
	m, err := libmigrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		config.Cfg.Database.Name,
		driver,
	)
	if err != nil || m == nil {
		log.Fatal("error in instantiating the migration instance")
	}
	if steps == 0 {
		err = m.Up()
	} else {
		err = m.Steps(steps)
	}
	if err != nil {
		if errors.Is(err, libmigrate.ErrNoChange) {
			log.Info("No changes applied")
		} else {
			log.Fatal(fmt.Sprintf("error in running migrations: %s", err.Error()))
		}
	} else {
		log.Info("Migrations ran successfully")
	}
}
