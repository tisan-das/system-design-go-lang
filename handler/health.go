package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tisan-das/system-design-go-lang/dto"
)

func Health(response http.ResponseWriter, request *http.Request) {

	healthStruct := dto.HealthResponse{Status: "running"}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	var responseBytes bytes.Buffer
	json.NewEncoder(&responseBytes).Encode(healthStruct)
	response.Write(responseBytes.Bytes())
}
