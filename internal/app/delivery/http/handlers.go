package http

import (
	"encoding/json"
	"net/http"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/service"
)

type WarehouseHandler struct {
	service *service.WarehouseService
}

func NewWarehouseHandler(service *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: service}
}

type createWarehouseRequest struct {
	Address string `json:"address" validate:"required,min=3,max=255"`
}

func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	warehouse, err := h.service.Create(r.Context(), req.Address)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create warehouse")
		return
	}

	respondWithJSON(w, http.StatusCreated, warehouse)
}

func (h *WarehouseHandler) List(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.service.List(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get warehouses")
		return
	}

	respondWithJSON(w, http.StatusOK, warehouses)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
