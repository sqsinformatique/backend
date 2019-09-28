package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sqsinformatique/backend/utils"
)

type IncidentsData struct {
	ID                 int       `json:"id"`
	SupplyOrganization int       `json:"supply_organization"`
	Object             int       `json:"object"`
	Date               time.Time `json:"date"`
	Results            string    `json:"results"`
	ResponsibleWorker  string    `json:"responsible_worker"`
}

func GetIncidentsByID(id int) (res IncidentsData, err error) {
	err = db.QueryRow(`select * from public.incidents where id=$1`, id).Scan(&res.ID, &res.SupplyOrganization,
		&res.Object, &res.Date, &res.Results, &res.ResponsibleWorker)
	return
}

func DeleteIncidentsByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.incidents where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertIncidents(supply_organization, object int, date time.Time, results, responsibleWorker string) (id int, err error) {
	err = db.QueryRow(`insert into public.incidents (supply_organization, object, date, results, responsible_worker) values ($1, $2, $3, $4, $5) returning id`,
		supply_organization, object, date, results, responsibleWorker).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllIncidents() (res []IncidentsData, err error) {
	rows, err := db.Query(`select * from public.incidents`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := IncidentsData{}
		err := rows.Scan(&p.ID, &p.SupplyOrganization,
			&p.Object, &p.Date, &p.Results, &p.ResponsibleWorker)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func AddIncidentsPartitionsBySupplyOrganization(supplyOrganization int) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	_, err = tx.Exec(fmt.Sprintf(`CREATE TABLE incidents_%s PARTITION OF incidents FOR VALUES IN (%s)`,
		supplyOrganization,
		supplyOrganization,
	))
	if err != nil {
		return
	}

	return tx.Commit()
}

func UpdateIncidents(id, supply_organization, object int, date time.Time, results, responsibleWorker string) (err error) {
	_, err = rollbackQuery(`insert into public.incidents (supply_organization, object, date, results, responsible_worker) values ($1, $2, $3, $4, $5) where id=$6`,
		supply_organization, object, date, results, responsibleWorker, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("Err insert model")
	}
	return
}
