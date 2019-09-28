package db

import (
	"database/sql"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type RefTypeMaintenanceData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetRefTypeMaintenanceByID(id int) (res RefTypeMaintenanceData, err error) {
	err = db.QueryRow(`select * from public.ref_type_maintenance where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func DeleteRefTypeMaintenanceByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.ref_type_maintenance where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertRefTypeMaintenance(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.ref_type_maintenance (name) values ($1) returning id`,
		name)
	if rows.Scan(&id) == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	rows.Scan(&id)
	return
}

func GetAllRefTypeMaintenance() (res []RefTypeMaintenanceData, err error) {
	rows, err := db.Query(`select * from public.ref_type_maintenance`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := RefTypeMaintenanceData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
