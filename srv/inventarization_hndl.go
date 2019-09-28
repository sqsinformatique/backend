package srv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sqsinformatique/backend/utils"

	"github.com/sqsinformatique/backend/db"
)

func inventarizationGetAllHandler(w http.ResponseWriter, r *http.Request) {
	inventarization, err := db.GetAllInventarization()
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(inventarization)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func inventarizationGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	objects, err := db.GetInventarizationByID(id)
	if err != nil {
		utils.Errorf("Error GetSupplyOrganizationByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jGetData, err := json.Marshal(objects)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func inventarizationPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.InventarizationData
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

	vars := mux.Vars(r)
	supply_organizations_id, _ := strconv.Atoi(vars["supply_organizations_id"])

	id, err := db.InsertInventarization(jPostData.Object, supply_organizations_id,
		jPostData.Percent, jPostData.PlanFault, jPostData.Cost, jPostData.PlanStatus)
	utils.Infoln("Insert InsertInventarization ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	objects, err := db.GetInventarizationByID(id)
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(objects)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func inventarizationDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.DeleteInventarizationByID(id)
	if err != nil {
		utils.Errorf("Error DeleteInventarizationByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	return
}
