package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/dto"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/webserver/services"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/webserver/utils"
)

func GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		http.Error(w, "Missing CEP parameter", http.StatusBadRequest)
		return
	}
	cep, ok := utils.ValidateCEP(cep)
	if !ok {
		http.Error(w, "Invalid CEP format. Must be 8 digits.", http.StatusBadRequest)
		return
	}

	var result interface{}
	brAPICh := make(chan *dto.APIResponse, 1)
	viaCEPCh := make(chan *dto.APIResponse, 1)

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go services.FetchFromBrasilAPI(ctx, cep, brAPICh)
	go services.FetchFromViaCEP(ctx, cep, viaCEPCh)

	select {
	case brAPIResponse := <-brAPICh:
		utils.LogAPIResponse(brAPIResponse, time.Since(start))
		result = brAPIResponse

	case viaCEPResponse := <-viaCEPCh:
		utils.LogAPIResponse(viaCEPResponse, time.Since(start))
		result = viaCEPResponse

	case <-ctx.Done():
		http.Error(w, "Timeout", http.StatusGatewayTimeout)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
