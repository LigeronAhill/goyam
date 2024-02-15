package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/LigeronAhill/goyam/models"
)

type MarketClient struct {
	accessToken string
	campaignId  int
	businessId  int
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
	fmt.Printf("%+v\n", c)
	st, err := c.GetAllCampaigns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", st)
	offers, err := c.GetOfferMappings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Offers on YandexMarket - %v\n", len(offers))
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
func (m *MarketClient) GetOfferMappings() ([]models.GetOfferMappingDTO, error) {
	base, err := url.Parse(fmt.Sprintf("https://api.partner.market.yandex.ru/businesses/%v/offer-mappings", m.businessId))
	if err != nil {
		return []models.GetOfferMappingDTO{}, err
	}
	var result []models.GetOfferMappingDTO
	nextPageToken := ""
	for {
		params := url.Values{}
		params.Add("limit", "20")
		if nextPageToken != "" {
			params.Add("page_token", nextPageToken)
		}
		base.RawQuery = params.Encode()
		req, err := http.NewRequest(http.MethodPost, base.String(), nil)
		if err != nil {
			return []models.GetOfferMappingDTO{}, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.accessToken))
		req.Header.Set("User-Agent", "GoYam")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return []models.GetOfferMappingDTO{}, err
		}
		defer resp.Body.Close()
		var offersResponse models.OfferMappingsResponse
		if err := json.NewDecoder(resp.Body).Decode(&offersResponse); err != nil {
			return []models.GetOfferMappingDTO{}, err
		}
		result = append(result, offersResponse.Result.OfferMappings...)
		nextPageToken = offersResponse.Result.Paging.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
	return result, nil
}
