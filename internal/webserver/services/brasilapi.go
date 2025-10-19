package services

import (
	"context"
	"encoding/json"

	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/dto"
)

func FetchFromBrasilAPI(ctx context.Context, cep string, ch chan<- *dto.APIResponse) {
	if envConfig == nil || envConfig.BrasilApiUrl == "" {
		panic("Missing BrasilApiUrl in config")
	}

	url := envConfig.BrasilApiUrl + cep
	body, err := Fetch(ctx, url)
	if err != nil {
		ch <- &dto.APIResponse{Source: "BrasilAPI", Err: err}
		return
	}

	var data dto.BrasilAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- &dto.APIResponse{Source: "BrasilAPI", Err: err}
		return
	}

	mapped := make(map[string]interface{})
	jsonBytes, _ := json.Marshal(data)
	json.Unmarshal(jsonBytes, &mapped)

	ch <- &dto.APIResponse{Source: "BrasilAPI", Data: mapped}
}
