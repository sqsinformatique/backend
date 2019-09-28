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

func ref_statusGetAllHandler(w http.ResponseWriter, r *http.Request) {
	ref_status, err := db.GetAllRefTypeStatus()
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

func ref_statusPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.RefTypeStatusData
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

	id, err := db.InsertRefTypeStatus(jPostData.Name)
	utils.Infoln("Insert RefTypeStatus ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ref_status, err := db.GetRefTypeStatusByID(id)
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

func ref_statusDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.DeleteRefTypeStatusByID(id)
	if err != nil {
		utils.Errorf("Error DeleteRefTypeStatusByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	return
}

func ref_statusGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ref_status, err := db.GetRefTypeStatusByID(id)
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
