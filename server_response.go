package server

import (
	"encoding/json"
	"net/http"
	"time"
)

// ErrorResponse is a default error response
type ErrorResponse struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// WriteRespWithStatus will write a response with a struct or a default error response with status 500
func WriteRespWithStatus(resp http.ResponseWriter, respStatus int, body interface{}) {

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(respStatus)
	if body != nil {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusCreated)
		json.NewEncoder(resp).Encode(body)
	}
}

// WriteErrorResp will write a default error response with status, message and timestamp
func WriteErrorResp(status int, message string, writer http.ResponseWriter) {
	errorResp := ErrorResponse{
		Status:    status,
		Message:   message,
		Timestamp: time.Now(),
	}
	writer.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(errorResp)
	writer.WriteHeader(status)
	writer.Write(resp)

}
