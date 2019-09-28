package db

import (
	"database/sql"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type RefTypePlanStatusData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetRefTypePlanStatusByID(id int) (res RefTypePlanStatusData, err error) {
	err = db.QueryRow(`select * from public.ref_type_status where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func DeleteRefTypePlanStatusByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.ref_type_status where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertRefTypePlanStatus(name string) (id int, err error) {
	err = db.QueryRow(`insert into public.ref_type_status (name) values ($1) returning id`,
		name).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllRefTypePlanStatus() (res []RefTypePlanStatusData, err error) {
	rows, err := db.Query(`select * from public.ref_type_status`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := RefTypePlanStatusData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}