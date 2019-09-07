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

func nodesGetAllHandler(w http.ResponseWriter, r *http.Request) {
	nodes, err := db.GetNodes()
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	res, err := json.Marshal(nodes)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)

	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func nodesGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	node, err := db.GetNodeByID(id)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	res, err := json.Marshal(node)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)

	}
	_, err = w.Write(res)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func nodesPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.NodeData
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

	err = db.InsertNode(jPostData.Name, jPostData.ModelID)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model, err := db.GetNode(jPostData.Name)
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
