package controller

import (
	"net/http"
	"strings"

	"github.com/partha-sen/ostd/webservice/service"
)

func HandleOpening(w http.ResponseWriter, r *http.Request) {

	strId := strings.TrimPrefix(r.URL.Path, "/openings/")

	switch r.Method {

	case http.MethodGet:
		processGet(w, strId, service.GetOpeningById)

	case http.MethodPut:
		processPut(w, r, strId, service.UpdateOpening)

	case http.MethodDelete:
		processDelete(w, strId, service.DeleteOpening)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func HandleOpenings(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		processGetAll(w, service.GetAllOpening)

	case http.MethodPost:
		processPost(w, r, service.SaveOpening)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
