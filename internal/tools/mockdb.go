package tools

import (
	"log"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"lopes": {
		AuthToken: "12345ABC",
		Username:  "lopes",
	},
	"noah": {
		AuthToken: "ABC12345",
		Username:  "noah",
	},
	"kieza": {
		AuthToken: "456789DEF",
		Username:  "kieza",
	},
	"samuel": {
		AuthToken: "67890DEF",
		Username:  "samuel",
	},
	"naiara": {
		AuthToken: "DEF45678",
		Username:  "naiara",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"lopes": {
		Coins:    100,
		Username: "lopes",
	},
	"noah": {
		Coins:    200,
		Username: "noah",
	},
	"kieza": {
		Coins:    300,
		Username: "kieza",
	},
	"samuel": {
		Coins:    400,
		Username: "samuel",
	},
	"naiara": {
		Coins:    500,
		Username: "naiara",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	log.Println(clientData)

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {

	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
