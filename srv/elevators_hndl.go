package srv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sqsinformatique/backend/utils"

	"github.com/sqsinformatique/backend/db"

	"github.com/gorilla/mux"
)

type WorkGetData struct {
	db.ContractorData `json:"contractor"`
	db.WorktypeData   `json:"worktype"`
	Description       string `json:"description"`
}

func elevatorWorksGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	works, err := db.GetWorksByEquID(id)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	var res []WorkGetData
	for _, w := range works {
		p := WorkGetData{}

		contractor, err := db.GetContractorByID(w.Contrid)
		if err != nil {
			utils.Error(err)
			continue
		}

		worktype, err := db.GetWorktypeByID(w.Worktypeid)
		if err != nil {
			utils.Error(err)
			continue
		}

		p.WorktypeData = worktype
		p.ContractorData = contractor
		p.Description = w.Description
		res = append(res, p)
	}

	jGetData, err := json.Marshal(res)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}

}

type ElevatorDataWithContractor struct {
	db.ElevatorData
	db.ContractorData `json:"contractor"`
}

func elevatorContractorGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	elevatorWithContractor, err := db.GetContractorForElevator(id)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	contractor, err := db.GetContractorByID(elevatorWithContractor.ContractorID)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	jGetData, err := json.Marshal(contractor)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}

}

func elevatorsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	elevators, err := db.GetElevators()
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	jGetData, err := json.Marshal(elevators)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func elevatorsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	elevator, err := db.GetElevatorByID(id)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	jGetData, err := json.Marshal(elevator)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func elevatorsPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.ElevatorData
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(contents, &jPostData)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := db.InsertElevator(jPostData.Name, jPostData.ModelID, jPostData.Street,
		jPostData.Build,
		jPostData.Gate,
		jPostData.TotalGates,
	)
	utils.Infoln("Insert elevator ID", id)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model, err := db.GetElevator(jPostData.Name)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(model)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}
