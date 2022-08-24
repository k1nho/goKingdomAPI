package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/k1nho/goKingdomAPI/model"
)

// /api/kingdom/exports
func GetExports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	exports, err := model.GetExports()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get Exports"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exports)
}

// /api/kingdom/exports/{export_id}
func GetExport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["export_id"]
	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parameter given is not valid"))
		return
	}

	export, err := model.GetExport(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not fetch export"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(export)
}
