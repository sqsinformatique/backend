package db

import (
	"database/sql"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type RefTypeStatusData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetRefTypeStatusByID(id int) (res RefTypeStatusData, err error) {
	err = db.QueryRow(`select * from public.ref_type_status where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func DeleteRefTypeStatusByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.ref_type_status where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertRefTypeStatus(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.ref_type_status (name) values ($1) returning id`,
		name)
	if rows.Scan(&id) == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	rows.Scan(&id)
	return
}

func GetAllRefTypeStatus() (res []RefTypeStatusData, err error) {
	rows, err := db.Query(`select * from public.ref_type_status`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := RefTypeStatusData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
