package srv

import (
	"net/http"
)

func healthGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}