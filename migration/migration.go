package migration

import (
	"context"
	"fmt"
	"github.com/huuthuan-nguyen/wager/app/model"
	"github.com/huuthuan-nguyen/wager/config"
	"github.com/uptrace/bun"
	"log"
	"strings"
)

type Migrator interface {
	Up() (err error)
	Down() (err error)
	GetVersion() string
	SetDB(db *bun.DB)
	GetDB() *bun.DB
}

var Migrators = []Migrator{
	&InitSchema{},
}

// Migration /**
type Migration struct {
	bun.BaseModel `bun:"table:migrations,alias:migration"`
	ID            int     `bun:"id,pk,autoincrement"`
	Migration     string  `bun:"migration"`
	Batch         int     `bun:"batch"`
	DB            *bun.DB `bun:"-"`
}

// SetDB /**
func (m *Migration) SetDB(db *bun.DB) {
	m.DB = db
}

// GetDB /**
func (m *Migration) GetDB() *bun.DB {
	return m.DB
}

// Engine /**
type Engine struct {
	DB *bun.DB
}

// NewEngine /**
func NewEngine(c *config.Config) *Engine {
	const migration = `CREATE TABLE IF NOT EXISTS migrations(
		id SERIAL PRIMARY KEY,
		migration VARCHAR(255) NOT NULL,
		batch INT NOT NULL)`

	db, err := model.NewConnection(c)
	if err != nil {
		log.Fatalln("Error connecting to database...")
	}

	// init the "migration" table
	_, err = db.Exec(migration)

	if err != nil {
		log.Fatalf("Error while creating migrations table:%v\n", err)
	}

	return &Engine{
		DB: db,
	}
}

// Migrate /**
func (engine *Engine) Migrate() (err error) {
	lastBatch, err := engine.getLastBatch()

	if err != nil {
		return err
	}

	successMigration := 0

	fmt.Println("Begin migrating...")

	for _, migrator := range Migrators {
		migrator.SetDB(engine.DB)

		version := migrator.GetVersion()

		migration, err := engine.getMigrationByVersion(version)
		if err != nil { // not migrated yet
			migration = Migration{
				Migration: version,
				Batch:     lastBatch + 1,
				DB:        engine.DB,
			}
		} else { // already migrated
			continue
		}

		// "up" the migration changes
		fmt.Println("Migrating ", migrator.GetVersion())
		err = migrator.Up()

		if err != nil {
			fmt.Printf("Error when migrating %s: %v\n", migrator.GetVersion(), err)
			panic(err)
		}

		// insert a new record to migration table
		err = migration.Create()
		if err != nil {
			panic(err)
		}
		successMigration++
		fmt.Println("Migrated ", migrator.GetVersion())
	}

	if successMigration > 0 {
		fmt.Println("Migrate Done.")
	} else {
		fmt.Println("Nothing to migrate.")
	}

	return err
}

// Rollback /**
func (engine *Engine) Rollback() (err error) {
	lastBatch, err := engine.getLastBatch()
	migrations, err := engine.getMigrationByBatch(lastBatch)

	successRollback := 0

	fmt.Println("Begin rolling back...")
	for _, m := range migrations {
		m.SetDB(engine.DB)
		fmt.Println("Rolling back ", m.Migration)
		err = m.Rollback()
		if err != nil {
			return err
		}
		err = m.Delete()
		if err != nil {
			return err
		}
		fmt.Println("Rolled back ", m.Migration)
		successRollback++
	}

	if successRollback > 0 {
		fmt.Println("Rollback Done.")
	} else {
		fmt.Println("Nothing to rollback.")
	}
	return
}

/**
Get the latest batch on migration table.
*/
func (engine *Engine) getLastBatch() (maxBatch int, err error) {
	statement := "SELECT COALESCE(MAX(batch), 0) AS b FROM migrations"
	err = engine.DB.QueryRow(statement).Scan(&maxBatch)
	return maxBatch, err
}

// Create /**
func (m *Migration) Create() (err error) {
	statement := "INSERT INTO migrations(migration, batch) VALUES(?, ?)"

	_, err = m.DB.Exec(statement, m.Migration, m.Batch)
	return err
}

// Delete /**
func (m *Migration) Delete() (err error) {
	_, err = m.DB.Exec("DELETE FROM migrations WHERE id=?", m.ID)
	return err
}

// Rollback /**
func (m *Migration) Rollback() (err error) {
	for _, migrator := range Migrators {
		migrator.SetDB(m.DB)
		if m.Migration == migrator.GetVersion() {
			err = migrator.Down()
		}
	}
	return err
}

/**
Get migrations by batch number
*/
func (engine *Engine) getMigrationByBatch(batch int) (migrations []Migration, err error) {
	err = engine.DB.NewSelect().
		Model(&migrations).
		Column("id", "migration", "batch").
		Where("batch=?", batch).
		Order("id DESC").
		Scan(context.Background())
	return migrations, err
}

/**
Get specific migration by version
*/
func (engine *Engine) getMigrationByVersion(version string) (migration Migration, err error) {
	err = engine.DB.NewSelect().
		Model(&migration).
		Column("id", "migration", "batch").
		Where("migration=?", version).
		Scan(context.Background())
	migration.SetDB(engine.DB)
	return migration, err
}

// Reset /**
func (engine *Engine) Reset() (err error) {
	ctx := context.Background()
	err = model.WithTransaction(ctx, engine.DB, func(tx bun.Tx) error {
		fmt.Println("Begin drop all tables...")
		if _, err := tx.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS = 0"); err != nil {
			return err
		}

		rows, err := tx.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = (SELECT DATABASE()) AND table_name <> 'migrations'")
		if err != nil {
			return err
		}

		defer rows.Close()

		tableNames := make([]string, 0)
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				return err
			} else {
				tableNames = append(tableNames, tableName)
			}
		}
		if len(tableNames) > 0 {
			if _, err := tx.Exec(fmt.Sprintf("DROP TABLES IF EXISTS %s", strings.Join(tableNames, ", "))); err != nil {
				return err
			}
		}

		fmt.Println("Drop all tables done.")

		// reset the "migrations" table
		if _, err := tx.Exec("TRUNCATE migrations"); err != nil {
			return err
		}

		if _, err := engine.DB.Exec("SET FOREIGN_KEY_CHECKS = 1"); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = engine.Migrate()
	return err
}
