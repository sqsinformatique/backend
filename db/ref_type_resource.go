package db

import (
	"database/sql"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type RefTypeResourceData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetRefTypeResourceByID(id int) (res RefTypeResourceData, err error) {
	err = db.QueryRow(`select * from public.ref_type_resource where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func DeleteRefTypeResourceByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.ref_type_resource where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertRefTypeResource(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.ref_type_resource (name) values ($1) returning id`,
		name)
	if rows.Scan(&id) == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	rows.Scan(&id)
	return
}

func GetAllRefTypeResource() (res []RefTypeResourceData, err error) {
	rows, err := db.Query(`select * from public.ref_type_resource`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := RefTypeResourceData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
