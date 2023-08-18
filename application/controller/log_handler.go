package application

import (
	"encoding/json"
	"net/http"

	infra "github.com/hokita/milk/infra/repository"
)

type LogHandler struct {
	Repository *infra.LogRepository
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.Repository.Get()
	if err := json.NewEncoder(w).Encode(log); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
