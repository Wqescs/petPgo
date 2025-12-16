package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Wqescs/petPgo/calc/internal/service"
	"github.com/Wqescs/petPgo/calc/pkg/decimal"
)

type HTTPHandler struct {
	calculator *service.Calculator
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		calculator: service.New(),
	}
}

func (h *HTTPHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/calculate", h.Calculate).Methods("POST")
	router.HandleFunc("/api/v1/health", h.HealthCheck).Methods("GET")
}

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result    string `json:"result"`
	Value     string `json:"value"`
	Error     string `json:"error,omitempty"`
	Precision int    `json:"precision,omitempty"`
}

func (h *HTTPHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CalculateResponse{
			Error: "Invalid request format",
		})
		return
	}
	
	result, err := h.calculator.Calculate(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CalculateResponse{
			Error: err.Error(),
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{
		Result:    result.Expression,
		Value:     decimal.Format(result.Value, result.Precision),
		Precision: result.Precision,
	})
}

func (h *HTTPHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "calculator",
	})
}