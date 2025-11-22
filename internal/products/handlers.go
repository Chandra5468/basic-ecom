package products

import (
	"log/slog"
	"net/http"

	"github.com/Chandra5468/basic-ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. call the service -> List Product
	// 2. Return JSON in an http response
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("db list products error ", "error", err)
		// notifying user
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// json.Write(w, http.StatusInternalServerError, "error from db for ")
		return
	}

	json.Write(w, http.StatusOK, products)
}
