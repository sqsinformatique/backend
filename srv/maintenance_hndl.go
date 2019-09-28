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

func maintenanceGetAllHandler(w http.ResponseWriter, r *http.Request) {
	maintenance, err := db.GetAllMaintenance()
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(maintenance)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func maintenancePostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.MaintenanceData
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

	id, err := db.InsertMaintenance(supply_organizations_id, jPostData.Object,
		jPostData.MaintenanceType, jPostData.MaintenanceStart, jPostData.MaintenanceEnd, jPostData.Checklist,
		jPostData.ResponsibleWorker)
	utils.Infoln("Insert InsertMaintenance ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	maintenance, err := db.GetMaintenanceByID(id)
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(maintenance)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func maintenanceGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	objects, err := db.GetMaintenanceByID(id)
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

func maintenanceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.DeleteMaintenanceByID(id)
	if err != nil {
		utils.Errorf("Error GetSupplyOrganizationByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	return
}
