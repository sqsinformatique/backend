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

func supply_organizationsGetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func supply_organizationsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	supply_organizations, err := db.GetSupplyOrganizationByID(id)
	if err != nil {
		utils.Errorf("Error GetSupplyOrganizationByID: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jGetData, err := json.Marshal(supply_organizations)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func supply_organizationsPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.SupplyOrganizationData
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

	id, err := db.InsertSupplyOrganization(jPostData.Name, jPostData.TypeOfResource,
		jPostData.Description, jPostData.ContactTel, jPostData.ContactEmail, jPostData.Head)
	utils.Infoln("Insert SupplyOrganization ID", id)
	if err != nil {
		utils.Errorf("Can't INSERT. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	supply_organization, err := db.GetSupplyOrganizationByID(id)
	if err != nil {
		utils.Errorf("Can't GET. Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(supply_organization)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}
