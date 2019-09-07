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

func contractorsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	elevators, err := db.GetContractors()
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

func contractorsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	contractor, err := db.GetContractorByID(id)
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

func contractorsPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.ContractorData
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

	id, err := db.InsertContractor(jPostData.Name)
	utils.Infoln("Insert elevator ID", id)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model, err := db.GetContractor(jPostData.Name)
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
