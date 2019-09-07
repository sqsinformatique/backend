package db

import (
	"fmt"
	"os"
	"time"

	"github.com/sqsinformatique/backend/utils"

	"github.com/gobuffalo/packr"
	"github.com/olekukonko/tablewriter"
	migrate "github.com/rubenv/sql-migrate"
)

const dialect = "postgres"

var migrations *migrate.PackrMigrationSource

type statusRow struct {
	Id        string
	Migrated  bool
	AppliedAt time.Time
}

func initMigration() {
	if migrations != nil {
		return
	}
	migrations = &migrate.PackrMigrationSource{
		Box: packr.NewBox("./sql"),
	}
}

func MigrateSQL() (err error) {
	utils.Info("Begin migrate")

	err = InitDB()
	if err != nil {
		return
	}

	initMigration()

	migrate.SetTable("migrations")
	_, err = migrate.Exec(db, dialect, migrations, migrate.Up)
	if err != nil {
		return
	}

	err = MigrateSQLStatus()
	if err != nil {
		return
	}

	utils.Info("Migrate successfully completed")
	return
}

func MigrateSQLStatus() (err error) {
	if db == nil {
		err = InitDB()
		if err != nil {
			return
		}
	}

	initMigration()

	records, err := migrate.GetMigrationRecords(db, dialect)
	if err != nil {
		utils.Error(err.Error())
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Migration", "Applied"})
	table.SetColWidth(60)

	rows := make(map[string]*statusRow)
	ms, _ := migrations.FindMigrations()

	for _, m := range ms {
		rows[m.Id] = &statusRow{
			Id:       m.Id,
			Migrated: false,
		}
	}

	for _, r := range records {
		if rows[r.Id] == nil {
			utils.Error(fmt.Sprintf("Could not find migration file: %v", r.Id))
			continue
		}

		rows[r.Id].Migrated = true
		rows[r.Id].AppliedAt = r.AppliedAt
	}

	for _, m := range ms {
		if rows[m.Id] != nil && rows[m.Id].Migrated {
			table.Append([]string{
				m.Id,
				rows[m.Id].AppliedAt.String(),
			})
		} else {
			table.Append([]string{
				m.Id,
				"no",
			})
		}
	}

	table.Render()
	return
}

func RollbackSQL() (err error) {
	utils.Info("Begin rollback")

	err = InitDB()
	if err != nil {
		return
	}

	initMigration()

	migrate.SetTable("migrations")
	_, err = migrate.Exec(db, dialect, migrations, migrate.Down)
	if err != nil {
		return
	}

	err = MigrateSQLStatus()
	if err != nil {
		return
	}

	utils.Info("Rollback successfully completed")
	return
}
