package models

type CampaignResponse struct {
	Campaigns []CampaignDTO    `json:"campaigns"`
	Pager     FlippingPagerDTO `json:"pager"`
}
type CampaignDTO struct {
	Domain        string      `json:"domain"`
	Id            int         `json:"id"`
	ClientId      int         `json:"clientId"`
	Business      BusinessDTO `json:"business"`
	PlacementType string      `json:"placementType"`
}
type FlippingPagerDTO struct {
	Total       int `json:"total"`
	From        int `json:"from"`
	To          int `json:"to"`
	CurrentPage int `json:"currentPage"`
	PagesCount  int `json:"pagesCount"`
	PageSize    int `json:"pageSize"`
}
type BusinessDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
