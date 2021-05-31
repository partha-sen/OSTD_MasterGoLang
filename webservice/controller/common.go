package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/partha-sen/ostd/webservice/model"
	"github.com/partha-sen/ostd/webservice/service"
)

func writeJson(w http.ResponseWriter, obj interface{}) {

	data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func validatePathParam(w http.ResponseWriter, strId string) (int, bool) {
	if len(strId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return 0, false
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func processGet(w http.ResponseWriter, strId string, m func(id int) (model.Any, error)) {

	id, ok := validatePathParam(w, strId)
	if !ok {
		return
	}

	value, err := m(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeJson(w, value)

}

func processPut(w http.ResponseWriter, r *http.Request, strId string, m func([]byte, int) (int64, error)) {

	pathId, ok := validatePathParam(w, strId)
	if !ok {
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := m(bodyBytes, pathId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(id, 10))
	w.Write(b)
}

func processGetAll(w http.ResponseWriter, m func() ([]model.Any, error)) {
	values, err := m()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(values) > 0 {
		writeJson(w, values)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func processPost(w http.ResponseWriter, r *http.Request, m func([]byte) (int64, error)) {

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := m(bodyBytes)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(id, 10))
	w.Write(b)
}

func processDelete(w http.ResponseWriter, strId string, m func(id int) (int64, error)) {
	id, ok := validatePathParam(w, strId)
	if !ok {
		return
	}
	count, err := service.DeleteOpening(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(count, 10))
	w.Write(b)

}
