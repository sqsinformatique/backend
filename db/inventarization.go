package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sqsinformatique/backend/utils"
)

type InventarizationData struct {
	ID                 int       `json:"id"`
	Object             int       `json:"object"`
	SupplyOrganization string    `json:"supply_organization"`
	Percent            int       `json:"percent"`
	PlanFault          time.Time `json:"plan_fault"`
	Cost               int       `json:"cost"`
	PlanStatus         int       `json:"plan_status"`
}

func GetInventarizationByID(id int) (res InventarizationData, err error) {
	err = db.QueryRow(`select * from public.inventarization where id=$1`, id).Scan(&res.ID, &res.Object, &res.SupplyOrganization,
		&res.Percent, &res.PlanFault, &res.Cost, &res.PlanStatus)
	return
}

func DeleteInventarizationByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.inventarization where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertInventarization(object int, supply_organization, percent int, planFault time.Time, cost, planStatus int) (id int, err error) {
	err = db.QueryRow(`insert into public.incidents (object, supply_organization, percent, plan_fault, cost, plan_status) values ($1, $2, $3, $4, $5, $6) returning id`,
		object, supply_organization, percent, planFault, cost, planStatus).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllInventarization() (res []InventarizationData, err error) {
	rows, err := db.Query(`select * from public.incidents`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := InventarizationData{}
		err := rows.Scan(&p.ID, &p.Object, &p.SupplyOrganization,
			&p.Percent, &p.PlanFault, &p.Cost, &p.PlanStatus)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func AddInventarizationPartitionsBySupplyOrganization(supplyOrganization int) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	_, err = tx.Exec(fmt.Sprintf(`CREATE TABLE inventarization_%s PARTITION OF inventarization FOR VALUES IN (%s)`,
		supplyOrganization,
		supplyOrganization,
	))
	if err != nil {
		return
	}

	return tx.Commit()
}
