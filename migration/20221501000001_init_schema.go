package migration

import (
	"context"
	"github.com/huuthuan-nguyen/wager/app/model"
	"github.com/uptrace/bun"
)

type InitSchema struct {
	Migration
	Version string
}

func (m *InitSchema) Up() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {

		// "wagers" table
		const wager = `CREATE TABLE IF NOT EXISTS wagers(
			id INT(11) AUTO_INCREMENT PRIMARY KEY,
			total_wager_value INT UNSIGNED NOT NULL,
			odds INT UNSIGNED NOT NULL,
			selling_percentage INT UNSIGNED NOT NULL,
			selling_price DECIMAL(11,2) UNSIGNED NOT NULL,
			current_selling_price DECIMAL(11,2) UNSIGNED NOT NULL,
			percentage_sold INT UNSIGNED NULL,
			amount_sold DECIMAL(11,2) UNSIGNED NULL,
			placed_at TIMESTAMP NOT NULL
		) ENGINE=INNODB`
		if _, err = tx.Exec(wager); err != nil {
			return err
		}

		// "purchases" table
		const purchase = `CREATE TABLE IF NOT EXISTS purchases(
			id INT(11) AUTO_INCREMENT PRIMARY KEY,
			wager_id INT(11) NOT NULL,
			buying_price DECIMAL(11,2) UNSIGNED NOT NULL,
			bought_at TIMESTAMP NOT NULL,
			FOREIGN KEY (wager_id) REFERENCES wagers(id) ON DELETE RESTRICT ON UPDATE RESTRICT
		) ENGINE=INNODB`
		if _, err = tx.Exec(purchase); err != nil {
			return err
		}

		return nil
	})
}

func (m *InitSchema) Down() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS purchases, wagers`)
		return err
	})
}

func (m *InitSchema) GetVersion() string {
	return "20221501000001"
}
