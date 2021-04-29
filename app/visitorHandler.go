package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/service"
)

type VisitorHandlers struct {
	service service.VisitorService
}

func (vh VisitorHandlers) getAllVisitors(w http.ResponseWriter, r *http.Request) {
	visitors, err := vh.service.GetAllVisitors()
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, visitors)
	}
}
func (vh VisitorHandlers) getVisitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nic := vars["nic"]
	visitor, err := vh.service.GetVisitor(nic)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, visitor)
	}
}
func (vh VisitorHandlers) newVisitor(w http.ResponseWriter, r *http.Request) {
	var req dto.NewVisitorRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		visitor, appErr := vh.service.NewVisitor(req)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, visitor)
		}
	}
}
func (vh VisitorHandlers) editVisitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nic := vars["nic"]
	var req dto.NewVisitorRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		newVis, appErr := vh.service.EditVisitor(nic, req)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, newVis)
		}
	}
}
func (vh VisitorHandlers) removeVisitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nic := vars["nic"]
	err := vh.service.RemoveVisitor(nic)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage().Code)
	} else {
		WriteResponse(w, http.StatusOK, "Deleted visitor successfully")
	}
}
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
