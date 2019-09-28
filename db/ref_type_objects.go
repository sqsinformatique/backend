package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type RefTypeObjectsData struct {
	ID              int             `json:"id"`
	Resource        int             `json:"resource"`
	Name            string          `json:"name"`
	Characteristics json.RawMessage `json:"characteristics"`
}

func GetRefTypeObjectsByID(id int) (res RefTypeObjectsData, err error) {
	err = db.QueryRow(`select * from public.ref_type_objects where id=$1`, id).Scan(&res.ID, &res.Resource, &res.Name, &res.Characteristics)
	return
}

func DeleteRefTypeObjectsByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.ref_type_objects where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertRefTypeObjects(resource int, name string, characteristics json.RawMessage) (id int, err error) {
	err = db.QueryRow(`insert into public.ref_type_objects (name) values ($1) returning id`,
		resource, name, characteristics).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllRefTypeObjects() (res []RefTypeObjectsData, err error) {
	rows, err := db.Query(`select * from public.ref_type_objects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := RefTypeObjectsData{}
		err := rows.Scan(&p.ID, &p.Resource, &p.Name, &p.Characteristics)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
