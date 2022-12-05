package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Hugovarellaa/curd-go/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParams(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}
