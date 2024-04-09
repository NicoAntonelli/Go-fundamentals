package tools

import (
	"fmt"
	"math/rand"
	ent "simple-api/internals/entities"
	"time"
)

type mockDB struct{}

// Mocked info
var mockLoginDetails = map[string]ent.LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
		Mail:      "alex@mail.com",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
		Mail:      "jason@mail.com",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
		Mail:      "marie@mail.com",
	},
}

var mockWalletDetails = map[string]ent.WalletDetails{
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
func (db *mockDB) GetLoginDetails(username string) *ent.LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second)

	var clientData = ent.LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetWalletDetails(username string) *ent.WalletDetails {
	// Simulate DB call
	time.Sleep(time.Second)

	var clientData = ent.WalletDetails{}
	clientData, ok := mockWalletDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUsers() []ent.UserDetails {
	// Simulate DB call
	time.Sleep(time.Second)

	if mockLoginDetails == nil {
		return nil
	}

	var clientData = []ent.UserDetails{}

	for _, value := range mockLoginDetails {
		clientData = append(clientData, ent.UserDetails{
			Username: value.Username,
			Mail:     value.Mail,
		})
	}

	return clientData
}

func (db *mockDB) PostUser(user ent.UserDetails) []ent.UserDetails {
	// Get all users
	var clientData = db.GetUsers()
	if clientData == nil {
		return nil
	}

	// Create new token
	var newToken string = ""
	var tokenLong int = rand.Intn(6) + 6 // Random length of token
	for i := 0; i < tokenLong; i++ {
		newToken += string(rune(rand.Intn(26) + 65)) // Random letters
	}

	// Print new token at server console
	fmt.Println("New token:", newToken)

	// Save new user into mocked data
	mockLoginDetails[user.Username] = ent.LoginDetails{
		AuthToken: newToken,
		Username:  user.Username,
		Mail:      user.Mail,
	}

	// Append new user in slice response
	clientData = append(clientData, user)

	// Return new completed list
	return clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
