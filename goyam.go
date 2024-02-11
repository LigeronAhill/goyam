package goyam

import (
	"os"

	"github.com/joho/godotenv"
)

type MarketClient struct {
	accessToken string
}

func Init(m *MarketClient) MarketClient {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	var accessToken = os.Getenv("ACCESS_TOKEN")
	return MarketClient{
		accessToken: accessToken,
	}
}
