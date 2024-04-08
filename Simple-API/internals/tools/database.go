package tools

import log "github.com/sirupsen/logrus"

// Database collections
type LoginDetails struct {
	AuthToken string
	Username  string
}

type WalletDetails struct {
	Username string
	Balance  float32
}

// Database interface
type DatabaseInterface interface {
	GetLoginDetails(username string) *LoginDetails
	GetWalletDetails(username string) *WalletDetails
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
