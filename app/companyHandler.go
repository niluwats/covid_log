package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/service"
)

type CompanyHandlers struct {
	service service.CompanyService
}

func (ch CompanyHandlers) getAllCompanies(w http.ResponseWriter, r *http.Request) {
	coms, err := ch.service.GetAllComanpies()
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, coms)
	}
}
func (ch CompanyHandlers) getCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	comp, err := ch.service.GetCompany(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, comp)
	}
}
func (ch CompanyHandlers) newCompany(w http.ResponseWriter, r *http.Request) {
	var req dto.NewCompanyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		comp, appErr := ch.service.NewCompany(req)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, comp)
		}
	}
}
func (ch CompanyHandlers) editCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var req dto.NewCompanyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		newComp, appErr := ch.service.EditCompany(id, req)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, newComp)
		}
	}
}
func (ch CompanyHandlers) removeCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := ch.service.RemoveCompany(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage().Code)
	} else {
		WriteResponse(w, http.StatusOK, "Deleted company successfully")
	}
}
