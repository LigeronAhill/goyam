package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/LigeronAhill/goyam/models"
)

type MarketClient struct {
	accessToken string
	campaignId  int
	businessId  int
}

func (m *MarketClient) GetAllCampaigns() ([]models.CampaignDTO, error) {
	var url = "https://api.partner.market.yandex.ru/campaigns"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.CampaignDTO{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.accessToken))
	req.Header.Set("User-Agent", "GoYam")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []models.CampaignDTO{}, err
	}
	defer resp.Body.Close()
	var campaign_response models.CampaignResponse
	if err := json.NewDecoder(resp.Body).Decode(&campaign_response); err != nil {
		return []models.CampaignDTO{}, err
	}
	return campaign_response.Campaigns, nil

}

func NewClientFromEnv() (MarketClient, error) {
	var accessToken = os.Getenv("MARKET_TOKEN")
	mc := MarketClient{accessToken: accessToken}
	campaigns, err := mc.GetAllCampaigns()
	if err != nil {
		return MarketClient{}, err
	}
	if len(campaigns) > 0 {
		return MarketClient{
			accessToken: accessToken,
			campaignId:  campaigns[0].Id,
			businessId:  campaigns[0].Business.Id,
		}, nil
	} else {
		return MarketClient{
			accessToken: accessToken,
		}, errors.New("no campaigns!")
	}
}

func main() {
	var c, err = NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
}
