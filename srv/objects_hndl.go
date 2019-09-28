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

func objectsGetAllByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceType, _ := strconv.Atoi(vars["type"])
	supply_organizations, err := db.GetAllSupplyOrganizationByType(resourceType)
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var res []db.ObjectsData
	for _, organization := range supply_organizations {
		objects, err := db.GetAllObjectsByType(organization.ID)
		if err != nil {
			utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res = append(res, objects...)
	}

	jGetData, err := json.Marshal(res)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func objectsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	objects, err := db.GetAllObjects()
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

func objectsPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.ObjectsData
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

	id, err := db.InsertObjects(supply_organizations_id, jPostData.Coordinates,
		jPostData.ObjectType, jPostData.Characteristics, jPostData.Address, jPostData.Description,
		jPostData.Status, jPostData.MaintenanceDate, jPostData.LastRepairsDate)
	utils.Infoln("Insert InsertObjects ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	objects, err := db.GetObjectsByID(id)
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

func objectsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	objects, err := db.GetObjectsByID(id)
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

func objectsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.DeleteObjectsByID(id)
	if err != nil {
		utils.Errorf("Error GetSupplyOrganizationByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	return
}

func objectsPutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var jPostData db.ObjectsData
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

	err = db.UpdateObjects(id, jPostData.SupplyOrganization, jPostData.Coordinates,
		jPostData.ObjectType, jPostData.Characteristics, jPostData.Address, jPostData.Description,
		jPostData.Status, jPostData.MaintenanceDate, jPostData.LastRepairsDate)
	utils.Infoln("Insert InsertObjects ID", id)
	if err != nil {
		utils.Errorf("Can't UPDATE. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	objects, err := db.GetObjectsByID(id)
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
