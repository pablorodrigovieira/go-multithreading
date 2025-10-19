package services

import (
	"context"
	"encoding/json"

	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/dto"
)

func FetchFromViaCEP(ctx context.Context, cep string, ch chan<- *dto.APIResponse) {
	if envConfig == nil || envConfig.ViaCepApiUrl == "" {
		panic("Missing ViaCepApiUrl in config")
	}

	url := envConfig.ViaCepApiUrl + cep + "/json/"
	body, err := Fetch(ctx, url)
	if err != nil {
		ch <- &dto.APIResponse{Source: "ViaCEP", Err: err}
		return
	}

	var data dto.ViaCepAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- &dto.APIResponse{Source: "ViaCEP", Err: err}
		return
	}

	mapped := make(map[string]interface{})
	jsonBytes, _ := json.Marshal(data)
	json.Unmarshal(jsonBytes, &mapped)

	ch <- &dto.APIResponse{Source: "ViaCEP", Data: mapped}
}
