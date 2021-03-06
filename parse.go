package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJson(
	w http.ResponseWriter,
	r *http.Request,
	o interface{},
) bool {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, ErrorBodyRead, http.StatusInternalServerError)
		return false
	}

	if err := json.Unmarshal(b, &o); err != nil {
		http.Error(w, ErrorBadJson, http.StatusBadRequest)
		return false
	}

	return true
}
