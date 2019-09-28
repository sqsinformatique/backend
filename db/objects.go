package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sqsinformatique/backend/utils"
)

type ObjectsData struct {
	ID                 int             `json:"id"`
	SupplyOrganization int             `json:"supply_organization"`
	Coordinates        json.RawMessage `json:"coordinates"`
	ObjectType         int             `json:"object_type"`
	Address            string          `json:"address"`
	Characteristics    json.RawMessage `json:"characteristics"`
	Description        string          `json:"description"`
	Status             int             `json:"status"`
	MaintenanceDate    time.Time       `json:"maintenance_date"`
	LastRepairsDate    time.Time       `json:"last_repairs_date"`
}

func GetObjectsByID(id int) (res ObjectsData, err error) {
	err = db.QueryRow(`select * from public.objects where id=$1`, id).Scan(&res.ID, &res.SupplyOrganization,
		&res.Coordinates, &res.ObjectType, &res.Address, &res.Characteristics, &res.Description, &res.Status, &res.MaintenanceDate, &res.LastRepairsDate)
	return
}

func DeleteObjectsByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.objects where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertObjects(supplyOrganization int, coordinates json.RawMessage, objectType int, characteristics json.RawMessage,
	address, description string, status int, maintenanceDate, lastRepairsDate time.Time) (id int, err error) {
	err = db.QueryRow(`insert into public.objects (supply_organization, coordinates, object_type, characteristics, address, description, status, maintenance_date, last_repairs_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`,
		supplyOrganization, coordinates, objectType, characteristics, address, description, status, maintenanceDate, lastRepairsDate).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllObjects() (res []ObjectsData, err error) {
	rows, err := db.Query(`select * from public.objects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := ObjectsData{}
		err := rows.Scan(&p.ID, &p.SupplyOrganization,
			&p.Coordinates, &p.ObjectType, &p.Address, &p.Characteristics, &p.Description, &p.Status, &p.MaintenanceDate, &p.LastRepairsDate)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func GetAllObjectsByType(supplyOrganization int) (res []ObjectsData, err error) {
	rows, err := db.Query(`select * from public.objects where supply_organization=$1`, supplyOrganization)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := ObjectsData{}
		err := rows.Scan(&p.ID, &p.SupplyOrganization,
			&p.Coordinates, &p.ObjectType, &p.Address, &p.Characteristics, &p.Description, &p.Status, &p.MaintenanceDate, &p.LastRepairsDate)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
