package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetJadkul(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	status := http.StatusOK
	response := GenericResponse{}
	defer func() {
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(response)
		w.WriteHeader(status)
		w.Write(jsonResponse)
	}()

	query := r.FormValue("query")
	if query == "" {
		status = http.StatusBadRequest
		response.Header.ErrorMessage = "Query param must not be empty"
	}

	jadkulList, err := h.Service.GetJadkul(query)
	if err != nil {
		status = http.StatusInternalServerError
		response.Header.ErrorMessage = "Something went wrong"
	}

	response.Header.Status = true
	response.Data = jadkulList

}
