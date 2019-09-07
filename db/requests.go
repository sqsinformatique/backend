package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sqsinformatique/backend/utils"
)

type WorkData struct {
	ID          int    `json:"id"`
	Equid       int    `json:"equid"`
	Contrid     int    `json:"contrid"`
	Worktypeid  int    `json:"worktypeid"`
	Description string `json:"description"`
}

func GetWorkByDesk(description string) (res WorkData, err error) {
	err = db.QueryRow(`select * from public.works where description=$1`, description).Scan(&res.ID, &res.Equid,
		&res.Contrid, &res.Worktypeid, &res.Description)
	return
}

func InsertWork(equid, contrid, worktypeid int, description string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.works (equid, contrid, worktypeid, description) values ($1, $2, $3, $4) returning id`,
		equid, contrid, worktypeid, description)
	if err != nil {
		return -1, err
	}
	rows.Scan(&id)
	return
}

func GetWorksByEquID(id int) (res []WorkData, err error) {
	rows, err := db.Query(`select * from public.works when equid=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := WorkData{}
		err := rows.Scan(&p.ID, &p.Equid, &p.Contrid, &p.Worktypeid, &p.Description)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

type WorktypeData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetWorktype(name string) (res WorktypeData, err error) {
	err = db.QueryRow(`select * from public.worktype where name=$1`, name).Scan(&res.ID, &res.Name)
	return
}

func GetWorktypeByID(id int) (res WorktypeData, err error) {
	err = db.QueryRow(`select * from public.worktype where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func GetWorktypes() (res []WorktypeData, err error) {
	rows, err := db.Query(`select * from public.worktype`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := WorktypeData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func InsertWorktype(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.worktype (name) values ($1) returning id`, name)
	if err != nil {
		return -1, err
	}
	rows.Scan(&id)
	return
}

type CurrentContractor struct {
	ID           int `json:"queid"`
	ContractorID int `json:"contid"`
}

func GetContractorForElevator(id int) (res CurrentContractor, err error) {
	err = db.QueryRow(`select * from public.current_contactor where queid=$1`, id).Scan(&res.ID, &res.ContractorID)
	return
}

type ContractorData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetContractor(name string) (res ContractorData, err error) {
	err = db.QueryRow(`select * from public.contractors where name=$1`, name).Scan(&res.ID, &res.Name)
	return
}

func GetContractorByID(id int) (res ContractorData, err error) {
	err = db.QueryRow(`select * from public.contractors where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func GetContractors() (res []ContractorData, err error) {
	rows, err := db.Query(`select * from public.contractors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := ContractorData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func InsertContractor(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.contractors (name) values ($1) returning id`, name)
	if err != nil {
		return -1, err
	}
	rows.Scan(&id)
	return
}

type PredictionData struct {
	Rul       float64   `json:"rul"`
	Timecycle int       `json:"timecycle"`
	Timestamp time.Time `json:"timestamp"`
}

func GetPredictionsByID(id, fileid int) (res []PredictionData, err error) {
	rows, err := db.Query(`select rul, timecycle, date from predictions where equid=$1 and fileid=$2`, id, fileid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		p := PredictionData{}
		err := rows.Scan(&p.Rul, &p.Timecycle, &p.Timestamp)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

type ElevatorData struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	ModelID    int    `json:"model_id"`
	Street     string `json:"street"`
	Build      string `json:"build"`
	Gate       int    `json:"gate"`
	TotalGates int    `json:"totalgates"`
}

func GetElevator(name string) (res ElevatorData, err error) {
	err = db.QueryRow(`select * from public.elevator where name=$1`, name).Scan(&res.ID, &res.Name, &res.ModelID,
		&res.Street,
		&res.Build,
		&res.Gate,
		&res.TotalGates,
	)
	return
}

func GetElevatorByID(id int) (res ElevatorData, err error) {
	err = db.QueryRow(`select * from public.elevator where id=$1`, id).Scan(&res.ID, &res.Name,
		&res.Street,
		&res.ModelID,
		&res.Build,
		&res.Gate,
		&res.TotalGates,
	)
	return
}

func GetElevators() (res []ElevatorData, err error) {
	rows, err := db.Query(`select * from public.elevator`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := ElevatorData{}
		err := rows.Scan(&p.ID, &p.ModelID, &p.Name,
			&p.Street,
			&p.Build,
			&p.Gate,
			&p.TotalGates)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func InsertElevator(name string, model_id int, street, build string, gate, totalgates int) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.elevator (name, model_id, street, build, gate, totalgates) values ($1, $2, $3, $4, $5, $6) returning id`,
		name, model_id, street, build, gate, totalgates)
	if err != nil {
		return -1, err
	}
	rows.Scan(&id)
	return
}

type NodeData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ModelID int    `json:"model_id"`
}

func GetNode(name string) (res NodeData, err error) {
	err = db.QueryRow(`select * from public.node where name=$1`, name).Scan(&res.ID, &res.ModelID, &res.Name)
	return
}

func GetNodeByID(id int) (res NodeData, err error) {
	err = db.QueryRow(`select * from public.node where id=$1`, id).Scan(&res.ID, &res.ModelID, &res.Name)
	return
}

func GetNodes() (res []NodeData, err error) {
	rows, err := db.Query(`select * from public.node`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := NodeData{}
		err := rows.Scan(&p.ID, &p.ModelID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func InsertNode(name string, model_id int) (err error) {
	_, err = rollbackQuery(`insert into public.node (name, model_id) values ($1, $2)`, name, model_id)
	if err != nil {
		return err
	}
	return
}

func GetNodesOfModel(id int) (res []NodeData, err error) {
	rows, err := db.Query(`select * from public.node where model_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := NodeData{}
		err := rows.Scan(&p.ID, &p.ModelID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

type ModelData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetModel(name string) (res ModelData, err error) {
	err = db.QueryRow(`select * from public.models where name=$1`, name).Scan(&res.ID, &res.Name)
	return
}

func GetModelByID(id int) (res ModelData, err error) {
	err = db.QueryRow(`select * from public.models where id=$1`, id).Scan(&res.ID, &res.Name)
	return
}

func GetModels() (res []ModelData, err error) {
	rows, err := db.Query(`select * from public.models`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := ModelData{}
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			utils.Error(err)
			continue
		}
		res = append(res, p)
	}
	return
}

func InsertModel(name string) (id int, err error) {
	rows, err := rollbackQuery(`insert into public.models (name) values ($1) returning id`, name)
	if err != nil {
		return -1, err
	}
	if rows.Scan(&id) == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func UpdateModel(id int, name string) (res_id int, err error) {
	rows, err := rollbackQuery(`update public.models set name=$1 where id=$2 returning id`, id, name)
	if err != nil {
		return -1, err
	}
	if rows.Scan(&res_id) == sql.ErrNoRows {
		return -1, fmt.Errorf("Err insert model")
	}
	return
}

func rollbackQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err = tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	rows.Close()

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}
