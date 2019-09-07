package srv

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sqsinformatique/backend/utils"

	"github.com/sqsinformatique/backend/db"

	"github.com/gorilla/mux"
)

func predictionGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fileid, _ := strconv.Atoi(vars["fileid"])
	predictions, err := db.GetPredictionsByID(id, fileid)
	if err != nil {
		utils.Errorf("Error GetModels: %s", err)
	}

	jGetData, err := json.Marshal(predictions)
	if err != nil {
		utils.Errorf("Can't marshaled request %s", err)
	}
	_, err = w.Write(jGetData)
	if err != nil {
		utils.Errorf("Can't send error request %s", err)
	}
}
