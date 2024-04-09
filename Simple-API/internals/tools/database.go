package tools

import (
	ent "simple-api/internals/entities"

	log "github.com/sirupsen/logrus"
)

// Database interface
type DatabaseInterface interface {
	GetLoginDetails(username string) *ent.LoginDetails
	GetWalletDetails(username string) *ent.WalletDetails
	GetUsers() []ent.UserDetails
	PostUser(user ent.UserDetails) []ent.UserDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
