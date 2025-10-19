package utils

import (
	"encoding/json"
	"log"
	"time"

	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/dto"
)

func LogAPIResponse(apiResp *dto.APIResponse, duration time.Duration) {
	if apiResp == nil {
		log.Println("API response is nil")
		return
	}

	jsonData, err := json.Marshal(apiResp.Data)
	if err != nil {
		log.Printf("Error marshalling API response: %v\n", err)
		return
	}

	log.Printf("Fastest response: %s (%s) - Data: %s",
		apiResp.Source,
		duration,
		string(jsonData),
	)
}
