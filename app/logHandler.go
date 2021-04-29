package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/service"
)

type LogHandlers struct {
	service service.LogService 
}

func (lh LogHandlers) newLog(w http.ResponseWriter, r *http.Request) {
	var req dto.NewLogRequest
	vars := mux.Vars(r)
	id := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		log, appErr := lh.service.NewLog(id, req)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, log)
		}
	}
}
func (lh LogHandlers) getAllLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	logs, err := lh.service.GetAllLogs(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, logs)
	}
}
func (lh LogHandlers) getLogsByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	date := vars["date"]
	logs, err := lh.service.GetLogsByDate(id, date)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, logs)
	}
}
func (lh LogHandlers) getLogsByNic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["nic"]
	logs, err := lh.service.GetLogsByVisitorId(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, logs)
	}
}
