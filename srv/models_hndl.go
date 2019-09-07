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

type ModelDataWithNode struct {
	db.ModelData
	Nodes []db.NodeData `json:"nodes"`
}

func modelsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	models, err := db.GetModels()
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	var res []ModelDataWithNode
	for _, m := range models {
		p := ModelDataWithNode{}

		modelNodes, err := db.GetNodesOfModel(m.ID)
		if err != nil {
			utils.Error(err)
			continue
		}
		p.ID = m.ID
		p.Name = m.Name
		p.Nodes = modelNodes
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

func modelsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	model, err := db.GetModelByID(id)

	modelNodes, err := db.GetNodesOfModel(model.ID)
	if err != nil {
		utils.Error(err)
	}

	var res ModelDataWithNode

	res.ID = model.ID
	res.Name = model.Name
	res.Nodes = modelNodes

	jGetData, err := json.Marshal(res)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}

func modelsPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.ModelData
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

	id, err := db.InsertModel(jPostData.Name)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.Infoln("Insert model ID", id)
	model, err := db.GetModel(jPostData.Name)
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

func modelsPutHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.ModelData
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

	id, err := db.UpdateModel(jPostData.ID, jPostData.Name)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.Infoln("Insert model ID", id)
	model, err := db.GetModel(jPostData.Name)
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
