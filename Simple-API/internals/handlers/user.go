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
type UserGetRequest struct {
	Username string
}

type UserGetResponse struct {
	Users []ent.UserDetails
}

type UserPostRequest struct {
	Username string
	Mail     string
}

type UserPostResponse struct {
	Users []ent.UserDetails
}

// GET
func UserGet(w http.ResponseWriter, r *http.Request) {
	var params = UserGetRequest{}
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

	var users []ent.UserDetails = (*database).GetUsers()
	if users == nil {
		log.Error(err) // This err is empty
		api.InternalErrorHandler(w)
		return
	}

	var response = UserGetResponse{
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

// POST
func UserPost(w http.ResponseWriter, r *http.Request) {
	var newUser ent.UserDetails
	var decoder *json.Decoder = json.NewDecoder(r.Body)

	var err error = decoder.Decode(&newUser)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Before exit function
	defer r.Body.Close()

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var users []ent.UserDetails = (*database).PostUser(newUser)
	if users == nil {
		log.Error(err) // This err is empty
		api.InternalErrorHandler(w)
		return
	}

	var response = UserGetResponse{
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
