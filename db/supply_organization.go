package db

import (
	"database/sql"
	"fmt"

	"github.com/sqsinformatique/backend/utils"
)

type SupplyOrganizationData struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	TypeOfResource int    `json:"type_of_resource"`
	Description    string `json:"description"`
	ContactTel     string `json:"contact_tel"`
	ContactEmail   string `json:"contact_email"`
	Head           string `json:"head"`
}

func GetSupplyOrganizationByID(id int) (res SupplyOrganizationData, err error) {
	err = db.QueryRow(`select * from public.supply_organization where id=$1`, id).Scan(&res.ID, &res.Name,
		&res.TypeOfResource, &res.Description, &res.ContactTel, &res.ContactEmail, &res.Head)
	return
}

func DeleteRefTypeOrganizationByID(id int) (err error) {
	_, err = rollbackQuery(`delete from public.supply_organization where id=$1`, id)
	if err != nil {
		return err
	}
	return
}

func InsertSupplyOrganization(name string, typeOfResource int, description, contactTel, contactEmail, head string) (id int, err error) {
	err = db.QueryRow(`insert into public.supply_organization (name, type_of_resource, description, contact_tel, contact_email, head) values ($1, $2, $3, $4, $5, $6) returning id`,
		name, typeOfResource, description, contactTel, contactEmail, head).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func GetAllSupplyOrganization() (res []SupplyOrganizationData, err error) {
	rows, err := db.Query(`select * from public.supply_organization`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := SupplyOrganizationData{}
		err := rows.Scan(&p.ID, &p.Name, &p.TypeOfResource, &p.Description, &p.ContactTel, &p.ContactEmail, &p.Head)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func GetAllSupplyOrganizationByType(typeOfResource int) (res []SupplyOrganizationData, err error) {
	rows, err := db.Query(`select * from public.supply_organization where type_of_resource=$1`, typeOfResource)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := SupplyOrganizationData{}
		err := rows.Scan(&p.ID, &p.Name, &p.TypeOfResource, &p.Description, &p.ContactTel, &p.ContactEmail, &p.Head)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}
