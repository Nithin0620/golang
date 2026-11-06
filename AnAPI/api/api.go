package api

import (
	"encoding/json"
	"net/http"
)
//coin balance params
type CointBalanceParams struct {
	UserName string 
}


//coiunt nBalance Respionse 
type CoinBalanceResponse struct {
	// Success Code ,usually 200 for success, 400 for bad request, 500 for server error
	code int
	// Account Balance
	Balance int64
}


// Error Response 
type Error struct {
	//Error Code
	code int

	//Error Message
	Message string
}


func writeError(w http.ResponseWriter, message string, code int) {
	resp:= Error{
		code :  code,
		Message: message,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}


var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(),http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter ) {
		writeError(w, "An Unexpected Error Occured." , http.statusInternalServerError)
	}
)