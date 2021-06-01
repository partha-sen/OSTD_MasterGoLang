package controller

import (
	"net/http"
	"strings"

	"github.com/partha-sen/ostd/webservice/model"
	"github.com/partha-sen/ostd/webservice/service"
)

func HandleInterview(w http.ResponseWriter, r *http.Request) {

	strId := strings.TrimPrefix(r.URL.Path, "/interviews/")

	switch r.Method {

	case http.MethodGet:
		processGet(w, strId, service.GetInterviewById)

	case http.MethodPut:
		id, err := parsePathParam(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var obj model.Interview
		err = requestBodyToObject(w, r, &obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		obj.Id = id
		callUpdate := func(obj model.Any) (int64, error) {
			return service.UpdateInterview(obj.(model.Interview))
		}
		processPut(w, obj, callUpdate)

	case http.MethodDelete:
		processDelete(w, strId, service.DeleteInterview)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func HandleInterviews(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		processGetAll(w, service.GetAllInterview)

	case http.MethodPost:
		var obj model.Interview
		err := requestBodyToObject(w, r, &obj)

		if err == nil {
			callSave := func(obj model.Any) (int64, error) {
				return service.SaveInterview(obj.(model.Interview))
			}
			processPost(w, obj, callSave)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
