package postgres

import (
	"fmt"

	"go-blog/internal/config"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/model/user"
	"go-blog/pkg/db"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func (d *Store) Transaction() *transaction {
	return newTransaction(d)
}

func (d *Store) Articles() *articles {
	return newArticles(d)
}

func (d *Store) ArticleBodys() *articleBodys {
	return newArticleBodys(d)
}

func (d *Store) ArticleTags() *articleTags {
	return newArticleTags(d)
}

func (d *Store) Categories() *categories {
	return newArticleCategories(d)
}

func (d *Store) ArticleSeries() *articleSeries {
	return newArticleSeries(d)
}

func (d *Store) ArticleArticleTag() *articleArticleTags {
	return newArticleArticleTags(d)
}

func (d *Store) Users() *users {
	return newUsers(d)
}

func (d *Store) Roles() *roles {
	return newRoles(d)
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
	// if err := db.Migrator().DropTable(&user.Role{}); err != nil {
	// 	return errors.Wrap(err, "drop role model failed")
	// }
	// if err := db.Migrator().DropTable(&user.User{}); err != nil {
	// 	return errors.Wrap(err, "drop user table failed")
	// }
	// if err := db.Migrator().DropTable(&article.Article{}); err != nil {
	// 	return errors.Wrap(err, "drop article model failed")
	// }
	// if err := db.Migrator().DropTable(&article.ArticleBody{}); err != nil {
	// 	return errors.Wrap(err, "drop articleBody model failed")
	// }
	// if err := db.Migrator().DropTable(&article.Tag{}); err != nil {
	// 	return errors.Wrap(err, "drop article tag model failed")
	// }
	// if err := db.Migrator().DropTable(&article.Category{}); err != nil {
	// 	return errors.Wrap(err, "drop article category model failed")
	// }
	// if err := db.Migrator().DropTable(&article.Series{}); err != nil {
	// 	return errors.Wrap(err, "drop article series model failed")
	// }
	// if err := db.Migrator().DropTable(&article.ArticleTag{}); err != nil {
	// 	return errors.Wrap(err, "drop article article_tag model failed")
	// }
	// if err := db.Migrator().DropTable(&list.Category{}); err != nil {
	// 	return errors.Wrap(err, "drop list category model failed")
	// }

	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.Role{}); err != nil {
		return errors.Wrap(err, "migrate role model failed")
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}
	if err := db.AutoMigrate(&article.Tag{}); err != nil {
		return errors.Wrap(err, "migrate article tag model failed")
	}
	if err := db.AutoMigrate(&category.Category{}); err != nil {
		return errors.Wrap(err, "migrate article category model failed")
	}
	if err := db.AutoMigrate(&article.ArticleBody{}); err != nil {
		return errors.Wrap(err, "migrate articleBody model failed")
	}
	if err := db.AutoMigrate(&article.Series{}); err != nil {
		return errors.Wrap(err, "migrate article series model failed")
	}
	if err := db.AutoMigrate(&article.Article{}); err != nil {
		return errors.Wrap(err, "migrate article model failed")
	}
	if err := db.AutoMigrate(&article.ArticleTag{}); err != nil {
		return errors.Wrap(err, "migrate article article_tag model failed")
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
