package postgres

import (
	"fmt"

	"go-blog/internal/config"
	"go-blog/internal/db/model/tag"
	"go-blog/internal/db/model/todo"
	"go-blog/internal/db/model/user"
	"go-blog/pkg/db"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type pgOptions struct {
	username string
	database string
	funcList []db.OptionFunc
}

func GetPGFactory() (*Store, error) {
	opts := &pgOptions{username: config.PgConfig.Username,
		database: config.PgConfig.Database,
		funcList: []db.OptionFunc{db.WithHost(config.PgConfig.Host), db.WithPort(config.PgConfig.Port), db.WithPassword(config.PgConfig.Password)}}
	if opts == nil {
		return nil, fmt.Errorf("failed to get postgresql store factory")
	}
	var err error
	var dbIns *gorm.DB
	if err != nil {
		return nil, err
	}
	var pgFactory *Store
	dbIns, err = db.NewPostgresql(opts.username, opts.database, opts.funcList...)
	pgFactory = &Store{dbIns}

	if pgFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get postgresql store factory, pgFactory: %+v, error: %w", pgFactory, err)
	}
	// 自动化迁移
	if err := migrateDatabase(dbIns); err != nil {
		fmt.Println(err)
	}

	return pgFactory, nil
}

// cleanDatabase tear downs the database tables.
func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&user.User{}); err != nil {
		return errors.Wrap(err, "drop user table failed")
	}

	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}
	if err := db.AutoMigrate(&tag.Tag{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}

	return nil
}

// resetDatabase resets the database tables.
func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}
