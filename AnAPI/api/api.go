package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coiunt nBalance Respionse
type CoinBalanceResponse struct {
	// Success Code ,usually 200 for success, 400 for bad request, 500 for server error
	Code int
	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	//Error Code
	Code int

	//Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error(err)
	}
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)
