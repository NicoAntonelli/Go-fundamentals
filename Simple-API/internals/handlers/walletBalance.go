package handlers

import (
	"encoding/json"
	"net/http"

	"simple-api/api"
	ent "simple-api/internals/entities"
	"simple-api/internals/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// Types
type WalletBalanceGetRequest struct {
	Username string
}

type WalletBalanceGetResponse struct {
	Status  int
	Balance float32
}

// GET
func WalletBalanceGetOne(w http.ResponseWriter, r *http.Request) {
	var params = WalletBalanceGetRequest{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *ent.WalletDetails = (*database).GetWalletDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err) // This err is empty
		api.InternalErrorHandler(w)
		return
	}

	var response = WalletBalanceGetResponse{
		Balance: (*tokenDetails).Balance,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
