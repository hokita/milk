package application

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	infra "github.com/hokita/milk/infra/repository"
)

type HelloHandler struct {
	Repository *infra.HelloRepository
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello := h.Repository.Get()
	if err := json.NewEncoder(w).Encode(hello); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
