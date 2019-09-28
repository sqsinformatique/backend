package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sqsinformatique/backend/utils"
)

type MaintenanceData struct {
	ID                 int       `json:"id"`
	SupplyOrganization string    `json:"supply_organization"`
	Object             int       `json:"object"`
	MaintenanceType    int       `json:"maintenance_type"`
	MaintenanceStart   time.Time `json:"maintenance_start"`
	MaintenanceEnd     time.Time `json:"maintenance_end"`
	Checklist          string    `json:"checklist`
	ResponsibleWorker  string    `json:"responsible_worker"`
	Cost               int       `json:"cost"`
	Progress           int       `json:"progress"`
}

func GetMaintenanceByID(id int) (res MaintenanceData, err error) {
	err = db.QueryRow(`select * from public.maintenance where id=$1`, id).Scan(&res.ID, &res.SupplyOrganization,
		&res.Object, &res.MaintenanceType, &res.MaintenanceStart, &res.MaintenanceEnd, &res.Checklist, &res.ResponsibleWorker, &res.Cost)
	return
}

func DeleteMaintenanceByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.maintenance where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertMaintenance(supply_organization, object int, maintenanceType int, maintenanceStart time.Time, maintenanceEnd time.Time, checklist, responsibleWorker string, cost, progress int) (id int, err error) {
	err = db.QueryRow(`insert into public.maintenance (supply_organization, object, maintenance_type, maintenance_start, maintenance_end, checklist, responsible_worker) values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		supply_organization, object, maintenanceType, maintenanceStart, maintenanceEnd, checklist, responsibleWorker, cost, progress).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert Incidents")
	}
	return
}

func GetAllMaintenance() (res []MaintenanceData, err error) {
	rows, err := db.Query(`select * from public.maintenance`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := MaintenanceData{}
		err := rows.Scan(&p.ID, &p.SupplyOrganization,
			&p.Object, &p.MaintenanceType, &p.MaintenanceStart, &p.MaintenanceEnd, &p.Checklist, &p.ResponsibleWorker, &p.Cost, &p.Progress)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func AddMaintenancePartitionsBySupplyOrganization(supplyOrganization int) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	_, err = tx.Exec(fmt.Sprintf(`CREATE TABLE maintenance_%s PARTITION OF maintenance FOR VALUES IN (%s)`,
		supplyOrganization,
		supplyOrganization,
	))
	if err != nil {
		return
	}

	return tx.Commit()
}
