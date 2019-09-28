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

func ref_planStatusGetAllHandler(w http.ResponseWriter, r *http.Request) {
	ref_status, err := db.GetAllRefTypePlanStatus()
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(ref_status)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func ref_planStatusPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.RefTypePlanStatusData
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

	id, err := db.InsertRefTypePlanStatus(jPostData.Name)
	utils.Infoln("Insert RefTypePlanStatus ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ref_status, err := db.GetRefTypePlanStatusByID(id)
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(ref_status)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func ref_planStatusDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.DeleteRefTypePlanStatusByID(id)
	if err != nil {
		utils.Errorf("Error DeleteRefTypePlanStatusByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	return
}

func ref_planStatusGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ref_status, err := db.GetRefTypePlanStatusByID(id)
	if err != nil {
		utils.Errorf("Error GetRefTypeResourceByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jGetData, err := json.Marshal(ref_status)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}
