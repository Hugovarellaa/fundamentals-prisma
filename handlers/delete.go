package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Hugovarellaa/curd-go/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParams(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}