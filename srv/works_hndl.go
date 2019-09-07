package srv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sqsinformatique/backend/utils"

	"github.com/sqsinformatique/backend/db"
)

func worksPostHandler(w http.ResponseWriter, r *http.Request) {
	var jPostData db.WorkData
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

	id, err := db.InsertWork(jPostData.Equid, jPostData.Contrid, jPostData.Worktypeid, jPostData.Description)
	if err != nil {
		utils.Errorf("Something wrong with the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.Infoln("Insert model ID", id)
	model, err := db.GetWorkByDesk(jPostData.Description)
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
