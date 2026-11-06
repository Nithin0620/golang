package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nithin0620/golang/api"
	"github.com/Nithin0620/golang/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var database tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = database.GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Errorf("no coin details found for user %q", params.Username)
		api.RequestErrorHandler(w, fmt.Errorf("user %q not found", params.Username))
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(responseBytes); err != nil {
		log.Error(err)
	}
}
