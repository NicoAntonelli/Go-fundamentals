package tools

import "time"

type mockDB struct{}

// Mocked info
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]WalletDetails{
	"alex": {
		Username: "alex",
		Balance:  100,
	},
	"jason": {
		Username: "jason",
		Balance:  200,
	},
	"marie": {
		Username: "marie",
		Balance:  300,
	},
}

// Interface implementation
func (d *mockDB) GetLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetWalletDetails(username string) *WalletDetails {
	// Simulate DB call
	time.Sleep(time.Second)

	var clientData = WalletDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
